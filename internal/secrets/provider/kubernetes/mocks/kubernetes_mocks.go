// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/juju/juju/internal/secrets/provider/kubernetes (interfaces: Broker)
//
// Generated by this command:
//
//	mockgen -package mocks -destination mocks/kubernetes_mocks.go github.com/juju/juju/internal/secrets/provider/kubernetes Broker
//

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	secrets "github.com/juju/juju/core/secrets"
	names "github.com/juju/names/v5"
	version "github.com/juju/version/v2"
	gomock "go.uber.org/mock/gomock"
)

// MockBroker is a mock of Broker interface.
type MockBroker struct {
	ctrl     *gomock.Controller
	recorder *MockBrokerMockRecorder
}

// MockBrokerMockRecorder is the mock recorder for MockBroker.
type MockBrokerMockRecorder struct {
	mock *MockBroker
}

// NewMockBroker creates a new mock instance.
func NewMockBroker(ctrl *gomock.Controller) *MockBroker {
	mock := &MockBroker{ctrl: ctrl}
	mock.recorder = &MockBrokerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBroker) EXPECT() *MockBrokerMockRecorder {
	return m.recorder
}

// DeleteJujuSecret mocks base method.
func (m *MockBroker) DeleteJujuSecret(arg0 context.Context, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteJujuSecret", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteJujuSecret indicates an expected call of DeleteJujuSecret.
func (mr *MockBrokerMockRecorder) DeleteJujuSecret(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteJujuSecret", reflect.TypeOf((*MockBroker)(nil).DeleteJujuSecret), arg0, arg1)
}

// EnsureSecretAccessToken mocks base method.
func (m *MockBroker) EnsureSecretAccessToken(arg0 names.Tag, arg1, arg2, arg3 []string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EnsureSecretAccessToken", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// EnsureSecretAccessToken indicates an expected call of EnsureSecretAccessToken.
func (mr *MockBrokerMockRecorder) EnsureSecretAccessToken(arg0, arg1, arg2, arg3 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EnsureSecretAccessToken", reflect.TypeOf((*MockBroker)(nil).EnsureSecretAccessToken), arg0, arg1, arg2, arg3)
}

// GetJujuSecret mocks base method.
func (m *MockBroker) GetJujuSecret(arg0 context.Context, arg1 string) (secrets.SecretValue, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetJujuSecret", arg0, arg1)
	ret0, _ := ret[0].(secrets.SecretValue)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetJujuSecret indicates an expected call of GetJujuSecret.
func (mr *MockBrokerMockRecorder) GetJujuSecret(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetJujuSecret", reflect.TypeOf((*MockBroker)(nil).GetJujuSecret), arg0, arg1)
}

// RemoveSecretAccessToken mocks base method.
func (m *MockBroker) RemoveSecretAccessToken(arg0 names.Tag) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveSecretAccessToken", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveSecretAccessToken indicates an expected call of RemoveSecretAccessToken.
func (mr *MockBrokerMockRecorder) RemoveSecretAccessToken(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveSecretAccessToken", reflect.TypeOf((*MockBroker)(nil).RemoveSecretAccessToken), arg0)
}

// SaveJujuSecret mocks base method.
func (m *MockBroker) SaveJujuSecret(arg0 context.Context, arg1 string, arg2 secrets.SecretValue) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SaveJujuSecret", arg0, arg1, arg2)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SaveJujuSecret indicates an expected call of SaveJujuSecret.
func (mr *MockBrokerMockRecorder) SaveJujuSecret(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SaveJujuSecret", reflect.TypeOf((*MockBroker)(nil).SaveJujuSecret), arg0, arg1, arg2)
}

// Version mocks base method.
func (m *MockBroker) Version() (*version.Number, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Version")
	ret0, _ := ret[0].(*version.Number)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Version indicates an expected call of Version.
func (mr *MockBrokerMockRecorder) Version() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Version", reflect.TypeOf((*MockBroker)(nil).Version))
}
