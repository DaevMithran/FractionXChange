package decorators

import (
	"fmt"
	"strings"

	fractionnfttypes "github.com/MANTRA-Chain/mantrachain/x/fractionnft/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/authz"
	ibctransfertypes "github.com/cosmos/ibc-go/v8/modules/apps/transfer/types"
)

// MsgFilterDecorator is an ante.go decorator template for filtering messages.
type MsgFilterDecorator struct {
	blockedTypes []sdk.Msg
}

// FilterDecorator returns a new MsgFilterDecorator. This errors if the transaction
// contains any of the blocked message types.
//
// Example:
// - decorators.FilterDecorator(&banktypes.MsgSend{})
// This would block any MsgSend messages from being included in a transaction if set in ante.go
func FilterDecorator() MsgFilterDecorator {
	return MsgFilterDecorator{}
}

func (mfd MsgFilterDecorator) AnteHandle(ctx sdk.Context, tx sdk.Tx, simulate bool, next sdk.AnteHandler) (newCtx sdk.Context, err error) {
	if mfd.HasDisallowedMessage(ctx, tx.GetMsgs()) {
		currHeight := ctx.BlockHeight()
		return ctx, fmt.Errorf("tx contains unsupported message types at height %d", currHeight)
	}

	return next(ctx, tx, simulate)
}

func (mfd MsgFilterDecorator) HasDisallowedMessage(ctx sdk.Context, msgs []sdk.Msg) bool {
	for _, msg := range msgs {
		// check nested messages in a recursive manner
		if execMsg, ok := msg.(*authz.MsgExec); ok {
			msgs, err := execMsg.GetMessages()
			if err != nil {
				return true
			}

			if mfd.HasDisallowedMessage(ctx, msgs) {
				return true
			}
		}
		if transferMsg, ok := msg.(*ibctransfertypes.MsgTransfer); ok {
			if strings.HasPrefix(transferMsg.Token.Denom, fractionnfttypes.ModuleName) {
				return true
			}
		}

	}

	return false
}
