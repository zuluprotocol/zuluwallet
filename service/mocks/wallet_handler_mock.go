// Code generated by MockGen. DO NOT EDIT.
// Source: code.vegaprotocol.io/go-wallet/service (interfaces: WalletHandler)

// Package mocks is a generated GoMock package.
package mocks

import (
	v1 "code.vegaprotocol.io/go-wallet/internal/proto/commands/v1"
	v10 "code.vegaprotocol.io/go-wallet/internal/proto/wallet/v1"
	wallet "code.vegaprotocol.io/go-wallet/wallet"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockWalletHandler is a mock of WalletHandler interface
type MockWalletHandler struct {
	ctrl     *gomock.Controller
	recorder *MockWalletHandlerMockRecorder
}

// MockWalletHandlerMockRecorder is the mock recorder for MockWalletHandler
type MockWalletHandlerMockRecorder struct {
	mock *MockWalletHandler
}

// NewMockWalletHandler creates a new mock instance
func NewMockWalletHandler(ctrl *gomock.Controller) *MockWalletHandler {
	mock := &MockWalletHandler{ctrl: ctrl}
	mock.recorder = &MockWalletHandlerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockWalletHandler) EXPECT() *MockWalletHandlerMockRecorder {
	return m.recorder
}

// CreateWallet mocks base method
func (m *MockWalletHandler) CreateWallet(arg0, arg1 string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateWallet", arg0, arg1)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateWallet indicates an expected call of CreateWallet
func (mr *MockWalletHandlerMockRecorder) CreateWallet(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateWallet", reflect.TypeOf((*MockWalletHandler)(nil).CreateWallet), arg0, arg1)
}

// GetPublicKey mocks base method
func (m *MockWalletHandler) GetPublicKey(arg0, arg1 string) (wallet.PublicKey, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPublicKey", arg0, arg1)
	ret0, _ := ret[0].(wallet.PublicKey)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPublicKey indicates an expected call of GetPublicKey
func (mr *MockWalletHandlerMockRecorder) GetPublicKey(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPublicKey", reflect.TypeOf((*MockWalletHandler)(nil).GetPublicKey), arg0, arg1)
}

// GetWalletPath mocks base method
func (m *MockWalletHandler) GetWalletPath(arg0 string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetWalletPath", arg0)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetWalletPath indicates an expected call of GetWalletPath
func (mr *MockWalletHandlerMockRecorder) GetWalletPath(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetWalletPath", reflect.TypeOf((*MockWalletHandler)(nil).GetWalletPath), arg0)
}

// ImportWallet mocks base method
func (m *MockWalletHandler) ImportWallet(arg0, arg1, arg2 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ImportWallet", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// ImportWallet indicates an expected call of ImportWallet
func (mr *MockWalletHandlerMockRecorder) ImportWallet(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ImportWallet", reflect.TypeOf((*MockWalletHandler)(nil).ImportWallet), arg0, arg1, arg2)
}

// ListPublicKeys mocks base method
func (m *MockWalletHandler) ListPublicKeys(arg0 string) ([]wallet.PublicKey, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListPublicKeys", arg0)
	ret0, _ := ret[0].([]wallet.PublicKey)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListPublicKeys indicates an expected call of ListPublicKeys
func (mr *MockWalletHandlerMockRecorder) ListPublicKeys(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListPublicKeys", reflect.TypeOf((*MockWalletHandler)(nil).ListPublicKeys), arg0)
}

// LoginWallet mocks base method
func (m *MockWalletHandler) LoginWallet(arg0, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LoginWallet", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// LoginWallet indicates an expected call of LoginWallet
func (mr *MockWalletHandlerMockRecorder) LoginWallet(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LoginWallet", reflect.TypeOf((*MockWalletHandler)(nil).LoginWallet), arg0, arg1)
}

// LogoutWallet mocks base method
func (m *MockWalletHandler) LogoutWallet(arg0 string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "LogoutWallet", arg0)
}

// LogoutWallet indicates an expected call of LogoutWallet
func (mr *MockWalletHandlerMockRecorder) LogoutWallet(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LogoutWallet", reflect.TypeOf((*MockWalletHandler)(nil).LogoutWallet), arg0)
}

// SecureGenerateKeyPair mocks base method
func (m *MockWalletHandler) SecureGenerateKeyPair(arg0, arg1 string, arg2 []wallet.Meta) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SecureGenerateKeyPair", arg0, arg1, arg2)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SecureGenerateKeyPair indicates an expected call of SecureGenerateKeyPair
func (mr *MockWalletHandlerMockRecorder) SecureGenerateKeyPair(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SecureGenerateKeyPair", reflect.TypeOf((*MockWalletHandler)(nil).SecureGenerateKeyPair), arg0, arg1, arg2)
}

// SignAny mocks base method
func (m *MockWalletHandler) SignAny(arg0 string, arg1 []byte, arg2 string) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SignAny", arg0, arg1, arg2)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SignAny indicates an expected call of SignAny
func (mr *MockWalletHandlerMockRecorder) SignAny(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SignAny", reflect.TypeOf((*MockWalletHandler)(nil).SignAny), arg0, arg1, arg2)
}

// SignTx mocks base method
func (m *MockWalletHandler) SignTx(arg0, arg1, arg2 string, arg3 uint64) (wallet.SignedBundle, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SignTx", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(wallet.SignedBundle)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SignTx indicates an expected call of SignTx
func (mr *MockWalletHandlerMockRecorder) SignTx(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SignTx", reflect.TypeOf((*MockWalletHandler)(nil).SignTx), arg0, arg1, arg2, arg3)
}

// SignTxV2 mocks base method
func (m *MockWalletHandler) SignTxV2(arg0 string, arg1 *v10.SubmitTransactionRequest, arg2 uint64) (*v1.Transaction, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SignTxV2", arg0, arg1, arg2)
	ret0, _ := ret[0].(*v1.Transaction)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SignTxV2 indicates an expected call of SignTxV2
func (mr *MockWalletHandlerMockRecorder) SignTxV2(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SignTxV2", reflect.TypeOf((*MockWalletHandler)(nil).SignTxV2), arg0, arg1, arg2)
}

// TaintKey mocks base method
func (m *MockWalletHandler) TaintKey(arg0, arg1, arg2 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TaintKey", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// TaintKey indicates an expected call of TaintKey
func (mr *MockWalletHandlerMockRecorder) TaintKey(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TaintKey", reflect.TypeOf((*MockWalletHandler)(nil).TaintKey), arg0, arg1, arg2)
}

// UpdateMeta mocks base method
func (m *MockWalletHandler) UpdateMeta(arg0, arg1, arg2 string, arg3 []wallet.Meta) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateMeta", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateMeta indicates an expected call of UpdateMeta
func (mr *MockWalletHandlerMockRecorder) UpdateMeta(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateMeta", reflect.TypeOf((*MockWalletHandler)(nil).UpdateMeta), arg0, arg1, arg2, arg3)
}

// VerifyAny mocks base method
func (m *MockWalletHandler) VerifyAny(arg0 string, arg1, arg2 []byte, arg3 string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "VerifyAny", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// VerifyAny indicates an expected call of VerifyAny
func (mr *MockWalletHandlerMockRecorder) VerifyAny(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "VerifyAny", reflect.TypeOf((*MockWalletHandler)(nil).VerifyAny), arg0, arg1, arg2, arg3)
}
