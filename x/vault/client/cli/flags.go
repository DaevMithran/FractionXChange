package cli

import (
	flag "github.com/spf13/pflag"
)

const (
	FlagMarketplaceCreator = "marketplace-creator"
	FlagMarketplaceId      = "marketplace-id"
	FlagCollectionCreator  = "collection-creator"
	FlagCollectionId       = "collection-id"
	FlagReceiver           = "receiver"
	FlagStakingChain       = "staking-chain"
	FlagStakingValidator   = "staking-validator"
)

var (
	FsWithdrawNftRewards   = flag.NewFlagSet("", flag.ContinueOnError)
	FsUpdateNftStakeStaked = flag.NewFlagSet("", flag.ContinueOnError)
)

func init() {
	FsWithdrawNftRewards.String(FlagMarketplaceCreator, "", "The marketplace creator address")
	FsWithdrawNftRewards.String(FlagMarketplaceId, "", "The marketplace id")
	FsWithdrawNftRewards.String(FlagCollectionCreator, "", "The collection creator address")
	FsWithdrawNftRewards.String(FlagCollectionId, "", "The collection id")
	FsWithdrawNftRewards.String(FlagReceiver, "", "The withdraw reward receiver")
	FsWithdrawNftRewards.String(FlagStakingChain, "", "The staking chain")
	FsWithdrawNftRewards.String(FlagStakingValidator, "", "The staking validator")

	FsUpdateNftStakeStaked.String(FlagMarketplaceCreator, "", "The marketplace creator address")
	FsUpdateNftStakeStaked.String(FlagMarketplaceId, "", "The marketplace id")
	FsUpdateNftStakeStaked.String(FlagCollectionCreator, "", "The collection creator address")
	FsUpdateNftStakeStaked.String(FlagCollectionId, "", "The collection id")
	FsUpdateNftStakeStaked.String(FlagStakingChain, "", "The staking chain")
	FsUpdateNftStakeStaked.String(FlagStakingValidator, "", "The staking validator")
}
