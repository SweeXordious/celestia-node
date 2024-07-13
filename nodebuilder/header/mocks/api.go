// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/celestiaorg/celestia-node/nodebuilder/header (interfaces: Module)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	header "github.com/celestiaorg/celestia-node/header"
	header0 "github.com/celestiaorg/celestia-node/nodebuilder/header"
	header1 "github.com/celestiaorg/go-header"
	sync "github.com/celestiaorg/go-header/sync"
	gomock "github.com/golang/mock/gomock"
)

// MockModule is a mock of Module interface.
type MockModule struct {
	ctrl     *gomock.Controller
	recorder *MockModuleMockRecorder
}

// MockModuleMockRecorder is the mock recorder for MockModule.
type MockModuleMockRecorder struct {
	mock *MockModule
}

// NewMockModule creates a new mock instance.
func NewMockModule(ctrl *gomock.Controller) *MockModule {
	mock := &MockModule{ctrl: ctrl}
	mock.recorder = &MockModuleMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockModule) EXPECT() *MockModuleMockRecorder {
	return m.recorder
}

// GetByHash mocks base method.
func (m *MockModule) GetByHash(arg0 context.Context, arg1 header1.Hash) (*header.ExtendedHeader, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByHash", arg0, arg1)
	ret0, _ := ret[0].(*header.ExtendedHeader)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByHash indicates an expected call of GetByHash.
func (mr *MockModuleMockRecorder) GetByHash(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByHash", reflect.TypeOf((*MockModule)(nil).GetByHash), arg0, arg1)
}

// GetByHeight mocks base method.
func (m *MockModule) GetByHeight(arg0 context.Context, arg1 uint64) (*header.ExtendedHeader, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByHeight", arg0, arg1)
	ret0, _ := ret[0].(*header.ExtendedHeader)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByHeight indicates an expected call of GetByHeight.
func (mr *MockModuleMockRecorder) GetByHeight(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByHeight", reflect.TypeOf((*MockModule)(nil).GetByHeight), arg0, arg1)
}

// GetDataCommitment mocks base method.
func (m *MockModule) GetDataCommitment(arg0 context.Context, arg1, arg2 uint64) (*header0.DataCommitment, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDataCommitment", arg0, arg1, arg2)
	ret0, _ := ret[0].(*header0.DataCommitment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDataCommitment indicates an expected call of GetDataCommitment.
func (mr *MockModuleMockRecorder) GetDataCommitment(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDataCommitment", reflect.TypeOf((*MockModule)(nil).GetDataCommitment), arg0, arg1, arg2)
}

// GetDataRootInclusionProof mocks base method.
func (m *MockModule) GetDataRootInclusionProof(arg0 context.Context, arg1 int64, arg2, arg3 uint64) (*header0.DataRootTupleInclusionProof, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDataRootInclusionProof", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(*header0.DataRootTupleInclusionProof)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDataRootInclusionProof indicates an expected call of GetDataRootInclusionProof.
func (mr *MockModuleMockRecorder) GetDataRootInclusionProof(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDataRootInclusionProof", reflect.TypeOf((*MockModule)(nil).GetDataRootInclusionProof), arg0, arg1, arg2, arg3)
}

// GetRangeByHeight mocks base method.
func (m *MockModule) GetRangeByHeight(arg0 context.Context, arg1 *header.ExtendedHeader, arg2 uint64) ([]*header.ExtendedHeader, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRangeByHeight", arg0, arg1, arg2)
	ret0, _ := ret[0].([]*header.ExtendedHeader)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRangeByHeight indicates an expected call of GetRangeByHeight.
func (mr *MockModuleMockRecorder) GetRangeByHeight(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRangeByHeight", reflect.TypeOf((*MockModule)(nil).GetRangeByHeight), arg0, arg1, arg2)
}

// LocalHead mocks base method.
func (m *MockModule) LocalHead(arg0 context.Context) (*header.ExtendedHeader, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LocalHead", arg0)
	ret0, _ := ret[0].(*header.ExtendedHeader)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// LocalHead indicates an expected call of LocalHead.
func (mr *MockModuleMockRecorder) LocalHead(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LocalHead", reflect.TypeOf((*MockModule)(nil).LocalHead), arg0)
}

// NetworkHead mocks base method.
func (m *MockModule) NetworkHead(arg0 context.Context) (*header.ExtendedHeader, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NetworkHead", arg0)
	ret0, _ := ret[0].(*header.ExtendedHeader)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NetworkHead indicates an expected call of NetworkHead.
func (mr *MockModuleMockRecorder) NetworkHead(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NetworkHead", reflect.TypeOf((*MockModule)(nil).NetworkHead), arg0)
}

// Subscribe mocks base method.
func (m *MockModule) Subscribe(arg0 context.Context) (<-chan *header.ExtendedHeader, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Subscribe", arg0)
	ret0, _ := ret[0].(<-chan *header.ExtendedHeader)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Subscribe indicates an expected call of Subscribe.
func (mr *MockModuleMockRecorder) Subscribe(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Subscribe", reflect.TypeOf((*MockModule)(nil).Subscribe), arg0)
}

// SyncState mocks base method.
func (m *MockModule) SyncState(arg0 context.Context) (sync.State, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SyncState", arg0)
	ret0, _ := ret[0].(sync.State)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SyncState indicates an expected call of SyncState.
func (mr *MockModuleMockRecorder) SyncState(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SyncState", reflect.TypeOf((*MockModule)(nil).SyncState), arg0)
}

// SyncWait mocks base method.
func (m *MockModule) SyncWait(arg0 context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SyncWait", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SyncWait indicates an expected call of SyncWait.
func (mr *MockModuleMockRecorder) SyncWait(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SyncWait", reflect.TypeOf((*MockModule)(nil).SyncWait), arg0)
}

// WaitForHeight mocks base method.
func (m *MockModule) WaitForHeight(arg0 context.Context, arg1 uint64) (*header.ExtendedHeader, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WaitForHeight", arg0, arg1)
	ret0, _ := ret[0].(*header.ExtendedHeader)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// WaitForHeight indicates an expected call of WaitForHeight.
func (mr *MockModuleMockRecorder) WaitForHeight(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WaitForHeight", reflect.TypeOf((*MockModule)(nil).WaitForHeight), arg0, arg1)
}
