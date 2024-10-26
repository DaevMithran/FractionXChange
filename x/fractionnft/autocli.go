package module

import (
	autocliv1 "cosmossdk.io/api/cosmos/autocli/v1"
	modulev1 "github.com/MANTRA-Chain/mantrachain/api/fractionnft/v1"
)

// AutoCLIOptions implements the autocli.HasAutoCLIConfig interface.
func (am AppModule) AutoCLIOptions() *autocliv1.ModuleOptions {
	return &autocliv1.ModuleOptions{
		Query: &autocliv1.ServiceCommandDescriptor{
			Service: modulev1.Query_ServiceDesc.ServiceName,
			RpcCommandOptions: []*autocliv1.RpcCommandOptions{
				{
                    RpcMethod: "GetNftToken",
                    Use:       "get <class_id> <nft_id>",
                    Short:     "Get nfttoken",
                    PositionalArgs: []*autocliv1.PositionalArgDescriptor{
                        {ProtoField: "class_id"},
						{ProtoField: "nft_id"},
                    },
                },
			},
		},
		Tx: &autocliv1.ServiceCommandDescriptor{
			Service: modulev1.Msg_ServiceDesc.ServiceName,
			RpcCommandOptions: []*autocliv1.RpcCommandOptions{
				{
                    RpcMethod: "MsgTokenizeNFT",
                    Use:       "tokenize <class_id> <nft_id> <token_supply> <timeout_height>",
                    Short:     "Tokenize nft",
                    PositionalArgs: []*autocliv1.PositionalArgDescriptor{
                        {ProtoField: "class_id"},
						{ProtoField: "nft_id"},
						{ProtoField: "token_supply"},
						{ProtoField: "timeout_height"},
                    },
                },
				{
                    RpcMethod: "MsgRemintNFT",
                    Use:       "remint <class_id> <nft_id>",
                    Short:     "Remint nfttoken",
                    PositionalArgs: []*autocliv1.PositionalArgDescriptor{
                        {ProtoField: "class_id"},
						{ProtoField: "nft_id"},
                    },
                },
				{
                    RpcMethod: "MsgMintNFT",
                    Use:       "mint ",
                    Short:     "Mint nfttoken",
                    PositionalArgs: []*autocliv1.PositionalArgDescriptor{
                        {ProtoField: "name"},
						{ProtoField: "id"},
						{ProtoField: "description"},
						{ProtoField: "image"},
						{ProtoField: "category"},
                    },
                },
			},
		},
	}
}
