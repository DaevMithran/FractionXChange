package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/MANTRA-Chain/mantrachain/x/fractionnft/types"
)

var _ types.QueryServer = Querier{}

type Querier struct {
	Keeper
}

func NewQuerier(keeper Keeper) Querier {
	return Querier{Keeper: keeper}
}

func (k Querier) Params(c context.Context, req *types.QueryParamsRequest) (*types.QueryParamsResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)

	p, err := k.Keeper.Params.Get(ctx)
	if err != nil {
		return nil, err
	}

	return &types.QueryParamsResponse{Params: &p}, nil
}

func (k Querier) GetNftToken(goCtx context.Context, req *types.QueryGetNftTokenRequest) (*types.QueryNftTokenResponse, error) {
	denom := "fractionNFT-"+req.ClassId+"-"+req.NftId
	nftres,  err := k.Keeper.NFTMapping.Get(goCtx, denom)
	if err != nil {
		return nil, err
	}

	return &types.QueryNftTokenResponse{
		Owner: nftres.Owner,
		TimeoutHeight: nftres.TimeoutHeight,
	}, nil

}
