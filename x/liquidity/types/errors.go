package types

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// DONTCOVER

// x/liquidity module sentinel errors
var (
	ErrInsufficientDepositAmount = sdkerrors.Register(ModuleName, 2, "insufficient deposit amount")
	ErrPairAlreadyExists         = sdkerrors.Register(ModuleName, 3, "pair already exists")
	ErrPoolAlreadyExists         = sdkerrors.Register(ModuleName, 4, "pool already exists")
	ErrWrongPoolCoinDenom        = sdkerrors.Register(ModuleName, 5, "wrong pool coin denom")
	ErrInvalidCoinDenom          = sdkerrors.Register(ModuleName, 6, "invalid coin denom")
	ErrNoLastPrice               = sdkerrors.Register(ModuleName, 7, "cannot make a market order to a pair with no last price")
	ErrInsufficientOfferCoin     = sdkerrors.Register(ModuleName, 8, "insufficient offer coin")
	ErrPriceOutOfRange           = sdkerrors.Register(ModuleName, 9, "price out of range limit")
	ErrTooLongOrderLifespan      = sdkerrors.Register(ModuleName, 10, "order lifespan is too long")
	ErrDisabledPool              = sdkerrors.Register(ModuleName, 11, "disabled pool")
	ErrWrongPair                 = sdkerrors.Register(ModuleName, 12, "wrong denom pair")
	ErrSameBatch                 = sdkerrors.Register(ModuleName, 13, "cannot cancel an order within the same batch")
	ErrAlreadyCanceled           = sdkerrors.Register(ModuleName, 14, "the order is already canceled")
	ErrDuplicatePairId           = sdkerrors.Register(ModuleName, 15, "duplicate pair id presents in the pair id list")
	ErrTooSmallOrder             = sdkerrors.Register(ModuleName, 16, "too small order")
	ErrTooLargePool              = sdkerrors.Register(ModuleName, 17, "too large pool")
	ErrTooManyPools              = sdkerrors.Register(ModuleName, 18, "too many pools in the pair")
	ErrPriceNotOnTicks           = sdkerrors.Register(ModuleName, 19, "price is not on ticks")
)
