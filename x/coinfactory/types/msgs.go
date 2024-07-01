package types

import (
	"strings"

	"cosmossdk.io/errors"
	"cosmossdk.io/math"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/x/auth/migrations/legacytx"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
)

// constants
const (
	TypeMsgCreateDenom      = "create_denom"
	TypeMsgMint             = "cf_mint"
	TypeMsgBurn             = "cf_burn"
	TypeMsgForceTransfer    = "force_transfer"
	TypeMsgChangeAdmin      = "change_admin"
	TypeMsgSetDenomMetadata = "set_denom_metadata"
)

var (
	_ legacytx.LegacyMsg = &MsgCreateDenom{}
	_ sdk.Msg            = &MsgCreateDenom{}
)

// NewMsgCreateDenom creates a msg to create a new denom
func NewMsgCreateDenom(sender, subdenom string) *MsgCreateDenom {
	return &MsgCreateDenom{
		Sender:   sender,
		Subdenom: subdenom,
	}
}

func (m MsgCreateDenom) Route() string { return RouterKey }
func (m MsgCreateDenom) Type() string  { return TypeMsgCreateDenom }
func (m MsgCreateDenom) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(m.Sender)
	if err != nil {
		return errors.Wrapf(sdkerrors.ErrInvalidAddress, "Invalid sender address (%s)", err)
	}

	_, err = GetTokenDenom(m.Sender, m.Subdenom)
	if err != nil {
		return errors.Wrap(ErrInvalidDenom, err.Error())
	}

	return nil
}

func (msg *MsgCreateDenom) GetSignBytes() []byte {
	bz := Amino.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

var (
	_ legacytx.LegacyMsg = &MsgMint{}
	_ sdk.Msg            = &MsgMint{}
)

// NewMsgMint creates a message to mint tokens
func NewMsgMint(sender string, amount sdk.Coin) *MsgMint {
	return &MsgMint{
		Sender: sender,
		Amount: amount,
	}
}

func NewMsgMintTo(sender string, amount sdk.Coin, mintToAddress string) *MsgMint {
	return &MsgMint{
		Sender:        sender,
		Amount:        amount,
		MintToAddress: mintToAddress,
	}
}

func (m MsgMint) Route() string { return RouterKey }
func (m MsgMint) Type() string  { return TypeMsgMint }
func (m MsgMint) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(m.Sender)
	if err != nil {
		return errors.Wrapf(sdkerrors.ErrInvalidAddress, "Invalid sender address (%s)", err)
	}

	if !m.Amount.IsValid() || m.Amount.Amount.Equal(math.ZeroInt()) {
		return errors.Wrap(sdkerrors.ErrInvalidCoins, m.Amount.String())
	}

	if strings.TrimSpace(m.MintToAddress) != "" {
		_, err = sdk.AccAddressFromBech32(m.MintToAddress)
		if err != nil {
			return errors.Wrapf(sdkerrors.ErrInvalidAddress, "Invalid mint to address (%s)", err)
		}
	}

	return nil
}

func (msg *MsgMint) GetSignBytes() []byte {
	bz := Amino.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

var (
	_ legacytx.LegacyMsg = &MsgBurn{}
	_ sdk.Msg            = &MsgBurn{}
)

// NewMsgBurn creates a message to burn tokens
func NewMsgBurn(sender string, amount sdk.Coin) *MsgBurn {
	return &MsgBurn{
		Sender: sender,
		Amount: amount,
	}
}

func NewMsgBurnFrom(sender string, amount sdk.Coin, burnFromAddress string) *MsgBurn {
	return &MsgBurn{
		Sender:          sender,
		Amount:          amount,
		BurnFromAddress: burnFromAddress,
	}
}

func (m MsgBurn) Route() string { return RouterKey }
func (m MsgBurn) Type() string  { return TypeMsgBurn }
func (m MsgBurn) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(m.Sender)
	if err != nil {
		return errors.Wrapf(sdkerrors.ErrInvalidAddress, "Invalid sender address (%s)", err)
	}

	if !m.Amount.IsValid() || m.Amount.Amount.Equal(math.ZeroInt()) {
		return errors.Wrap(sdkerrors.ErrInvalidCoins, m.Amount.String())
	}

	if strings.TrimSpace(m.BurnFromAddress) != "" {
		_, err = sdk.AccAddressFromBech32(m.BurnFromAddress)
		if err != nil {
			return errors.Wrapf(sdkerrors.ErrInvalidAddress, "Invalid burn from address (%s)", err)
		}
	}

	return nil
}

func (msg *MsgBurn) GetSignBytes() []byte {
	bz := Amino.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

var (
	_ legacytx.LegacyMsg = &MsgForceTransfer{}
	_ sdk.Msg            = &MsgForceTransfer{}
)

// NewMsgForceTransfer creates a transfer funds from one account to another
func NewMsgForceTransfer(sender string, amount sdk.Coin, fromAddr, toAddr string) *MsgForceTransfer {
	return &MsgForceTransfer{
		Sender:              sender,
		Amount:              amount,
		TransferFromAddress: fromAddr,
		TransferToAddress:   toAddr,
	}
}

func (m MsgForceTransfer) Route() string { return RouterKey }
func (m MsgForceTransfer) Type() string  { return TypeMsgForceTransfer }
func (m MsgForceTransfer) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(m.Sender)
	if err != nil {
		return errors.Wrapf(sdkerrors.ErrInvalidAddress, "Invalid sender address (%s)", err)
	}

	_, err = sdk.AccAddressFromBech32(m.TransferFromAddress)
	if err != nil {
		return errors.Wrapf(sdkerrors.ErrInvalidAddress, "Invalid address (%s)", err)
	}
	_, err = sdk.AccAddressFromBech32(m.TransferToAddress)
	if err != nil {
		return errors.Wrapf(sdkerrors.ErrInvalidAddress, "Invalid address (%s)", err)
	}

	if !m.Amount.IsValid() {
		return errors.Wrap(sdkerrors.ErrInvalidCoins, m.Amount.String())
	}

	return nil
}

func (msg *MsgForceTransfer) GetSignBytes() []byte {
	bz := Amino.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

var (
	_ legacytx.LegacyMsg = &MsgChangeAdmin{}
	_ sdk.Msg            = &MsgChangeAdmin{}
)

// NewMsgChangeAdmin creates a message to burn tokens
func NewMsgChangeAdmin(sender, denom, newAdmin string) *MsgChangeAdmin {
	return &MsgChangeAdmin{
		Sender:   sender,
		Denom:    denom,
		NewAdmin: newAdmin,
	}
}

func (m MsgChangeAdmin) Route() string { return RouterKey }
func (m MsgChangeAdmin) Type() string  { return TypeMsgChangeAdmin }
func (m MsgChangeAdmin) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(m.Sender)
	if err != nil {
		return errors.Wrapf(sdkerrors.ErrInvalidAddress, "Invalid sender address (%s)", err)
	}

	_, err = sdk.AccAddressFromBech32(m.NewAdmin)
	if err != nil {
		return errors.Wrapf(sdkerrors.ErrInvalidAddress, "Invalid address (%s)", err)
	}

	_, _, err = DeconstructDenom(m.Denom)
	if err != nil {
		return err
	}

	return nil
}

func (msg *MsgChangeAdmin) GetSignBytes() []byte {
	bz := Amino.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

var (
	_ legacytx.LegacyMsg = &MsgSetDenomMetadata{}
	_ sdk.Msg            = &MsgSetDenomMetadata{}
)

// NewMsgChangeAdmin creates a message to burn tokens
func NewMsgSetDenomMetadata(sender string, metadata banktypes.Metadata) *MsgSetDenomMetadata {
	return &MsgSetDenomMetadata{
		Sender:   sender,
		Metadata: metadata,
	}
}

func (m MsgSetDenomMetadata) Route() string { return RouterKey }
func (m MsgSetDenomMetadata) Type() string  { return TypeMsgSetDenomMetadata }
func (m MsgSetDenomMetadata) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(m.Sender)
	if err != nil {
		return errors.Wrapf(sdkerrors.ErrInvalidAddress, "Invalid sender address (%s)", err)
	}

	err = m.Metadata.Validate()
	if err != nil {
		return err
	}

	_, _, err = DeconstructDenom(m.Metadata.Base)
	if err != nil {
		return err
	}

	return nil
}

func (msg *MsgSetDenomMetadata) GetSignBytes() []byte {
	bz := Amino.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}
