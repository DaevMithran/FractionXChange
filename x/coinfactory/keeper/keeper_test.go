package keeper_test

import (
	"testing"

	cmtproto "github.com/cometbft/cometbft/proto/tendermint/types"
	"github.com/stretchr/testify/suite"

	"cosmossdk.io/math"

	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"golang.org/x/exp/slices"

	"github.com/MANTRA-Finance/mantrachain/app"
	"github.com/MANTRA-Finance/mantrachain/testutil"
	utils "github.com/MANTRA-Finance/mantrachain/types"

	"github.com/MANTRA-Finance/mantrachain/x/coinfactory/keeper"
	"github.com/MANTRA-Finance/mantrachain/x/coinfactory/types"
)

type KeeperTestSuite struct {
	suite.Suite

	app         *app.App
	addrs       []sdk.AccAddress
	ctx         sdk.Context
	keeper      keeper.Keeper
	QueryHelper *baseapp.QueryServiceTestHelper
	queryClient types.QueryClient
	msgServer   types.MsgServer
	// defaultDenom is on the suite, as it depends on the creator test address.
	defaultDenom string
}

var (
	SecondaryDenom  = "ucoin"
	SecondaryAmount = math.NewInt(100000000)
)

func TestKeeperTestSuite(t *testing.T) {
	suite.Run(t, new(KeeperTestSuite))
}

func (suite *KeeperTestSuite) SetupTest() {
	app, err := testutil.Setup()
	suite.NoError(err)

	suite.app = app
	suite.ctx = app.BaseApp.NewContextLegacy(false, cmtproto.Header{
		Height: 1,
		Time:   utils.ParseTime("2022-01-01T00:00:05Z"),
	})
	suite.keeper = suite.app.CoinfactoryKeeper

	initialBalances := sdk.NewCoins(sdk.NewCoin(SecondaryDenom, SecondaryAmount))
	suite.addrs = testutil.AddTestAddrsAndAdmin(suite.app, suite.ctx, 6, math.ZeroInt())
	for _, addr := range suite.addrs {
		err := testutil.FundAccount(suite.app.BankKeeper, suite.ctx, addr, initialBalances)
		suite.Require().NoError(err)
	}
	suite.QueryHelper = &baseapp.QueryServiceTestHelper{
		GRPCQueryRouter: suite.app.GRPCQueryRouter(),
		Ctx:             suite.ctx,
	}

	suite.queryClient = types.NewQueryClient(suite.QueryHelper)
	suite.msgServer = keeper.NewMsgServerImpl(suite.keeper)
}

func (suite *KeeperTestSuite) CreateDefaultDenom() {
	res, err := suite.msgServer.CreateDenom(sdk.WrapSDKContext(suite.ctx), types.NewMsgCreateDenom(suite.addrs[0].String(), "bitcoin"))
	suite.NoError(err)
	suite.defaultDenom = res.GetNewTokenDenom()
}

// AssertEventEmitted asserts that ctx's event manager has emitted the given number of events
// of the given type.
func (s *KeeperTestSuite) AssertEventEmitted(ctx sdk.Context, eventTypeExpected string, numEventsExpected int) {
	allEvents := ctx.EventManager().Events()
	// filter out other events
	actualEvents := make([]sdk.Event, 0)
	for _, event := range allEvents {
		if event.Type == eventTypeExpected {
			actualEvents = append(actualEvents, event)
		}
	}
	s.Equal(numEventsExpected, len(actualEvents))
}

func (s *KeeperTestSuite) FindEvent(events []sdk.Event, name string) sdk.Event {
	index := slices.IndexFunc(events, func(e sdk.Event) bool { return e.Type == name })
	if index == -1 {
		return sdk.Event{}
	}
	return events[index]
}

func (s *KeeperTestSuite) ExtractAttributes(event sdk.Event) map[string]string {
	attrs := make(map[string]string)
	if event.Attributes == nil {
		return attrs
	}
	for _, a := range event.Attributes {
		attrs[string(a.Key)] = string(a.Value)
	}
	return attrs
}
