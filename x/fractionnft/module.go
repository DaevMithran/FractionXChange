package module

import (
	"context"
	"encoding/json"
	"strings"

	"github.com/gorilla/mux"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"

	abci "github.com/cometbft/cometbft/abci/types"

	"cosmossdk.io/client/v2/autocli"
	errorsmod "cosmossdk.io/errors"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"

	"github.com/MANTRA-Chain/mantrachain/x/fractionnft/keeper"
	"github.com/MANTRA-Chain/mantrachain/x/fractionnft/types"

	nftkeeper "cosmossdk.io/x/nft/keeper"
	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
)

const (
	// ConsensusVersion defines the current x/fractionnft module consensus version.
	ConsensusVersion = 1
)

var (
	_ module.AppModuleBasic    = AppModuleBasic{}
	_ module.AppModuleGenesis  = AppModule{}
	_ module.AppModule         = AppModule{}
	_ module.HasABCIEndBlock   = AppModule{}
	_ autocli.HasAutoCLIConfig = AppModule{}
)

// AppModuleBasic defines the basic application module used by the wasm module.
type AppModuleBasic struct {
	cdc codec.Codec
}

type AppModule struct {
	AppModuleBasic

	keeper keeper.Keeper

	accountKeeper authkeeper.AccountKeeper
	nftKeeper     nftkeeper.Keeper
	bankKeeper    bankkeeper.Keeper
}

// EndBlock implements module.HasABCIEndBlock.
func (am AppModule) EndBlock(ctx context.Context) ([]abci.ValidatorUpdate, error) {
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	blockHeight := sdkCtx.BlockHeight()

	tokenizedNFT, err := am.keeper.NFTMapping.Iterate(ctx, nil)
	if err != nil {
		return nil, err
	}

	keyValues, err := tokenizedNFT.KeyValues()
	if err != nil {
		return nil, err
	}

	for _, v := range keyValues {
		if v.Value.TimeoutHeight == blockHeight {
			splitted := strings.Split(v.Key, "-")
			am.keeper.RemintNFT(ctx, splitted[1], splitted[2], authtypes.NewModuleAddress(types.ModuleName))
		}
	}
	return nil, nil
}

// Name implements module.HasABCIEndBlock.
// Subtle: this method shadows the method (AppModuleBasic).Name of AppModule.AppModuleBasic.
func (am AppModule) Name() string {
	return types.ModuleName
}

// RegisterGRPCGatewayRoutes implements module.HasABCIEndBlock.
// Subtle: this method shadows the method (AppModuleBasic).RegisterGRPCGatewayRoutes of AppModule.AppModuleBasic.
func (am AppModule) RegisterGRPCGatewayRoutes(clientCtx client.Context, mux *runtime.ServeMux) {
	err := types.RegisterQueryHandlerClient(context.Background(), mux, types.NewQueryClient(clientCtx))
	if err != nil {
		// same behavior as in cosmos-sdk
		panic(err)
	}
}

// RegisterInterfaces implements module.HasABCIEndBlock.
// Subtle: this method shadows the method (AppModuleBasic).RegisterInterfaces of AppModule.AppModuleBasic.
func (am AppModule) RegisterInterfaces(r codectypes.InterfaceRegistry) {
	types.RegisterInterfaces(r)
}

// RegisterLegacyAminoCodec implements module.HasABCIEndBlock.
// Subtle: this method shadows the method (AppModuleBasic).RegisterLegacyAminoCodec of AppModule.AppModuleBasic.
func (am AppModule) RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {
	types.RegisterLegacyAminoCodec(cdc)
}

// NewAppModule constructor
func NewAppModule(
	cdc codec.Codec,
	keeper keeper.Keeper,
	accountKeeper authkeeper.AccountKeeper,
	nftKeeper nftkeeper.Keeper,
	bankKeeper bankkeeper.Keeper,
) *AppModule {
	return &AppModule{
		AppModuleBasic: AppModuleBasic{cdc: cdc},
		keeper:         keeper,
		accountKeeper:  accountKeeper,
		nftKeeper:      nftKeeper,
		bankKeeper:     bankKeeper,
	}
}

func (a AppModuleBasic) Name() string {
	return types.ModuleName
}

func (a AppModuleBasic) DefaultGenesis(cdc codec.JSONCodec) json.RawMessage {
	return cdc.MustMarshalJSON(&types.GenesisState{
		Params: types.DefaultParams(),
	})
}

func (a AppModuleBasic) ValidateGenesis(marshaler codec.JSONCodec, _ client.TxEncodingConfig, message json.RawMessage) error {
	var data types.GenesisState
	err := marshaler.UnmarshalJSON(message, &data)
	if err != nil {
		return err
	}
	if err := data.Params.Validate(); err != nil {
		return errorsmod.Wrap(err, "params")
	}
	return nil
}

func (a AppModuleBasic) RegisterRESTRoutes(_ client.Context, _ *mux.Router) {
}

func (a AppModuleBasic) RegisterGRPCGatewayRoutes(clientCtx client.Context, mux *runtime.ServeMux) {
	err := types.RegisterQueryHandlerClient(context.Background(), mux, types.NewQueryClient(clientCtx))
	if err != nil {
		// same behavior as in cosmos-sdk
		panic(err)
	}
}

// Disable in favor of autocli.go. If you wish to use these, it will override AutoCLI methods.
/*
func (a AppModuleBasic) GetTxCmd() *cobra.Command {
	return cli.NewTxCmd()
}

func (a AppModuleBasic) GetQueryCmd() *cobra.Command {
	return cli.GetQueryCmd()
}
*/

func (AppModuleBasic) RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {
	types.RegisterLegacyAminoCodec(cdc)
}

func (a AppModuleBasic) RegisterInterfaces(r codectypes.InterfaceRegistry) {
	types.RegisterInterfaces(r)
}

func (a AppModule) InitGenesis(ctx sdk.Context, marshaler codec.JSONCodec, message json.RawMessage) []abci.ValidatorUpdate {
	var genesisState types.GenesisState
	marshaler.MustUnmarshalJSON(message, &genesisState)

	if err := a.keeper.Params.Set(ctx, genesisState.Params); err != nil {
		panic(err)
	}

	return nil
}

func (a AppModule) ExportGenesis(ctx sdk.Context, marshaler codec.JSONCodec) json.RawMessage {
	genState := a.keeper.ExportGenesis(ctx)
	return marshaler.MustMarshalJSON(genState)
}

func (a AppModule) RegisterInvariants(_ sdk.InvariantRegistry) {
}

func (a AppModule) QuerierRoute() string {
	return types.QuerierRoute
}

func (a AppModule) RegisterServices(cfg module.Configurator) {
	types.RegisterMsgServer(cfg.MsgServer(), keeper.NewMsgServerImpl(a.keeper))
	types.RegisterQueryServer(cfg.QueryServer(), keeper.NewQuerier(a.keeper))
}

// ConsensusVersion is a sequence number for state-breaking change of the
// module. It should be incremented on each consensus-breaking change
// introduced by the module. To avoid wrong/empty versions, the initial version
// should be set to 1.
func (a AppModule) ConsensusVersion() uint64 {
	return ConsensusVersion
}
