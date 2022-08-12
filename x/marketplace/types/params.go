package types

import (
	"fmt"
	"regexp"

	"github.com/cosmos/cosmos-sdk/codec"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"gopkg.in/yaml.v2"
)

const (
	DefaultValidMarketplaceId                           string = "[a-zA-Z0-9_/:-]{0,100}"
	DefaultValidMarketplaceMetadataDescriptionMaxLength uint32 = 1000
	DefaultValidMarketplaceMetadataNameMaxLength        uint32 = 100
	DefaultValidNftsEarningsOnSaleMaxCount              uint32 = 5
	DefaultValidNftsEarningsOnYieldRewardMaxCount       uint32 = 5
)

var (
	KeyValidMarketplaceId                           = []byte("ValidMarketplaceId")
	KeyValidMarketplaceMetadataDescriptionMaxLength = []byte("ValidMarketplaceMetadataDescriptionMaxLength")
	KeyValidMarketplaceMetadataNameMaxLength        = []byte("ValidMarketplaceMetadataNameMaxLength")
	KeyValidNftsEarningsOnSaleMaxCount              = []byte("ValidNftsEarningsOnSaleMaxCount")
	KeyValidNftsEarningsOnYieldRewardMaxCount       = []byte("ValidNftsEarningsOnYieldRewardMaxCount")
)

var _ paramtypes.ParamSet = (*Params)(nil)

// ParamKeyTable the param key table for launch module
func ParamKeyTable() paramtypes.KeyTable {
	return paramtypes.NewKeyTable().RegisterParamSet(&Params{})
}

// NewParams creates a new Params instance
func NewParams(
	validMarketplaceId string,
	validMarketplaceMetadataDescriptionMaxLength uint32,
	validMarketplaceMetadataNameMaxLength uint32,
	validNftsEarningsOnSaleMaxCount uint32,
	validNftsEarningsOnYieldRewardMaxCount uint32,
) Params {
	return Params{
		ValidMarketplaceId:                           validMarketplaceId,
		ValidMarketplaceMetadataDescriptionMaxLength: validMarketplaceMetadataDescriptionMaxLength,
		ValidMarketplaceMetadataNameMaxLength:        validMarketplaceMetadataNameMaxLength,
		ValidNftsEarningsOnSaleMaxCount:              validNftsEarningsOnSaleMaxCount,
		ValidNftsEarningsOnYieldRewardMaxCount:       validNftsEarningsOnYieldRewardMaxCount,
	}
}

// DefaultParams returns a default set of parameters
func DefaultParams() Params {
	return NewParams(
		DefaultValidMarketplaceId,
		DefaultValidMarketplaceMetadataDescriptionMaxLength,
		DefaultValidMarketplaceMetadataNameMaxLength,
		DefaultValidNftsEarningsOnSaleMaxCount,
		DefaultValidNftsEarningsOnYieldRewardMaxCount,
	)
}

// ParamSetPairs get the params.ParamSet
func (p *Params) ParamSetPairs() paramtypes.ParamSetPairs {
	return paramtypes.ParamSetPairs{
		paramtypes.NewParamSetPair(KeyValidMarketplaceId, &p.ValidMarketplaceId, validateValidMarketplaceId),
		paramtypes.NewParamSetPair(KeyValidMarketplaceMetadataDescriptionMaxLength, &p.ValidMarketplaceMetadataDescriptionMaxLength, validateValidMarketplaceMetadataDescriptionMaxLength),
		paramtypes.NewParamSetPair(KeyValidMarketplaceMetadataNameMaxLength, &p.ValidMarketplaceMetadataNameMaxLength, validateValidMarketplaceMetadataNameMaxLength),
		paramtypes.NewParamSetPair(KeyValidNftsEarningsOnSaleMaxCount, &p.ValidNftsEarningsOnSaleMaxCount, validateValidNftsEarningsOnSaleMaxCount),
		paramtypes.NewParamSetPair(KeyValidNftsEarningsOnYieldRewardMaxCount, &p.ValidNftsEarningsOnYieldRewardMaxCount, validateValidNftsEarningsOnYieldRewardMaxCount),
	}
}

// unmarshal the current staking params value from store key or panic
func MustUnmarshalParams(cdc *codec.LegacyAmino, value []byte) Params {
	params, err := UnmarshalParams(cdc, value)
	if err != nil {
		panic(err)
	}

	return params
}

// unmarshal the current staking params value from store key
func UnmarshalParams(cdc *codec.LegacyAmino, value []byte) (params Params, err error) {
	err = cdc.Unmarshal(value, &params)
	if err != nil {
		return
	}

	return
}

// Validate validates the set of params
func (p Params) Validate() error {
	if err := validateValidMarketplaceId(p.ValidMarketplaceId); err != nil {
		return err
	}

	if err := validateValidMarketplaceMetadataDescriptionMaxLength(p.ValidMarketplaceMetadataDescriptionMaxLength); err != nil {
		return err
	}

	if err := validateValidMarketplaceMetadataNameMaxLength(p.ValidMarketplaceMetadataNameMaxLength); err != nil {
		return err
	}

	if err := validateValidNftsEarningsOnSaleMaxCount(p.ValidNftsEarningsOnSaleMaxCount); err != nil {
		return err
	}

	if err := validateValidNftsEarningsOnYieldRewardMaxCount(p.ValidNftsEarningsOnYieldRewardMaxCount); err != nil {
		return err
	}

	return nil
}

// String implements the Stringer interface.
func (p Params) String() string {
	out, _ := yaml.Marshal(p)
	return string(out)
}

func validateValidMarketplaceId(i interface{}) error {
	v, ok := i.(string)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if v == "" {
		return fmt.Errorf("valid marketplace id regex should not be empty")
	}

	_, err := regexp.Compile(v)
	if err != nil {
		return fmt.Errorf("valid marketplace id is invalid regex %s", v)
	}

	return nil
}

func validateValidMarketplaceMetadataDescriptionMaxLength(i interface{}) error {
	v, ok := i.(uint32)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if v == 0 {
		return fmt.Errorf("max validators must be positive: %d", v)
	}

	return nil
}

func validateValidMarketplaceMetadataNameMaxLength(i interface{}) error {
	v, ok := i.(uint32)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if v == 0 {
		return fmt.Errorf("max validators must be positive: %d", v)
	}

	return nil
}

func validateValidNftsEarningsOnSaleMaxCount(i interface{}) error {
	v, ok := i.(uint32)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if v == 0 {
		return fmt.Errorf("max validators must be positive: %d", v)
	}

	return nil
}

func validateValidNftsEarningsOnYieldRewardMaxCount(i interface{}) error {
	v, ok := i.(uint32)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if v == 0 {
		return fmt.Errorf("max validators must be positive: %d", v)
	}

	return nil
}