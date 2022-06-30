package keeper

import (
	"context"

	"github.com/LimeChain/mantrachain/x/mdb/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) NftCollection(c context.Context, req *types.QueryGetNftCollectionRequest) (*types.QueryGetNftCollectionResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	creator, err := sdk.AccAddressFromBech32(req.Creator)

	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	if req.Id == "" {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	index := types.GetNftCollectionIndex(creator, req.Id)

	meta, found := k.GetNftCollection(
		ctx,
		sdk.AccAddress(creator),
		index,
	)
	if !found {
		return nil, status.Error(codes.InvalidArgument, "not found")
	}

	nftExecutor := NewNftExecutor(ctx, k.nftKeeper)
	nftColl, found := nftExecutor.GetClass(ctx, string(index))
	if !found {
		return nil, status.Error(codes.InvalidArgument, "not found")
	}

	return &types.QueryGetNftCollectionResponse{
		Id:          nftColl.UriHash,
		Name:        nftColl.Name,
		Symbol:      nftColl.Symbol,
		Description: nftColl.Description,
		Did:         meta.Did,
		Images:      meta.Images,
		Url:         meta.Url,
		Links:       meta.Links,
		Category:    meta.Category,
		Options:     meta.Options,
		Creator:     meta.Creator.String(),
		Owner:       meta.Owner.String(),
		Opened:      meta.Opened,
		Data:        nftColl.Data,
	}, nil
}
