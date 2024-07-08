package types

import (
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
)

const (
	CampaignReserveAddressPrefix = "CampaignReserveAddress"
	MaxCampaignNameLen           = 100 // Maximum length of a campaign's name
	MaxCampaignDescriptionLen    = 200 // Maximum length of a campaign's description
	ModuleAddressNameSplitter    = "|"
)

var _ paramtypes.ParamSet = (*Params)(nil)

// ParamKeyTable the param key table for launch module
func ParamKeyTable() paramtypes.KeyTable {
	return paramtypes.NewKeyTable().RegisterParamSet(&Params{})
}

// NewParams creates a new Params instance
func NewParams() Params {
	return Params{}
}

// DefaultParams returns a default set of parameters
func DefaultParams() Params {
	return NewParams()
}

// ParamSetPairs get the params.ParamSet
func (p *Params) ParamSetPairs() paramtypes.ParamSetPairs {
	return paramtypes.ParamSetPairs{}
}

// Validate validates the set of params
func (p Params) Validate() error {
	return nil
}