package ante

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
)

type GuardTransferDecorator struct {
	guardKeeper GuardKeeper
	tokenKeeper TokenKeeper
	nftKeeper   NFTKeeper
}

func NewGuardTransferDecorator(gk GuardKeeper, tk TokenKeeper, nk NFTKeeper) GuardTransferDecorator {
	return GuardTransferDecorator{guardKeeper: gk, tokenKeeper: tk, nftKeeper: nk}
}

func (gtd GuardTransferDecorator) AnteHandle(ctx sdk.Context, tx sdk.Tx, simulate bool, next sdk.AnteHandler) (newCtx sdk.Context, err error) {
	switch tx.(type) {
	case sdk.Tx:
		for _, msg := range tx.GetMsgs() {
			switch msg := msg.(type) {
			case *banktypes.MsgSend:
				from, err := sdk.AccAddressFromBech32(msg.FromAddress)

				if err != nil {
					return ctx, sdkerrors.ErrInvalidAddress.Wrapf("invalid from address: %s", err)
				}

				to, err := sdk.AccAddressFromBech32(msg.ToAddress)

				if err != nil {
					return ctx, sdkerrors.ErrInvalidAddress.Wrapf("invalid to address: %s", err)
				}

				ok, err := gtd.guardKeeper.CheckCanTransfer(ctx, gtd.tokenKeeper, gtd.nftKeeper, []sdk.AccAddress{from, to}, msg.Amount)

				if err != nil {
					return ctx, errorsmod.Wrap(err, "send guard: fail")
				}

				if !ok {
					return ctx, errorsmod.Wrap(err, "send guard: cannot transfer")
				}
			}
		}
	}

	return next(ctx, tx, simulate)
}