package keeper

import (
	"context"
	"strconv"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	coinfactorytypes "mantrachain/x/coinfactory/types"
	"mantrachain/x/guard/types"
)

func (k msgServer) UpdateLocked(goCtx context.Context, msg *types.MsgUpdateLocked) (*types.MsgUpdateLockedResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if err := k.CheckIsAdmin(ctx, msg.GetCreator()); err != nil {
		return nil, sdkerrors.Wrap(err, "unauthorized")
	}

	if len(msg.Index) == 0 {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "invalid index")
	}

	kind, err := types.ParseLockedKind(msg.Kind)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "invalid kind")
	}

	denom := string(msg.Index)

	if kind == types.LockedCoin {
		_, _, err := coinfactorytypes.DeconstructDenom(denom)
		if err != nil {
			return nil, err
		}

		_, found := k.bk.GetDenomMetaData(ctx, denom)
		if !found {
			return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "denom not found")
		}
	}

	exists := k.HasLocked(ctx, msg.Index, kind)
	updated := false

	if exists && !msg.Locked {
		k.RemoveLocked(ctx, msg.Index, kind)
		updated = true
	} else if !exists && msg.Locked {
		k.SetLocked(ctx, msg.Index, kind)
		updated = true
	}

	if updated {
		ctx.EventManager().EmitEvent(
			sdk.NewEvent(
				sdk.EventTypeMessage,
				sdk.NewAttribute(sdk.AttributeKeyAction, types.TypeMsgUpdateLocked),
				sdk.NewAttribute(types.AttributeKeyCreator, msg.Creator),
				sdk.NewAttribute(types.AttributeKeyLocked, strconv.FormatBool(msg.Locked)),
				sdk.NewAttribute(types.AttributeKeyIndex, string(msg.Index)),
				sdk.NewAttribute(types.AttributeKeyKind, kind.String()),
			),
		)
	}

	return &types.MsgUpdateLockedResponse{}, nil
}
