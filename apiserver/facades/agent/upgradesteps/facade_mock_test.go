// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/juju/juju/apiserver/facade (interfaces: Resources,Authorizer,WatcherRegistry)
//
// Generated by this command:
//
//	mockgen -package upgradesteps_test -destination facade_mock_test.go github.com/juju/juju/apiserver/facade Resources,Authorizer,WatcherRegistry
//

// Package upgradesteps_test is a generated GoMock package.
package upgradesteps_test

import (
	reflect "reflect"

	permission "github.com/juju/juju/core/permission"
	names "github.com/juju/names/v5"
	worker "github.com/juju/worker/v3"
	gomock "go.uber.org/mock/gomock"
)

// MockResources is a mock of Resources interface.
type MockResources struct {
	ctrl     *gomock.Controller
	recorder *MockResourcesMockRecorder
}

// MockResourcesMockRecorder is the mock recorder for MockResources.
type MockResourcesMockRecorder struct {
	mock *MockResources
}

// NewMockResources creates a new mock instance.
func NewMockResources(ctrl *gomock.Controller) *MockResources {
	mock := &MockResources{ctrl: ctrl}
	mock.recorder = &MockResourcesMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockResources) EXPECT() *MockResourcesMockRecorder {
	return m.recorder
}

// Get mocks base method.
func (m *MockResources) Get(arg0 string) worker.Worker {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0)
	ret0, _ := ret[0].(worker.Worker)
	return ret0
}

// Get indicates an expected call of Get.
func (mr *MockResourcesMockRecorder) Get(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockResources)(nil).Get), arg0)
}

// Register mocks base method.
func (m *MockResources) Register(arg0 worker.Worker) string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Register", arg0)
	ret0, _ := ret[0].(string)
	return ret0
}

// Register indicates an expected call of Register.
func (mr *MockResourcesMockRecorder) Register(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Register", reflect.TypeOf((*MockResources)(nil).Register), arg0)
}

// Stop mocks base method.
func (m *MockResources) Stop(arg0 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Stop", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Stop indicates an expected call of Stop.
func (mr *MockResourcesMockRecorder) Stop(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Stop", reflect.TypeOf((*MockResources)(nil).Stop), arg0)
}

// MockAuthorizer is a mock of Authorizer interface.
type MockAuthorizer struct {
	ctrl     *gomock.Controller
	recorder *MockAuthorizerMockRecorder
}

// MockAuthorizerMockRecorder is the mock recorder for MockAuthorizer.
type MockAuthorizerMockRecorder struct {
	mock *MockAuthorizer
}

// NewMockAuthorizer creates a new mock instance.
func NewMockAuthorizer(ctrl *gomock.Controller) *MockAuthorizer {
	mock := &MockAuthorizer{ctrl: ctrl}
	mock.recorder = &MockAuthorizerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAuthorizer) EXPECT() *MockAuthorizerMockRecorder {
	return m.recorder
}

// AuthApplicationAgent mocks base method.
func (m *MockAuthorizer) AuthApplicationAgent() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AuthApplicationAgent")
	ret0, _ := ret[0].(bool)
	return ret0
}

// AuthApplicationAgent indicates an expected call of AuthApplicationAgent.
func (mr *MockAuthorizerMockRecorder) AuthApplicationAgent() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AuthApplicationAgent", reflect.TypeOf((*MockAuthorizer)(nil).AuthApplicationAgent))
}

// AuthClient mocks base method.
func (m *MockAuthorizer) AuthClient() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AuthClient")
	ret0, _ := ret[0].(bool)
	return ret0
}

// AuthClient indicates an expected call of AuthClient.
func (mr *MockAuthorizerMockRecorder) AuthClient() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AuthClient", reflect.TypeOf((*MockAuthorizer)(nil).AuthClient))
}

// AuthController mocks base method.
func (m *MockAuthorizer) AuthController() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AuthController")
	ret0, _ := ret[0].(bool)
	return ret0
}

// AuthController indicates an expected call of AuthController.
func (mr *MockAuthorizerMockRecorder) AuthController() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AuthController", reflect.TypeOf((*MockAuthorizer)(nil).AuthController))
}

// AuthMachineAgent mocks base method.
func (m *MockAuthorizer) AuthMachineAgent() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AuthMachineAgent")
	ret0, _ := ret[0].(bool)
	return ret0
}

// AuthMachineAgent indicates an expected call of AuthMachineAgent.
func (mr *MockAuthorizerMockRecorder) AuthMachineAgent() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AuthMachineAgent", reflect.TypeOf((*MockAuthorizer)(nil).AuthMachineAgent))
}

// AuthModelAgent mocks base method.
func (m *MockAuthorizer) AuthModelAgent() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AuthModelAgent")
	ret0, _ := ret[0].(bool)
	return ret0
}

// AuthModelAgent indicates an expected call of AuthModelAgent.
func (mr *MockAuthorizerMockRecorder) AuthModelAgent() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AuthModelAgent", reflect.TypeOf((*MockAuthorizer)(nil).AuthModelAgent))
}

// AuthOwner mocks base method.
func (m *MockAuthorizer) AuthOwner(arg0 names.Tag) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AuthOwner", arg0)
	ret0, _ := ret[0].(bool)
	return ret0
}

// AuthOwner indicates an expected call of AuthOwner.
func (mr *MockAuthorizerMockRecorder) AuthOwner(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AuthOwner", reflect.TypeOf((*MockAuthorizer)(nil).AuthOwner), arg0)
}

// AuthUnitAgent mocks base method.
func (m *MockAuthorizer) AuthUnitAgent() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AuthUnitAgent")
	ret0, _ := ret[0].(bool)
	return ret0
}

// AuthUnitAgent indicates an expected call of AuthUnitAgent.
func (mr *MockAuthorizerMockRecorder) AuthUnitAgent() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AuthUnitAgent", reflect.TypeOf((*MockAuthorizer)(nil).AuthUnitAgent))
}

// ConnectedModel mocks base method.
func (m *MockAuthorizer) ConnectedModel() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ConnectedModel")
	ret0, _ := ret[0].(string)
	return ret0
}

// ConnectedModel indicates an expected call of ConnectedModel.
func (mr *MockAuthorizerMockRecorder) ConnectedModel() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ConnectedModel", reflect.TypeOf((*MockAuthorizer)(nil).ConnectedModel))
}

// EntityHasPermission mocks base method.
func (m *MockAuthorizer) EntityHasPermission(arg0 names.Tag, arg1 permission.Access, arg2 names.Tag) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EntityHasPermission", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// EntityHasPermission indicates an expected call of EntityHasPermission.
func (mr *MockAuthorizerMockRecorder) EntityHasPermission(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EntityHasPermission", reflect.TypeOf((*MockAuthorizer)(nil).EntityHasPermission), arg0, arg1, arg2)
}

// GetAuthTag mocks base method.
func (m *MockAuthorizer) GetAuthTag() names.Tag {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAuthTag")
	ret0, _ := ret[0].(names.Tag)
	return ret0
}

// GetAuthTag indicates an expected call of GetAuthTag.
func (mr *MockAuthorizerMockRecorder) GetAuthTag() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAuthTag", reflect.TypeOf((*MockAuthorizer)(nil).GetAuthTag))
}

// HasPermission mocks base method.
func (m *MockAuthorizer) HasPermission(arg0 permission.Access, arg1 names.Tag) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HasPermission", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// HasPermission indicates an expected call of HasPermission.
func (mr *MockAuthorizerMockRecorder) HasPermission(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HasPermission", reflect.TypeOf((*MockAuthorizer)(nil).HasPermission), arg0, arg1)
}

// MockWatcherRegistry is a mock of WatcherRegistry interface.
type MockWatcherRegistry struct {
	ctrl     *gomock.Controller
	recorder *MockWatcherRegistryMockRecorder
}

// MockWatcherRegistryMockRecorder is the mock recorder for MockWatcherRegistry.
type MockWatcherRegistryMockRecorder struct {
	mock *MockWatcherRegistry
}

// NewMockWatcherRegistry creates a new mock instance.
func NewMockWatcherRegistry(ctrl *gomock.Controller) *MockWatcherRegistry {
	mock := &MockWatcherRegistry{ctrl: ctrl}
	mock.recorder = &MockWatcherRegistryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockWatcherRegistry) EXPECT() *MockWatcherRegistryMockRecorder {
	return m.recorder
}

// Count mocks base method.
func (m *MockWatcherRegistry) Count() int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Count")
	ret0, _ := ret[0].(int)
	return ret0
}

// Count indicates an expected call of Count.
func (mr *MockWatcherRegistryMockRecorder) Count() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Count", reflect.TypeOf((*MockWatcherRegistry)(nil).Count))
}

// Get mocks base method.
func (m *MockWatcherRegistry) Get(arg0 string) (worker.Worker, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0)
	ret0, _ := ret[0].(worker.Worker)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockWatcherRegistryMockRecorder) Get(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockWatcherRegistry)(nil).Get), arg0)
}

// Kill mocks base method.
func (m *MockWatcherRegistry) Kill() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Kill")
}

// Kill indicates an expected call of Kill.
func (mr *MockWatcherRegistryMockRecorder) Kill() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Kill", reflect.TypeOf((*MockWatcherRegistry)(nil).Kill))
}

// Register mocks base method.
func (m *MockWatcherRegistry) Register(arg0 worker.Worker) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Register", arg0)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Register indicates an expected call of Register.
func (mr *MockWatcherRegistryMockRecorder) Register(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Register", reflect.TypeOf((*MockWatcherRegistry)(nil).Register), arg0)
}

// RegisterNamed mocks base method.
func (m *MockWatcherRegistry) RegisterNamed(arg0 string, arg1 worker.Worker) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RegisterNamed", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// RegisterNamed indicates an expected call of RegisterNamed.
func (mr *MockWatcherRegistryMockRecorder) RegisterNamed(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegisterNamed", reflect.TypeOf((*MockWatcherRegistry)(nil).RegisterNamed), arg0, arg1)
}

// Stop mocks base method.
func (m *MockWatcherRegistry) Stop(arg0 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Stop", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Stop indicates an expected call of Stop.
func (mr *MockWatcherRegistryMockRecorder) Stop(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Stop", reflect.TypeOf((*MockWatcherRegistry)(nil).Stop), arg0)
}

// Wait mocks base method.
func (m *MockWatcherRegistry) Wait() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Wait")
	ret0, _ := ret[0].(error)
	return ret0
}

// Wait indicates an expected call of Wait.
func (mr *MockWatcherRegistryMockRecorder) Wait() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Wait", reflect.TypeOf((*MockWatcherRegistry)(nil).Wait))
}
