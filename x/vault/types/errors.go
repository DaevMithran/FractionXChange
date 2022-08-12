package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/vault module sentinel errors
var (
	ErrKeyFormatNotSupported          = sdkerrors.Register(ModuleName, 1111, "key format not supported")
	ErrInvalidMarketplaceId           = sdkerrors.Register(ModuleName, 1112, "marketplace id provided is invalid")
	ErrInvalidCollectionId            = sdkerrors.Register(ModuleName, 1113, "collection id provided is invalid")
	ErrInvalidNftId                   = sdkerrors.Register(ModuleName, 1114, "nft id provided is invalid")
	ErrValidatorDoesNotExist          = sdkerrors.Register(ModuleName, 1115, "validator does not exists")
	ErrInvalidStakingValidatorAddress = sdkerrors.Register(ModuleName, 1116, "staking validator address is invalid")
)