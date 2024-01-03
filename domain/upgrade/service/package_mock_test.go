// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/juju/juju/domain/upgrade/service (interfaces: State,WatcherFactory)
//
// Generated by this command:
//
//	mockgen -package service -destination package_mock_test.go github.com/juju/juju/domain/upgrade/service State,WatcherFactory
//

// Package service is a generated GoMock package.
package service

import (
	context "context"
	reflect "reflect"

	changestream "github.com/juju/juju/core/changestream"
	upgrade "github.com/juju/juju/core/upgrade"
	watcher "github.com/juju/juju/core/watcher"
	eventsource "github.com/juju/juju/core/watcher/eventsource"
	upgrade0 "github.com/juju/juju/domain/upgrade"
	version "github.com/juju/version/v2"
	gomock "go.uber.org/mock/gomock"
)

// MockState is a mock of State interface.
type MockState struct {
	ctrl     *gomock.Controller
	recorder *MockStateMockRecorder
}

// MockStateMockRecorder is the mock recorder for MockState.
type MockStateMockRecorder struct {
	mock *MockState
}

// NewMockState creates a new mock instance.
func NewMockState(ctrl *gomock.Controller) *MockState {
	mock := &MockState{ctrl: ctrl}
	mock.recorder = &MockStateMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockState) EXPECT() *MockStateMockRecorder {
	return m.recorder
}

// ActiveUpgrade mocks base method.
func (m *MockState) ActiveUpgrade(arg0 context.Context) (upgrade0.UUID, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ActiveUpgrade", arg0)
	ret0, _ := ret[0].(upgrade0.UUID)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ActiveUpgrade indicates an expected call of ActiveUpgrade.
func (mr *MockStateMockRecorder) ActiveUpgrade(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ActiveUpgrade", reflect.TypeOf((*MockState)(nil).ActiveUpgrade), arg0)
}

// AllProvisionedControllersReady mocks base method.
func (m *MockState) AllProvisionedControllersReady(arg0 context.Context, arg1 upgrade0.UUID) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AllProvisionedControllersReady", arg0, arg1)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AllProvisionedControllersReady indicates an expected call of AllProvisionedControllersReady.
func (mr *MockStateMockRecorder) AllProvisionedControllersReady(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AllProvisionedControllersReady", reflect.TypeOf((*MockState)(nil).AllProvisionedControllersReady), arg0, arg1)
}

// CreateUpgrade mocks base method.
func (m *MockState) CreateUpgrade(arg0 context.Context, arg1, arg2 version.Number) (upgrade0.UUID, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUpgrade", arg0, arg1, arg2)
	ret0, _ := ret[0].(upgrade0.UUID)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateUpgrade indicates an expected call of CreateUpgrade.
func (mr *MockStateMockRecorder) CreateUpgrade(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUpgrade", reflect.TypeOf((*MockState)(nil).CreateUpgrade), arg0, arg1, arg2)
}

// SetControllerDone mocks base method.
func (m *MockState) SetControllerDone(arg0 context.Context, arg1 upgrade0.UUID, arg2 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetControllerDone", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetControllerDone indicates an expected call of SetControllerDone.
func (mr *MockStateMockRecorder) SetControllerDone(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetControllerDone", reflect.TypeOf((*MockState)(nil).SetControllerDone), arg0, arg1, arg2)
}

// SetControllerReady mocks base method.
func (m *MockState) SetControllerReady(arg0 context.Context, arg1 upgrade0.UUID, arg2 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetControllerReady", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetControllerReady indicates an expected call of SetControllerReady.
func (mr *MockStateMockRecorder) SetControllerReady(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetControllerReady", reflect.TypeOf((*MockState)(nil).SetControllerReady), arg0, arg1, arg2)
}

// SetDBUpgradeCompleted mocks base method.
func (m *MockState) SetDBUpgradeCompleted(arg0 context.Context, arg1 upgrade0.UUID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetDBUpgradeCompleted", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetDBUpgradeCompleted indicates an expected call of SetDBUpgradeCompleted.
func (mr *MockStateMockRecorder) SetDBUpgradeCompleted(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetDBUpgradeCompleted", reflect.TypeOf((*MockState)(nil).SetDBUpgradeCompleted), arg0, arg1)
}

// SetDBUpgradeFailed mocks base method.
func (m *MockState) SetDBUpgradeFailed(arg0 context.Context, arg1 upgrade0.UUID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetDBUpgradeFailed", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetDBUpgradeFailed indicates an expected call of SetDBUpgradeFailed.
func (mr *MockStateMockRecorder) SetDBUpgradeFailed(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetDBUpgradeFailed", reflect.TypeOf((*MockState)(nil).SetDBUpgradeFailed), arg0, arg1)
}

// StartUpgrade mocks base method.
func (m *MockState) StartUpgrade(arg0 context.Context, arg1 upgrade0.UUID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StartUpgrade", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// StartUpgrade indicates an expected call of StartUpgrade.
func (mr *MockStateMockRecorder) StartUpgrade(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StartUpgrade", reflect.TypeOf((*MockState)(nil).StartUpgrade), arg0, arg1)
}

// UpgradeInfo mocks base method.
func (m *MockState) UpgradeInfo(arg0 context.Context, arg1 upgrade0.UUID) (upgrade.Info, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpgradeInfo", arg0, arg1)
	ret0, _ := ret[0].(upgrade.Info)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpgradeInfo indicates an expected call of UpgradeInfo.
func (mr *MockStateMockRecorder) UpgradeInfo(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpgradeInfo", reflect.TypeOf((*MockState)(nil).UpgradeInfo), arg0, arg1)
}

// MockWatcherFactory is a mock of WatcherFactory interface.
type MockWatcherFactory struct {
	ctrl     *gomock.Controller
	recorder *MockWatcherFactoryMockRecorder
}

// MockWatcherFactoryMockRecorder is the mock recorder for MockWatcherFactory.
type MockWatcherFactoryMockRecorder struct {
	mock *MockWatcherFactory
}

// NewMockWatcherFactory creates a new mock instance.
func NewMockWatcherFactory(ctrl *gomock.Controller) *MockWatcherFactory {
	mock := &MockWatcherFactory{ctrl: ctrl}
	mock.recorder = &MockWatcherFactoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockWatcherFactory) EXPECT() *MockWatcherFactoryMockRecorder {
	return m.recorder
}

// NewValuePredicateWatcher mocks base method.
func (m *MockWatcherFactory) NewValuePredicateWatcher(arg0, arg1 string, arg2 changestream.ChangeType, arg3 eventsource.Predicate) (watcher.Watcher[struct{}], error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewValuePredicateWatcher", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(watcher.Watcher[struct{}])
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NewValuePredicateWatcher indicates an expected call of NewValuePredicateWatcher.
func (mr *MockWatcherFactoryMockRecorder) NewValuePredicateWatcher(arg0, arg1, arg2, arg3 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewValuePredicateWatcher", reflect.TypeOf((*MockWatcherFactory)(nil).NewValuePredicateWatcher), arg0, arg1, arg2, arg3)
}
