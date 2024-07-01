package types

import (
	"time"

	"cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	errorstypes "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/x/auth/migrations/legacytx"
)

const (
	TypeMsgCreateCampaign  = "create_campaign"
	TypeMsgDeleteCampaign  = "delete_campaign"
	TypeMsgPauseCampaign   = "pause_campaign"
	TypeMsgUnpauseCampaign = "unpause_campaign"
	TypeMsgCampaignClaim   = "campaign_claim"
)

var (
	_ legacytx.LegacyMsg = &MsgCreateCampaign{}
	_ legacytx.LegacyMsg = &MsgDeleteCampaign{}
	_ legacytx.LegacyMsg = &MsgPauseCampaign{}
	_ legacytx.LegacyMsg = &MsgUnpauseCampaign{}
	_ legacytx.LegacyMsg = &MsgCampaignClaim{}
	_ sdk.Msg            = &MsgCreateCampaign{}
	_ sdk.Msg            = &MsgDeleteCampaign{}
	_ sdk.Msg            = &MsgPauseCampaign{}
	_ sdk.Msg            = &MsgUnpauseCampaign{}
	_ sdk.Msg            = &MsgCampaignClaim{}
)

func NewMsgCreateCampaign(creator string, name string, desc string, startTime time.Time, endTime time.Time, mtRoot []byte, amount sdk.Coin) *MsgCreateCampaign {
	return &MsgCreateCampaign{
		Creator:   creator,
		Name:      name,
		Desc:      desc,
		StartTime: startTime,
		EndTime:   endTime,
		MtRoot:    mtRoot,
		Amount:    amount,
	}
}

func (msg *MsgCreateCampaign) Route() string {
	return RouterKey
}

func (msg *MsgCreateCampaign) Type() string {
	return TypeMsgCreateCampaign
}

func (msg *MsgCreateCampaign) GetSignBytes() []byte {
	bz := Amino.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateCampaign) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errors.Wrapf(errorstypes.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	if len(msg.Name) > MaxCampaignNameLen {
		return errors.Wrapf(ErrCampaignNameTooLong, "too long campaign name, maximum %d", MaxCampaignNameLen)
	}
	if len(msg.Desc) > MaxCampaignDescriptionLen {
		return errors.Wrapf(ErrCampaignDescTooLong, "too long campaign description, maximum %d", MaxCampaignDescriptionLen)
	}
	if len(msg.MtRoot) != 32 {
		return errors.Wrap(ErrCampaignMtRootInvalid, "merkle tree root hash must be 32 bytes")
	}
	if msg.Amount.IsZero() || !msg.Amount.IsValid() {
		return errors.Wrapf(ErrCampaignInvalidAmount, "invalid amount: %s", msg.Amount)
	}
	if !msg.StartTime.Before(msg.EndTime) {
		return errors.Wrap(ErrCampaignStartTimeInvalid, "end time must be after start time")
	}
	return nil
}

func NewMsgDeleteCampaign(creator string, id uint64) *MsgDeleteCampaign {
	return &MsgDeleteCampaign{
		Creator: creator,
		Id:      id,
	}
}

func (msg *MsgDeleteCampaign) Route() string {
	return RouterKey
}

func (msg *MsgDeleteCampaign) Type() string {
	return TypeMsgDeleteCampaign
}

func (msg *MsgDeleteCampaign) GetSignBytes() []byte {
	bz := Amino.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgDeleteCampaign) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errors.Wrapf(errorstypes.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	if msg.Id == 0 {
		return errors.Wrap(ErrCampaignInvalidId, "campaign id cannot be 0")
	}
	return nil
}

func NewMsgPauseCampaign(creator string, id uint64) *MsgPauseCampaign {
	return &MsgPauseCampaign{
		Creator: creator,
		Id:      id,
	}
}

func (msg *MsgPauseCampaign) Route() string {
	return RouterKey
}

func (msg *MsgPauseCampaign) Type() string {
	return TypeMsgPauseCampaign
}

func (msg *MsgPauseCampaign) GetSignBytes() []byte {
	bz := Amino.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgPauseCampaign) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errors.Wrapf(errorstypes.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	if msg.Id == 0 {
		return errors.Wrap(ErrCampaignInvalidId, "campaign id cannot be 0")
	}
	return nil
}

func NewMsgUnpauseCampaign(creator string, id uint64) *MsgUnpauseCampaign {
	return &MsgUnpauseCampaign{
		Creator: creator,
		Id:      id,
	}
}

func (msg *MsgUnpauseCampaign) Route() string {
	return RouterKey
}

func (msg *MsgUnpauseCampaign) Type() string {
	return TypeMsgUnpauseCampaign
}

func (msg *MsgUnpauseCampaign) GetSignBytes() []byte {
	bz := Amino.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUnpauseCampaign) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errors.Wrapf(errorstypes.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	if msg.Id == 0 {
		return errors.Wrap(ErrCampaignInvalidId, "campaign id cannot be 0")
	}
	return nil
}

func NewMsgCampaignClaim(creator string, id uint64, amount sdk.Coin, mip []byte, index uint64) *MsgCampaignClaim {
	return &MsgCampaignClaim{
		Creator: creator,
		Id:      id,
		Amount:  &amount,
		Mip:     mip,
		Index:   index,
	}
}

func (msg *MsgCampaignClaim) Route() string {
	return RouterKey
}

func (msg *MsgCampaignClaim) Type() string {
	return TypeMsgCampaignClaim
}

func (msg *MsgCampaignClaim) GetSignBytes() []byte {
	bz := Amino.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCampaignClaim) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errors.Wrapf(errorstypes.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	if len(msg.Mip) == 0 {
		return errors.Wrap(ErrInvalidMerklePath, "empty merkle path")
	}
	if len(msg.Mip)%32 != 0 {
		return errors.Wrap(ErrInvalidMerklePath, "invalid merkle path length")
	}
	if msg.Amount == nil || msg.Amount.IsZero() || !msg.Amount.IsValid() {
		return errors.Wrapf(ErrCampaignInvalidAmount, "invalid amount: %s", msg.Amount)
	}
	if msg.Id == 0 {
		return errors.Wrap(ErrCampaignInvalidId, "campaign id cannot be 0")
	}
	if msg.Index == 0 {
		return errors.Wrap(ErrInvalidMerklePathIndex, "invalid merkle path index")
	}
	return nil
}
