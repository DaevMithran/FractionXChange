package keeper_test

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/MANTRA-Finance/mantrachain/x/liquidity/types"
)

func (s *KeeperTestSuite) TestPairIndexes() {
	pair := s.createPair(s.addr(0), "denom1", "denom2", true)

	pair2, found := s.keeper.GetPairByDenoms(s.ctx, "denom1", "denom2")
	s.Require().True(found)
	s.Require().Equal(pair.Id, pair2.Id)

	resp, err := s.queryClient.Pairs(sdk.WrapSDKContext(s.ctx), &types.QueryPairsRequest{
		Denoms: []string{"denom2"},
	})
	s.Require().NoError(err)
	s.Require().Len(resp.Pairs, 1)
	s.Require().Equal(pair.Id, resp.Pairs[0].Id)
}