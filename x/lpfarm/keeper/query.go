package keeper

import (
	"github.com/MANTRA-Finance/mantrachain/x/lpfarm/types"
)

var _ types.QueryServer = Keeper{}
