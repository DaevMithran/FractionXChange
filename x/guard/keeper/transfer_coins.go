package keeper

import (
	"strings"

	"cosmossdk.io/errors"
	tokentypes "github.com/LimeChain/mantrachain/x/token/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"

	"github.com/LimeChain/mantrachain/x/guard/types"
)

func (k Keeper) CheckCanTransferCoins(ctx sdk.Context, address string, coins sdk.Coins) error {
	accAddress, err := sdk.AccAddressFromBech32(address)
	if err != nil {
		return err
	}

	conf := k.GetParams(ctx)

	// Check if it is a module address or it is an admin wallet address
	if k.modAccAddrs[address] || sdk.MustAccAddressFromBech32(conf.AdminAccount).Equals(accAddress) {
		return nil
	}

	collectionCreator := conf.AccountPrivilegesTokenCollectionCreator
	collectionId := conf.AccountPrivilegesTokenCollectionId

	if strings.TrimSpace(collectionId) == "" {
		return errors.Wrap(types.ErrInvalidTokenCollectionId, "nft collection id should not be empty")
	}

	creator, err := sdk.AccAddressFromBech32(collectionCreator)

	if err != nil {
		return errors.Wrap(types.ErrInvalidTokenCollectionCreator, "collection creator should not be empty")
	}

	collectionIndex := tokentypes.GetNftCollectionIndex(creator, collectionId)

	index := tokentypes.GetNftIndex(collectionIndex, address)

	owner := k.nk.GetOwner(ctx, string(collectionIndex), string(index))

	if owner.Empty() || !accAddress.Equals(owner) {
		return errors.Wrapf(types.ErrIncorrectNftOwner, "incorrect nft owner, address %s", address)
	}

	var indexes [][]byte

	for _, coin := range coins {
		denom := coin.GetDenom()
		hasAdmin := k.ck.HasAdmin(ctx, denom)

		if hasAdmin {
			admin, found := k.ck.GetAdmin(ctx, denom)

			if !found {
				return errors.Wrapf(sdkerrors.ErrInvalidCoins, "coin %s admin invalid", denom)
			}

			if admin.Equals(accAddress) {
				continue
			}

			indexes = append(indexes, []byte(denom))
		}
	}

	if len(indexes) > 0 {
		requiredPrivilegesList := k.GetRequiredPrivilegesMany(ctx, indexes, types.Coin)

		if len(requiredPrivilegesList) == 0 {
			return errors.Wrap(types.ErrRequiredPrivilegesNotFound, "required privileges not found")
		}

		if len(requiredPrivilegesList) != len(indexes) {
			return errors.Wrap(types.ErrCoinRequiredPrivilegesNotFound, "coin required privileges not found")
		}

		hasPrivileges, err := k.CheckAccountFulfillsRequiredPrivileges(ctx, accAddress, requiredPrivilegesList)

		if err != nil {
			return err
		}

		if !hasPrivileges {
			return errors.Wrapf(types.ErrInsufficientPrivileges, "insufficient privileges, address %s", accAddress)
		}
	}

	return nil
}

func (k Keeper) ValidateCanTransferCoins(ctx sdk.Context, inputs []banktypes.Input, outputs []banktypes.Output) error {
	if !k.HasGuardTransferCoins(ctx) {
		return nil
	}

	for _, in := range inputs {
		err := k.CheckCanTransferCoins(ctx, in.Address, in.Coins)

		if err != nil {
			return err
		}
	}

	for _, out := range outputs {
		err := k.CheckCanTransferCoins(ctx, out.Address, out.Coins)

		if err != nil {
			return err
		}
	}

	return nil
}
