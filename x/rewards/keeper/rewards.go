package keeper

import (
	"math"

	sdkmath "cosmossdk.io/math"
	"github.com/MANTRA-Finance/mantrachain/x/rewards/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) CalculateRewards(ctx sdk.Context, pairId uint64, provider types.Provider, params *types.ClaimParams) (types.Provider, error) {
	conf := k.GetParams(ctx)
	minDepositTime := conf.MinDepositTime
	startClaimedSnapshotId := uint64(0)
	endClaimedSnapshotId := uint64(0)

	if params != nil && params.EndClaimedSnapshotId != nil {
		endClaimedSnapshotId = *params.EndClaimedSnapshotId
	} else {
		endClaimedSnapshotId = k.GetEndClaimedSnapshotId(ctx, pairId)
	}

	pairIdx, found := provider.PairIdToIdx[pairId]
	if !found {
		return provider, nil
	}

	providerPair := provider.Pairs[pairIdx]

	if providerPair == nil {
		return provider, nil
	}

	if params != nil && params.StartClaimedSnapshotId != nil {
		startClaimedSnapshotId = *params.StartClaimedSnapshotId
	} else {
		startClaimedSnapshotId = providerPair.LastClaimedSnapshotId

		// If LastClaimedSnapshotId is zero then the provider has never claimed rewards
		if startClaimedSnapshotId == uint64(math.MaxUint64) {
			startClaimedSnapshotId = 0
		} else {
			startClaimedSnapshotId++
		}

		snapshotStartId, found := k.GetSnapshotStartId(ctx, pairId)

		if !found {
			snapshotStartId = types.SnapshotStartId{
				PairId:     pairId,
				SnapshotId: 0,
			}
		}

		// If we purge the snapshots then we need to start from the first available snapshot if it is greater than the last claimed snapshot
		if snapshotStartId.SnapshotId > startClaimedSnapshotId {
			startClaimedSnapshotId = snapshotStartId.SnapshotId
		}
	}

	if conf.MaxClaimedRangeLength > 0 && endClaimedSnapshotId-startClaimedSnapshotId >= conf.MaxClaimedRangeLength {
		endClaimedSnapshotId = startClaimedSnapshotId + conf.MaxClaimedRangeLength - 1
	}

	// If the provider has never claimed rewards then we need to start from the first snapshot
	snapshots := k.GetSnapshotsInRange(ctx, pairId, startClaimedSnapshotId, endClaimedSnapshotId)

	for _, snapshot := range snapshots {
		if !snapshot.Distributed {
			continue
		}

		if minDepositTime > 0 && snapshot.DistributedAt.Unix() < providerPair.LastDepositTime.Unix()+int64(minDepositTime) {
			continue
		}

		for _, pool := range snapshot.Pools {
			if pool.CoinSupply.IsPositive() {
				poolIdx, found := providerPair.PoolIdToIdx[pool.PoolId]
				if !found {
					continue
				}

				balance := providerPair.Balances[poolIdx]

				if balance.IsPositive() {
					for _, rewardPerToken := range pool.RewardsPerToken {
						reward := sdk.NewDecCoinFromDec(rewardPerToken.Denom, rewardPerToken.Amount.Mul(sdkmath.LegacyNewDecFromInt(balance.Amount)).TruncateDec())
						// In case of a mismatch between the rewards and the remaining rewards, we need to adjust the rewards
						for _, remaining := range snapshot.Remaining {
							if remaining.Denom == reward.Denom && remaining.IsLT(reward) {
								reward.Amount = remaining.Amount
							}
						}
						providerPair.OwedRewards = providerPair.OwedRewards.Add(reward)

						// If the provider is claiming rewards then we need to update the snapshot
						if params != nil && !params.IsQuery {
							snapshot.Remaining = snapshot.Remaining.Sub(sdk.NewDecCoins(reward))
							k.SetSnapshot(ctx, snapshot)
						}
					}
				}
			}

		}
	}

	provider.Pairs[pairIdx].LastClaimedSnapshotId = endClaimedSnapshotId

	return provider, nil
}
