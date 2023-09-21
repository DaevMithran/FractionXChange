// Code generated by MockGen. DO NOT EDIT.
// Source: x/txfees/types/expected_keepers.go

// Package testutil is a generated GoMock package.
package testutil

import (
	reflect "reflect"

	types "github.com/MANTRA-Finance/mantrachain/x/txfees/types"
	types0 "github.com/cosmos/cosmos-sdk/types"
	types1 "github.com/cosmos/cosmos-sdk/x/auth/types"
	gomock "github.com/golang/mock/gomock"
)

// MockAccountKeeper is a mock of AccountKeeper interface.
type MockAccountKeeper struct {
	ctrl     *gomock.Controller
	recorder *MockAccountKeeperMockRecorder
}

// MockAccountKeeperMockRecorder is the mock recorder for MockAccountKeeper.
type MockAccountKeeperMockRecorder struct {
	mock *MockAccountKeeper
}

// NewMockAccountKeeper creates a new mock instance.
func NewMockAccountKeeper(ctrl *gomock.Controller) *MockAccountKeeper {
	mock := &MockAccountKeeper{ctrl: ctrl}
	mock.recorder = &MockAccountKeeperMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAccountKeeper) EXPECT() *MockAccountKeeperMockRecorder {
	return m.recorder
}

// GetAccount mocks base method.
func (m *MockAccountKeeper) GetAccount(ctx types0.Context, addr types0.AccAddress) types1.AccountI {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAccount", ctx, addr)
	ret0, _ := ret[0].(types1.AccountI)
	return ret0
}

// GetAccount indicates an expected call of GetAccount.
func (mr *MockAccountKeeperMockRecorder) GetAccount(ctx, addr interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAccount", reflect.TypeOf((*MockAccountKeeper)(nil).GetAccount), ctx, addr)
}

// GetModuleAddress mocks base method.
func (m *MockAccountKeeper) GetModuleAddress(moduleName string) types0.AccAddress {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetModuleAddress", moduleName)
	ret0, _ := ret[0].(types0.AccAddress)
	return ret0
}

// GetModuleAddress indicates an expected call of GetModuleAddress.
func (mr *MockAccountKeeperMockRecorder) GetModuleAddress(moduleName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetModuleAddress", reflect.TypeOf((*MockAccountKeeper)(nil).GetModuleAddress), moduleName)
}

// GetParams mocks base method.
func (m *MockAccountKeeper) GetParams(ctx types0.Context) types1.Params {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetParams", ctx)
	ret0, _ := ret[0].(types1.Params)
	return ret0
}

// GetParams indicates an expected call of GetParams.
func (mr *MockAccountKeeperMockRecorder) GetParams(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetParams", reflect.TypeOf((*MockAccountKeeper)(nil).GetParams), ctx)
}

// SetAccount mocks base method.
func (m *MockAccountKeeper) SetAccount(ctx types0.Context, acc types1.AccountI) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetAccount", ctx, acc)
}

// SetAccount indicates an expected call of SetAccount.
func (mr *MockAccountKeeperMockRecorder) SetAccount(ctx, acc interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetAccount", reflect.TypeOf((*MockAccountKeeper)(nil).SetAccount), ctx, acc)
}

// MockFeegrantKeeper is a mock of FeegrantKeeper interface.
type MockFeegrantKeeper struct {
	ctrl     *gomock.Controller
	recorder *MockFeegrantKeeperMockRecorder
}

// MockFeegrantKeeperMockRecorder is the mock recorder for MockFeegrantKeeper.
type MockFeegrantKeeperMockRecorder struct {
	mock *MockFeegrantKeeper
}

// NewMockFeegrantKeeper creates a new mock instance.
func NewMockFeegrantKeeper(ctrl *gomock.Controller) *MockFeegrantKeeper {
	mock := &MockFeegrantKeeper{ctrl: ctrl}
	mock.recorder = &MockFeegrantKeeperMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockFeegrantKeeper) EXPECT() *MockFeegrantKeeperMockRecorder {
	return m.recorder
}

// UseGrantedFees mocks base method.
func (m *MockFeegrantKeeper) UseGrantedFees(ctx types0.Context, granter, grantee types0.AccAddress, fee types0.Coins, msgs []types0.Msg) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UseGrantedFees", ctx, granter, grantee, fee, msgs)
	ret0, _ := ret[0].(error)
	return ret0
}

// UseGrantedFees indicates an expected call of UseGrantedFees.
func (mr *MockFeegrantKeeperMockRecorder) UseGrantedFees(ctx, granter, grantee, fee, msgs interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UseGrantedFees", reflect.TypeOf((*MockFeegrantKeeper)(nil).UseGrantedFees), ctx, granter, grantee, fee, msgs)
}

// MockBankKeeper is a mock of BankKeeper interface.
type MockBankKeeper struct {
	ctrl     *gomock.Controller
	recorder *MockBankKeeperMockRecorder
}

// MockBankKeeperMockRecorder is the mock recorder for MockBankKeeper.
type MockBankKeeperMockRecorder struct {
	mock *MockBankKeeper
}

// NewMockBankKeeper creates a new mock instance.
func NewMockBankKeeper(ctrl *gomock.Controller) *MockBankKeeper {
	mock := &MockBankKeeper{ctrl: ctrl}
	mock.recorder = &MockBankKeeperMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBankKeeper) EXPECT() *MockBankKeeperMockRecorder {
	return m.recorder
}

// MockGuardKeeper is a mock of GuardKeeper interface.
type MockGuardKeeper struct {
	ctrl     *gomock.Controller
	recorder *MockGuardKeeperMockRecorder
}

// MockGuardKeeperMockRecorder is the mock recorder for MockGuardKeeper.
type MockGuardKeeperMockRecorder struct {
	mock *MockGuardKeeper
}

// NewMockGuardKeeper creates a new mock instance.
func NewMockGuardKeeper(ctrl *gomock.Controller) *MockGuardKeeper {
	mock := &MockGuardKeeper{ctrl: ctrl}
	mock.recorder = &MockGuardKeeperMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockGuardKeeper) EXPECT() *MockGuardKeeperMockRecorder {
	return m.recorder
}

// CheckIsAdmin mocks base method.
func (m *MockGuardKeeper) CheckIsAdmin(ctx types0.Context, address string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckIsAdmin", ctx, address)
	ret0, _ := ret[0].(error)
	return ret0
}

// CheckIsAdmin indicates an expected call of CheckIsAdmin.
func (mr *MockGuardKeeperMockRecorder) CheckIsAdmin(ctx, address interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckIsAdmin", reflect.TypeOf((*MockGuardKeeper)(nil).CheckIsAdmin), ctx, address)
}

// GetAdmin mocks base method.
func (m *MockGuardKeeper) GetAdmin(ctx types0.Context) types0.AccAddress {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAdmin", ctx)
	ret0, _ := ret[0].(types0.AccAddress)
	return ret0
}

// GetAdmin indicates an expected call of GetAdmin.
func (mr *MockGuardKeeperMockRecorder) GetAdmin(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAdmin", reflect.TypeOf((*MockGuardKeeper)(nil).GetAdmin), ctx)
}

// MockLiquidityKeeper is a mock of LiquidityKeeper interface.
type MockLiquidityKeeper struct {
	ctrl     *gomock.Controller
	recorder *MockLiquidityKeeperMockRecorder
}

// MockLiquidityKeeperMockRecorder is the mock recorder for MockLiquidityKeeper.
type MockLiquidityKeeperMockRecorder struct {
	mock *MockLiquidityKeeper
}

// NewMockLiquidityKeeper creates a new mock instance.
func NewMockLiquidityKeeper(ctrl *gomock.Controller) *MockLiquidityKeeper {
	mock := &MockLiquidityKeeper{ctrl: ctrl}
	mock.recorder = &MockLiquidityKeeperMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockLiquidityKeeper) EXPECT() *MockLiquidityKeeperMockRecorder {
	return m.recorder
}

// GetSwapAmount mocks base method.
func (m *MockLiquidityKeeper) GetSwapAmount(ctx types0.Context, pairId uint64, swapCoin types0.Coin) (types0.Coin, types0.Dec, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSwapAmount", ctx, pairId, swapCoin)
	ret0, _ := ret[0].(types0.Coin)
	ret1, _ := ret[1].(types0.Dec)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetSwapAmount indicates an expected call of GetSwapAmount.
func (mr *MockLiquidityKeeperMockRecorder) GetSwapAmount(ctx, pairId, swapCoin interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSwapAmount", reflect.TypeOf((*MockLiquidityKeeper)(nil).GetSwapAmount), ctx, pairId, swapCoin)
}

// MockTxfeesKeeper is a mock of TxfeesKeeper interface.
type MockTxfeesKeeper struct {
	ctrl     *gomock.Controller
	recorder *MockTxfeesKeeperMockRecorder
}

// MockTxfeesKeeperMockRecorder is the mock recorder for MockTxfeesKeeper.
type MockTxfeesKeeperMockRecorder struct {
	mock *MockTxfeesKeeper
}

// NewMockTxfeesKeeper creates a new mock instance.
func NewMockTxfeesKeeper(ctrl *gomock.Controller) *MockTxfeesKeeper {
	mock := &MockTxfeesKeeper{ctrl: ctrl}
	mock.recorder = &MockTxfeesKeeperMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTxfeesKeeper) EXPECT() *MockTxfeesKeeperMockRecorder {
	return m.recorder
}

// GetFeeToken mocks base method.
func (m *MockTxfeesKeeper) GetFeeToken(ctx types0.Context, denom string) (types.FeeToken, bool) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFeeToken", ctx, denom)
	ret0, _ := ret[0].(types.FeeToken)
	ret1, _ := ret[1].(bool)
	return ret0, ret1
}

// GetFeeToken indicates an expected call of GetFeeToken.
func (mr *MockTxfeesKeeperMockRecorder) GetFeeToken(ctx, denom interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFeeToken", reflect.TypeOf((*MockTxfeesKeeper)(nil).GetFeeToken), ctx, denom)
}

// GetParams mocks base method.
func (m *MockTxfeesKeeper) GetParams(ctx types0.Context) types.Params {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetParams", ctx)
	ret0, _ := ret[0].(types.Params)
	return ret0
}

// GetParams indicates an expected call of GetParams.
func (mr *MockTxfeesKeeperMockRecorder) GetParams(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetParams", reflect.TypeOf((*MockTxfeesKeeper)(nil).GetParams), ctx)
}