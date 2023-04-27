// Code generated by MockGen. DO NOT EDIT.
// Source: x/token/types/expected_keepers.go

// Package testutil is a generated GoMock package.
package testutil

import (
	reflect "reflect"

	types "github.com/MANTRA-Finance/mantrachain/x/nft/types"
	types0 "github.com/cosmos/cosmos-sdk/types"
	gomock "github.com/golang/mock/gomock"
)

// MockNFTKeeper is a mock of NFTKeeper interface.
type MockNFTKeeper struct {
	ctrl     *gomock.Controller
	recorder *MockNFTKeeperMockRecorder
}

// MockNFTKeeperMockRecorder is the mock recorder for MockNFTKeeper.
type MockNFTKeeperMockRecorder struct {
	mock *MockNFTKeeper
}

// NewMockNFTKeeper creates a new mock instance.
func NewMockNFTKeeper(ctrl *gomock.Controller) *MockNFTKeeper {
	mock := &MockNFTKeeper{ctrl: ctrl}
	mock.recorder = &MockNFTKeeperMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockNFTKeeper) EXPECT() *MockNFTKeeperMockRecorder {
	return m.recorder
}

// BatchBurn mocks base method.
func (m *MockNFTKeeper) BatchBurn(ctx types0.Context, classID string, nftIDs []string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BatchBurn", ctx, classID, nftIDs)
	ret0, _ := ret[0].(error)
	return ret0
}

// BatchBurn indicates an expected call of BatchBurn.
func (mr *MockNFTKeeperMockRecorder) BatchBurn(ctx, classID, nftIDs interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BatchBurn", reflect.TypeOf((*MockNFTKeeper)(nil).BatchBurn), ctx, classID, nftIDs)
}

// BatchMint mocks base method.
func (m *MockNFTKeeper) BatchMint(ctx types0.Context, tokens []types.NFT, receiver types0.AccAddress) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BatchMint", ctx, tokens, receiver)
	ret0, _ := ret[0].(error)
	return ret0
}

// BatchMint indicates an expected call of BatchMint.
func (mr *MockNFTKeeperMockRecorder) BatchMint(ctx, tokens, receiver interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BatchMint", reflect.TypeOf((*MockNFTKeeper)(nil).BatchMint), ctx, tokens, receiver)
}

// BatchTransfer mocks base method.
func (m *MockNFTKeeper) BatchTransfer(ctx types0.Context, classID string, nftIDs []string, receiver types0.AccAddress) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BatchTransfer", ctx, classID, nftIDs, receiver)
	ret0, _ := ret[0].(error)
	return ret0
}

// BatchTransfer indicates an expected call of BatchTransfer.
func (mr *MockNFTKeeperMockRecorder) BatchTransfer(ctx, classID, nftIDs, receiver interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BatchTransfer", reflect.TypeOf((*MockNFTKeeper)(nil).BatchTransfer), ctx, classID, nftIDs, receiver)
}

// Burn mocks base method.
func (m *MockNFTKeeper) Burn(ctx types0.Context, classID, nftID string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Burn", ctx, classID, nftID)
	ret0, _ := ret[0].(error)
	return ret0
}

// Burn indicates an expected call of Burn.
func (mr *MockNFTKeeperMockRecorder) Burn(ctx, classID, nftID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Burn", reflect.TypeOf((*MockNFTKeeper)(nil).Burn), ctx, classID, nftID)
}

// GetNFT mocks base method.
func (m *MockNFTKeeper) GetNFT(ctx types0.Context, classID, nftID string) (types.NFT, bool) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetNFT", ctx, classID, nftID)
	ret0, _ := ret[0].(types.NFT)
	ret1, _ := ret[1].(bool)
	return ret0, ret1
}

// GetNFT indicates an expected call of GetNFT.
func (mr *MockNFTKeeperMockRecorder) GetNFT(ctx, classID, nftID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetNFT", reflect.TypeOf((*MockNFTKeeper)(nil).GetNFT), ctx, classID, nftID)
}

// GetOwner mocks base method.
func (m *MockNFTKeeper) GetOwner(ctx types0.Context, classID, nftID string) types0.AccAddress {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOwner", ctx, classID, nftID)
	ret0, _ := ret[0].(types0.AccAddress)
	return ret0
}

// GetOwner indicates an expected call of GetOwner.
func (mr *MockNFTKeeperMockRecorder) GetOwner(ctx, classID, nftID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOwner", reflect.TypeOf((*MockNFTKeeper)(nil).GetOwner), ctx, classID, nftID)
}

// Mint mocks base method.
func (m *MockNFTKeeper) Mint(ctx types0.Context, token types.NFT, receiver types0.AccAddress) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Mint", ctx, token, receiver)
	ret0, _ := ret[0].(error)
	return ret0
}

// Mint indicates an expected call of Mint.
func (mr *MockNFTKeeperMockRecorder) Mint(ctx, token, receiver interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Mint", reflect.TypeOf((*MockNFTKeeper)(nil).Mint), ctx, token, receiver)
}

// SaveClass mocks base method.
func (m *MockNFTKeeper) SaveClass(ctx types0.Context, class types.Class) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SaveClass", ctx, class)
	ret0, _ := ret[0].(error)
	return ret0
}

// SaveClass indicates an expected call of SaveClass.
func (mr *MockNFTKeeperMockRecorder) SaveClass(ctx, class interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SaveClass", reflect.TypeOf((*MockNFTKeeper)(nil).SaveClass), ctx, class)
}

// Transfer mocks base method.
func (m *MockNFTKeeper) Transfer(ctx types0.Context, classID, nftID string, receiver types0.AccAddress) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Transfer", ctx, classID, nftID, receiver)
	ret0, _ := ret[0].(error)
	return ret0
}

// Transfer indicates an expected call of Transfer.
func (mr *MockNFTKeeperMockRecorder) Transfer(ctx, classID, nftID, receiver interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Transfer", reflect.TypeOf((*MockNFTKeeper)(nil).Transfer), ctx, classID, nftID, receiver)
}
