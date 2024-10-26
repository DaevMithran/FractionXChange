package keeper

import (
	"context"
	"fmt"

	"github.com/cosmos/cosmos-sdk/codec"

	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"

	"cosmossdk.io/collections"
	"cosmossdk.io/core/address"
	storetypes "cosmossdk.io/core/store"
	"cosmossdk.io/log"
	"cosmossdk.io/math"
	"cosmossdk.io/orm/model/ormdb"

	apiv1 "github.com/MANTRA-Chain/mantrachain/api/fractionnft/v1"
	"github.com/MANTRA-Chain/mantrachain/x/fractionnft/types"

	nft "cosmossdk.io/x/nft"
	nftkeeper "cosmossdk.io/x/nft/keeper"
	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
)

type Keeper struct {
	cdc codec.BinaryCodec
	addressCodec address.Codec

	logger log.Logger

	// state management
	Schema collections.Schema
	Params collections.Item[types.Params]
	OrmDB  apiv1.StateStore

	authority string

	NFTMapping collections.Map[string, types.TokenizedNFT]

	AccountKeeper authkeeper.AccountKeeper
	Nftkeeper nftkeeper.Keeper
	BankKeeper bankkeeper.Keeper
}

// NewKeeper creates a new Keeper instance
func NewKeeper(
	cdc codec.BinaryCodec,
	storeService storetypes.KVStoreService,
	logger log.Logger,
	authority string,
	accountKepper authkeeper.AccountKeeper,
	nftKeeper nftkeeper.Keeper,
	bankKeeper bankkeeper.Keeper,
) Keeper {
	logger = logger.With(log.ModuleKey, "x/"+types.ModuleName)

	sb := collections.NewSchemaBuilder(storeService)

	if authority == "" {
		authority = authtypes.NewModuleAddress(govtypes.ModuleName).String()
	}

	db, err := ormdb.NewModuleDB(&types.ORMModuleSchema, ormdb.ModuleDBOptions{KVStoreService: storeService})
	if err != nil {
		panic(err)
	}

	store, err := apiv1.NewStateStore(db)
	if err != nil {
		panic(err)
	}

	

	k := Keeper{
		cdc:    cdc,
		logger: logger,

		Params: collections.NewItem(sb, types.ParamsKey, "params", codec.CollValue[types.Params](cdc)),
		OrmDB:  store,

		authority: authority,

		NFTMapping: collections.NewMap(sb, collections.NewPrefix(1), "nft_mapping", collections.StringKey, codec.CollValue[types.TokenizedNFT](cdc)),

		AccountKeeper: accountKepper,
		Nftkeeper: nftKeeper,
		BankKeeper: bankKeeper,
	}

	schema, err := sb.Build()
	if err != nil {
		panic(err)
	}

	k.Schema = schema

	return k
}

func (k Keeper) Logger() log.Logger {
	return k.logger
}

// InitGenesis initializes the module's state from a genesis state.
func (k *Keeper) InitGenesis(ctx context.Context, data *types.GenesisState) error {

	if err := data.Params.Validate(); err != nil {
		return err
	}

	return k.Params.Set(ctx, data.Params)
}

// ExportGenesis exports the module's state to a genesis state.
func (k *Keeper) ExportGenesis(ctx context.Context) *types.GenesisState {
	params, err := k.Params.Get(ctx)
	if err != nil {
		panic(err)
	}

	return &types.GenesisState{
		Params: params,
	}
}

func (k *Keeper) TokenizeNFT(ctx context.Context, classId string, nftId string, sender sdk.AccAddress, tokenSupply int64, timeoutHeight int64) error {
	if !k.Nftkeeper.HasClass(ctx, nftId) {
        return fmt.Errorf("NFT class not found: %s", classId)
	}

	if !k.Nftkeeper.HasNFT(ctx, classId, nftId) {
        return fmt.Errorf("NFT not found: %s", nftId)
	}

	denom := "fractionNFT-"+classId+"-"+nftId
	
	// validate if NFT is tokenized already
    has, err := k.NFTMapping.Has(ctx, denom)
    if err != nil {
        return err
    }
    if has {
        return fmt.Errorf("NFT is tokenized already : %s", denom)
    }

	// Lock the NFT in module address
	// Note: Owner validation is done within the Send function
	_, err = k.Nftkeeper.Send(ctx, &nft.MsgSend {
		ClassId: classId, 
		Id: nftId, 
		Sender: sender.String(),
		Receiver: types.ModuleName,
	})
	if err != nil {
		return err
	}

	coins := sdk.NewCoins(sdk.NewCoin(denom, math.NewInt(tokenSupply)))
	
	// tokenize NFT to create new coins
	err = k.BankKeeper.MintCoins(ctx, types.ModuleName, coins)
	if err != nil {
		return err
	}

	// sends coins to the NFT owner account
	err = k.BankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, sender, coins)
	if err != nil {
		return err
	}

	return k.NFTMapping.Set(ctx, denom, types.TokenizedNFT{
		Owner: sender.String(),
		TimeoutHeight: timeoutHeight,
	})
}

func(k *Keeper) RemintNFT(ctx context.Context, classId string, nftId string, sender sdk.AccAddress) error {
	if !k.Nftkeeper.HasClass(ctx, nftId) {
        return fmt.Errorf("NFT class not found: %s", classId)
	}

	if !k.Nftkeeper.HasNFT(ctx, classId, nftId) {
        return fmt.Errorf("NFT not found: %s", nftId)
	}

	denom := "fractionNFT-"+classId+"-"+nftId
	
	// validate if NFT is tokenized already
    tokenizedNFT, err := k.NFTMapping.Get(ctx, denom)
    if err != nil {
        return err
    }

	// validate block height
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	blockHeight := sdkCtx.BlockHeight()
	senderStr := sender.String()

	switch senderStr {
	case types.ModuleName:
		if blockHeight < tokenizedNFT.TimeoutHeight {
			return fmt.Errorf("remint can be performed on timeout: %s", denom)
		}
	case tokenizedNFT.Owner:
		supply := k.BankKeeper.GetSupply(ctx, denom)
		balance := k.BankKeeper.GetBalance(ctx, sender, denom)

		if supply != balance {
			return fmt.Errorf("remint can be called by owner only if they own all the tokens: %s", denom)
		}
	default:
		return fmt.Errorf("unauthorized sender: %s", senderStr)
	}

	// Burn tokens in the denom
	err = k.BurnAllTokensOfDenom(sdkCtx, denom)
	if err != nil {
		return err
	}

	// Transfer NFT back to the owner
	_, err = k.Nftkeeper.Send(ctx, &nft.MsgSend {
		ClassId: classId, 
		Id: nftId, 
		Sender: types.ModuleName,
		Receiver: tokenizedNFT.Owner,
	})
	if err != nil {
		return err
	}

	return nil
}

func (k Keeper) BurnAllTokensOfDenom(ctx sdk.Context, denom string) error {
    // Iterate over all accounts in the state
    k.AccountKeeper.IterateAccounts(ctx, func(account sdk.AccountI) bool {
        // Get the balance of the specified denom for this account
        balance := k.BankKeeper.GetBalance(ctx, account.GetAddress(), denom)
        // If balance is greater than zero, burn the balance
        if !balance.IsZero() {
			k.BankKeeper.BurnCoins(ctx, types.ModuleName, sdk.NewCoins(balance))
        }

        return false // continue iteration
    })

    return nil
}
