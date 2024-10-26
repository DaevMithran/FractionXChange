package keeper

import (
	"context"

	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"

	"cosmossdk.io/errors"
	"github.com/MANTRA-Chain/mantrachain/x/fractionnft/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

type msgServer struct {
	k Keeper
}

var _ types.MsgServer = msgServer{}

// NewMsgServerImpl returns an implementation of the module MsgServer interface.
func NewMsgServerImpl(keeper Keeper) types.MsgServer {
	return &msgServer{k: keeper}
}

func (ms msgServer) UpdateParams(ctx context.Context, msg *types.MsgUpdateParams) (*types.MsgUpdateParamsResponse, error) {
	if ms.k.authority != msg.Authority {
		return nil, errors.Wrapf(govtypes.ErrInvalidSigner, "invalid authority; expected %s, got %s", ms.k.authority, msg.Authority)
	}

	return nil, ms.k.Params.Set(ctx, msg.Params)
}

// MsgTokenizeNFT implements types.MsgServer.
func (ms msgServer) MsgTokenizeNFT(ctx context.Context, msg *types.MsgTokenizeNFTParams) (*types.MsgTokenizeNFTResponse, error) {
	owner, err := ms.k.addressCodec.StringToBytes(msg.Owner)
	if err != nil {
		return nil, sdkerrors.ErrInvalidAddress.Wrapf("invalid owner address: %s", err)
	}

	err = ms.k.TokenizeNFT(ctx, msg.ClassId, msg.NftId, owner, msg.TokenSupply, msg.TimeoutHeight)
	if err != nil {
		return nil, err
	}

	return &types.MsgTokenizeNFTResponse{}, nil
}

// MsgTokenizeNFT implements types.MsgServer.
func (ms msgServer) MsgRemintNFT(ctx context.Context, msg *types.MsgRemintNFTParams) (*types.MsgRemintNFTResponse, error) {
	owner, err := ms.k.addressCodec.StringToBytes(msg.Owner)
	if err != nil {
		return nil, sdkerrors.ErrInvalidAddress.Wrapf("invalid owner address: %s", err)
	}

	err = ms.k.RemintNFT(ctx, msg.ClassId, msg.NftId, owner)
	if err != nil {
		return nil, err
	}
	return &types.MsgRemintNFTResponse{}, nil
}

func (ms msgServer) MsgMintNFT(ctx context.Context, msg *types.MsgMintNFTParams) (*types.MsgMintNFTResponse, error) {
	owner, err := ms.k.addressCodec.StringToBytes(msg.Owner)
	if err != nil {
		return nil, sdkerrors.ErrInvalidAddress.Wrapf("invalid owner address: %s", err)
	}

	err = ms.k.MintNFT(ctx, owner, msg.Category, msg.Id, msg.Name, msg.Description, msg.Image)
	if err != nil {
		return nil, err
	}
	return &types.MsgMintNFTResponse{}, nil
}
