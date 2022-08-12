package keeper

import (
	"fmt"

	"github.com/tendermint/tendermint/libs/log"

	"github.com/LimeChain/mantrachain/x/marketplace/types"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
)

type (
	Keeper struct {
		cdc        codec.BinaryCodec
		storeKey   sdk.StoreKey
		memKey     sdk.StoreKey
		paramstore paramtypes.Subspace

		ac          types.AccountKeeper
		bk          types.BankKeeper
		tokenKeeper types.TokenKeeper
		nftKeeper   types.NFTKeeper
		vaultKeeper types.VaultKeeper
	}
)

func NewKeeper(
	cdc codec.BinaryCodec,
	storeKey,
	memKey sdk.StoreKey,
	ps paramtypes.Subspace,

	ac types.AccountKeeper,
	bk types.BankKeeper,
	tokenKeeper types.TokenKeeper,
	nftKeeper types.NFTKeeper,
	vaultKeeper types.VaultKeeper,
) *Keeper {
	// set KeyTable if it has not already been set
	if !ps.HasKeyTable() {
		ps = ps.WithKeyTable(types.ParamKeyTable())
	}

	return &Keeper{

		cdc:        cdc,
		storeKey:   storeKey,
		memKey:     memKey,
		paramstore: ps,

		ac: ac, bk: bk, tokenKeeper: tokenKeeper, nftKeeper: nftKeeper, vaultKeeper: vaultKeeper,
	}
}

func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}
