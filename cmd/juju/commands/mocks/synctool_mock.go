// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/juju/juju/cmd/juju/commands (interfaces: SyncToolAPI)
//
// Generated by this command:
//
//	mockgen -package mocks -destination mocks/synctool_mock.go github.com/juju/juju/cmd/juju/commands SyncToolAPI
//

// Package mocks is a generated GoMock package.
package mocks

import (
	io "io"
	reflect "reflect"

	tools "github.com/juju/juju/internal/tools"
	version "github.com/juju/version/v2"
	gomock "go.uber.org/mock/gomock"
)

// MockSyncToolAPI is a mock of SyncToolAPI interface.
type MockSyncToolAPI struct {
	ctrl     *gomock.Controller
	recorder *MockSyncToolAPIMockRecorder
}

// MockSyncToolAPIMockRecorder is the mock recorder for MockSyncToolAPI.
type MockSyncToolAPIMockRecorder struct {
	mock *MockSyncToolAPI
}

// NewMockSyncToolAPI creates a new mock instance.
func NewMockSyncToolAPI(ctrl *gomock.Controller) *MockSyncToolAPI {
	mock := &MockSyncToolAPI{ctrl: ctrl}
	mock.recorder = &MockSyncToolAPIMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSyncToolAPI) EXPECT() *MockSyncToolAPIMockRecorder {
	return m.recorder
}

// Close mocks base method.
func (m *MockSyncToolAPI) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close.
func (mr *MockSyncToolAPIMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockSyncToolAPI)(nil).Close))
}

// UploadTools mocks base method.
func (m *MockSyncToolAPI) UploadTools(arg0 io.ReadSeeker, arg1 version.Binary) (tools.List, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UploadTools", arg0, arg1)
	ret0, _ := ret[0].(tools.List)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UploadTools indicates an expected call of UploadTools.
func (mr *MockSyncToolAPIMockRecorder) UploadTools(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UploadTools", reflect.TypeOf((*MockSyncToolAPI)(nil).UploadTools), arg0, arg1)
}
