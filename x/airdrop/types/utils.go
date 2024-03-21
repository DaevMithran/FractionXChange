package types

import (
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/address"
)

const Day = 24 * time.Hour

// ParseTime parses string time to time in RFC3339 format.
// This is used only for internal testing purpose.
func ParseTime(s string) time.Time {
	t, err := time.Parse(time.RFC3339, s)
	if err != nil {
		panic(err)
	}
	return t
}

// DeriveAddress derives an address with the given address length type, module name, and
// address derivation name. It is used to derive private plan farming pool address, and staking reserve address.
func DeriveAddress(moduleName, name string) sdk.AccAddress {
	return sdk.AccAddress(address.Module(moduleName, []byte(name)))
}

// DateRangeIncludes returns true if the target date included on the start, end time range.
// End time is exclusive and start time is inclusive.
func DateRangeIncludes(startTime, endTime, targetTime time.Time) bool {
	return endTime.After(targetTime) && !startTime.After(targetTime)
}
