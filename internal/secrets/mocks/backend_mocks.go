// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/juju/juju/internal/secrets/provider (interfaces: SecretsBackend)
//
// Generated by this command:
//
//	mockgen -package mocks -destination mocks/backend_mocks.go github.com/juju/juju/internal/secrets/provider SecretsBackend
//

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	secrets "github.com/juju/juju/core/secrets"
	gomock "go.uber.org/mock/gomock"
)

// MockSecretsBackend is a mock of SecretsBackend interface.
type MockSecretsBackend struct {
	ctrl     *gomock.Controller
	recorder *MockSecretsBackendMockRecorder
}

// MockSecretsBackendMockRecorder is the mock recorder for MockSecretsBackend.
type MockSecretsBackendMockRecorder struct {
	mock *MockSecretsBackend
}

// NewMockSecretsBackend creates a new mock instance.
func NewMockSecretsBackend(ctrl *gomock.Controller) *MockSecretsBackend {
	mock := &MockSecretsBackend{ctrl: ctrl}
	mock.recorder = &MockSecretsBackendMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSecretsBackend) EXPECT() *MockSecretsBackendMockRecorder {
	return m.recorder
}

// DeleteContent mocks base method.
func (m *MockSecretsBackend) DeleteContent(arg0 context.Context, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteContent", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteContent indicates an expected call of DeleteContent.
func (mr *MockSecretsBackendMockRecorder) DeleteContent(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteContent", reflect.TypeOf((*MockSecretsBackend)(nil).DeleteContent), arg0, arg1)
}

// GetContent mocks base method.
func (m *MockSecretsBackend) GetContent(arg0 context.Context, arg1 string) (secrets.SecretValue, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetContent", arg0, arg1)
	ret0, _ := ret[0].(secrets.SecretValue)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetContent indicates an expected call of GetContent.
func (mr *MockSecretsBackendMockRecorder) GetContent(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetContent", reflect.TypeOf((*MockSecretsBackend)(nil).GetContent), arg0, arg1)
}

// Ping mocks base method.
func (m *MockSecretsBackend) Ping() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Ping")
	ret0, _ := ret[0].(error)
	return ret0
}

// Ping indicates an expected call of Ping.
func (mr *MockSecretsBackendMockRecorder) Ping() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Ping", reflect.TypeOf((*MockSecretsBackend)(nil).Ping))
}

// SaveContent mocks base method.
func (m *MockSecretsBackend) SaveContent(arg0 context.Context, arg1 *secrets.URI, arg2 int, arg3 secrets.SecretValue) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SaveContent", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SaveContent indicates an expected call of SaveContent.
func (mr *MockSecretsBackendMockRecorder) SaveContent(arg0, arg1, arg2, arg3 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SaveContent", reflect.TypeOf((*MockSecretsBackend)(nil).SaveContent), arg0, arg1, arg2, arg3)
}
