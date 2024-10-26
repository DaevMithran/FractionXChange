package keeper

import (
	"context"
	"fmt"

	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"

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
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
)

type Keeper struct {
	cdc          codec.BinaryCodec
	addressCodec address.Codec

	logger log.Logger

	// state management
	Schema collections.Schema
	Params collections.Item[types.Params]
	OrmDB  apiv1.StateStore

	authority string

	NFTMapping collections.Map[string, types.TokenizedNFT]

	AccountKeeper authkeeper.AccountKeeper
	Nftkeeper     nftkeeper.Keeper
	BankKeeper    bankkeeper.Keeper
}

// NewKeeper creates a new Keeper instance
func NewKeeper(
	cdc codec.BinaryCodec,
	addressCodec address.Codec,
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
		cdc:          cdc,
		addressCodec: addressCodec,
		logger:       logger,

		Params: collections.NewItem(sb, types.ParamsKey, "params", codec.CollValue[types.Params](cdc)),
		OrmDB:  store,

		authority: authority,

		NFTMapping: collections.NewMap(sb, collections.NewPrefix(1), "nft_mapping", collections.StringKey, codec.CollValue[types.TokenizedNFT](cdc)),

		AccountKeeper: accountKepper,
		Nftkeeper:     nftKeeper,
		BankKeeper:    bankKeeper,
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
	if !k.Nftkeeper.HasClass(ctx, classId) {
		return fmt.Errorf("NFT class not found: %s", classId)
	}

	if !k.Nftkeeper.HasNFT(ctx, classId, nftId) {
		return fmt.Errorf("NFT not found: %s", nftId)
	}

	sdkCtx := sdk.UnwrapSDKContext(ctx)
	blockHeight := sdkCtx.BlockHeight()
	// Restricting timeout: timeout should be alteast greater than 3 blocks
	if blockHeight+3 >= timeoutHeight {
		return fmt.Errorf("timeout should be greater than curent blockHeight %d", blockHeight)
	}

	denom := "fractionNFT-" + classId + "-" + nftId

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
	_, err = k.Nftkeeper.Send(ctx, &nft.MsgSend{
		ClassId:  classId,
		Id:       nftId,
		Sender:   sender.String(),
		Receiver: authtypes.NewModuleAddress(types.ModuleName).String(),
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
		Owner:         sender.String(),
		TimeoutHeight: timeoutHeight,
	})
}

func (k *Keeper) RemintNFT(ctx context.Context, classId string, nftId string, sender sdk.AccAddress) error {
	denom := "fractionNFT-" + classId + "-" + nftId

	// validate if NFT is tokenized already
	tokenizedNFT, err := k.NFTMapping.Get(ctx, denom)
	if err != nil {
		return err
	}

	// validate block height
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	blockHeight := sdkCtx.BlockHeight()
	senderStr := sender.String()
	fmt.Printf("senderStr %s\n", senderStr)
	fmt.Printf("moduleAddr %s", authtypes.NewModuleAddress(types.ModuleName).String())
	switch senderStr {
	case authtypes.NewModuleAddress(types.ModuleName).String():
		if blockHeight < tokenizedNFT.TimeoutHeight {
			return fmt.Errorf("remint can be performed on timeout: %s", denom)
		}
		// Burn tokens in the denom
		err = k.BurnAllTokensOfDenom(sdkCtx, denom)
		if err != nil {
			return err
		}
	case tokenizedNFT.Owner:
		supply := k.BankKeeper.GetSupply(ctx, denom)
		balance := k.BankKeeper.GetBalance(ctx, sender, denom)

		if !supply.Amount.Equal(balance.Amount) {
			return fmt.Errorf("remint can be called by owner only if they own all the tokens: %s", denom)
		}

		coins := sdk.NewCoins(balance)

		err = k.BankKeeper.SendCoinsFromAccountToModule(ctx, sender, types.ModuleName, coins)
		if err != nil {
			return err
		}

		k.BankKeeper.BurnCoins(ctx, types.ModuleName, coins)
	default:
		return fmt.Errorf("unauthorized sender: %s", senderStr)
	}

	// Transfer NFT back to the owner
	_, err = k.Nftkeeper.Send(ctx, &nft.MsgSend{
		ClassId:  classId,
		Id:       nftId,
		Sender:   authtypes.NewModuleAddress(types.ModuleName).String(),
		Receiver: tokenizedNFT.Owner,
	})
	if err != nil {
		return err
	}

	// Remove from NFT Mapping
	return k.NFTMapping.Remove(ctx, denom)
}

func (k Keeper) BurnAllTokensOfDenom(ctx sdk.Context, denom string) error {
	fmt.Printf("Burning All tokens of Denom")
	owners, err := k.BankKeeper.DenomOwners(ctx, &banktypes.QueryDenomOwnersRequest{
		Denom: denom,
	})
	if err != nil {
		return err
	}

	for _, owner := range owners.DenomOwners {
		addr, err := sdk.AccAddressFromBech32(owner.Address)
		if err != nil {
			return err
		}
		k.BankKeeper.SendCoinsFromAccountToModule(ctx, addr, types.ModuleName, sdk.NewCoins(owner.Balance))
	}

	coin := k.BankKeeper.GetBalance(ctx, authtypes.NewModuleAddress(types.ModuleName), denom)
	k.BankKeeper.BurnCoins(ctx, types.ModuleName, sdk.NewCoins(coin))

	return nil
}

func (k Keeper) MintNFT(ctx context.Context, owner sdk.AccAddress, category string, id string, name string, description string, image string) error {
	if !k.Nftkeeper.HasClass(ctx, category) {
		k.Nftkeeper.SaveClass(ctx, nft.Class{Id: category, Name: category})
	}

	if k.Nftkeeper.HasNFT(ctx, category, id) {
		return nil
	}

	coins := sdk.NewCoins(sdk.NewCoin("uom", math.NewInt(int64(5))))
	// receive payment
	err := k.BankKeeper.SendCoinsFromAccountToModule(ctx, owner, types.ModuleName, coins)
	if err != nil {
		return err
	}

	data, err := codectypes.NewAnyWithValue(&types.NFTData{
		Name:        name,
		Description: description,
		Image:       image,
	})
	if err != nil {
		return err
	}

	return k.Nftkeeper.Mint(ctx, nft.NFT{
		ClassId: category,
		Id:      id,
		Data:    data,
	}, owner)
}
