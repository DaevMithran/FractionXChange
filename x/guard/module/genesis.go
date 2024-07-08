package guard

import (
	"strings"

	"cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/MANTRA-Finance/mantrachain/x/guard/keeper"
	"github.com/MANTRA-Finance/mantrachain/x/guard/types"
)

// InitGenesis initializes the module's state from a provided genesis state.
func InitGenesis(ctx sdk.Context, k *keeper.Keeper, genState types.GenesisState) {
	// this line is used by starport scaffolding # genesis/module/init
	if err := k.SetParams(ctx, genState.Params); err != nil {
		panic(err)
	}

	if _, err := sdk.AccAddressFromBech32(genState.Params.AccountPrivilegesTokenCollectionCreator); err != nil {
		panic(err)
	}
	if strings.TrimSpace(genState.Params.AccountPrivilegesTokenCollectionId) == "" {
		panic(errors.Wrap(types.ErrInvalidAccountPrivilegesTokenCollectionIdParam, "account privileges token collection id param should not be empty"))
	}
	// Set all the accountPrivileges
	for _, elem := range genState.AccountPrivilegesList {
		k.SetAccountPrivileges(ctx, elem.Account, elem.Privileges)
	}
	// Set if defined
	if genState.GuardTransferCoins != nil {
		k.SetGuardTransferCoins(ctx)
	}
	// Set all the requiredPrivileges
	for _, elem := range genState.RequiredPrivilegesList {
		kind, err := types.ParseRequiredPrivilegesKind(elem.Kind)
		if err != nil {
			panic("kind is invalid")
		}
		k.SetRequiredPrivileges(ctx, elem.Index, kind, elem.Privileges)
	}
}

// ExportGenesis returns the module's exported genesis.
func ExportGenesis(ctx sdk.Context, k *keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)

	genesis.AccountPrivilegesList = k.GetAllAccountPrivileges(ctx)
	// Get all guardTransfer
	if k.HasGuardTransferCoins(ctx) {
		genesis.GuardTransferCoins = types.Placeholder
	}
	requiredPrivilegesCoin := k.GetAllRequiredPrivileges(ctx, types.RequiredPrivilegesCoin)
	requiredPrivilegesAuthz := k.GetAllRequiredPrivileges(ctx, types.RequiredPrivilegesAuthz)
	genesis.RequiredPrivilegesList = append(requiredPrivilegesCoin, requiredPrivilegesAuthz...)
	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}