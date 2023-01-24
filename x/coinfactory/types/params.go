package types

import (
	"fmt"

	cmdcfg "github.com/LimeChain/mantrachain/cmd/config"

	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
)

// Parameter store keys.
var (
	KeyDenomCreationFee = []byte("DenomCreationFee")
)

// ParamTable for gamm module.
func ParamKeyTable() paramtypes.KeyTable {
	return paramtypes.NewKeyTable().RegisterParamSet(&Params{})
}

func NewParams(denomCreationFee sdk.Coins) Params {
	return Params{
		DenomCreationFee: denomCreationFee,
	}
}

// default gamm module parameters.
func DefaultParams() Params {
	return Params{
		// TODO: this should be adjusted
		DenomCreationFee: sdk.NewCoins(sdk.NewCoin(cmdcfg.BaseDenom, sdk.NewInt(1000000000000000000))), // 1 XOM
	}
}

// validate params.
func (p Params) Validate() error {
	if err := validateDenomCreationFee(p.DenomCreationFee); err != nil {
		return err
	}

	return nil
}

// Implements params.ParamSet.
func (p *Params) ParamSetPairs() paramtypes.ParamSetPairs {
	return paramtypes.ParamSetPairs{
		paramtypes.NewParamSetPair(KeyDenomCreationFee, &p.DenomCreationFee, validateDenomCreationFee),
	}
}

func validateDenomCreationFee(i interface{}) error {
	v, ok := i.(sdk.Coins)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if v.Validate() != nil {
		return fmt.Errorf("invalid denom creation fee: %+v", i)
	}

	return nil
}