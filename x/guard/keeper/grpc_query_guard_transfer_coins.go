package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"mantrachain/x/guard/types"
)

func (k Keeper) GuardTransferCoins(c context.Context, req *types.QueryGetGuardTransferCoinsRequest) (*types.QueryGetGuardTransferCoinsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(c)

	return &types.QueryGetGuardTransferCoinsResponse{GuardTransferCoins: k.HasGuardTransferCoins(ctx)}, nil
}
