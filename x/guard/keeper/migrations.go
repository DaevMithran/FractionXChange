package keeper

import (
	"github.com/MANTRA-Finance/mantrachain/x/guard/exported"
	v3 "github.com/MANTRA-Finance/mantrachain/x/guard/migrations/v3"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// Migrator is a struct for handling in-place store migrations.
type Migrator struct {
	keeper                      *Keeper
	legacySubspace              exported.Subspace
	whitelistedTransferAccAddrs []string
}

// NewMigrator returns a new Migrator.
func NewMigrator(keeper *Keeper, legacySubspace exported.Subspace, whitelistedTransferAccAddrs []string) Migrator {
	return Migrator{
		keeper:                      keeper,
		legacySubspace:              legacySubspace,
		whitelistedTransferAccAddrs: whitelistedTransferAccAddrs,
	}
}

// Migrate2to3 migrates the store from consensus version 2 to 3
func (m Migrator) Migrate2to3(ctx sdk.Context) error {
	return v3.MigrateStore(ctx, m.keeper.storeService, m.keeper.cdc, m.legacySubspace, m.whitelistedTransferAccAddrs)
}
