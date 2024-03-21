package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/rewards module sentinel errors
var (
	ErrProviderNotFound     = sdkerrors.Register(ModuleName, 1801, "provider not found")
	ErrProviderPairNotFound = sdkerrors.Register(ModuleName, 1802, "provider pair not found")
	ErrProviderPoolNotFound = sdkerrors.Register(ModuleName, 1803, "provider pool not found")
	ErrPairNotFound         = sdkerrors.Register(ModuleName, 1804, "pair not found")
	ErrInvalidPairId        = sdkerrors.Register(ModuleName, 1805, "invalid pair id")
	ErrSnapshotNotFound     = sdkerrors.Register(ModuleName, 1806, "snapshot not found")
	ErrBalanceMismatch      = sdkerrors.Register(ModuleName, 1807, "balance mismatch")
)
