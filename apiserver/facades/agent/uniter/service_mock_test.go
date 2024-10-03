// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/juju/juju/apiserver/facades/agent/uniter (interfaces: ModelConfigService,ModelInfoService,NetworkService,MachineService)
//
// Generated by this command:
//
//	mockgen -package uniter_test -destination service_mock_test.go github.com/juju/juju/apiserver/facades/agent/uniter ModelConfigService,ModelInfoService,NetworkService,MachineService
//

// Package uniter_test is a generated GoMock package.
package uniter_test

import (
	context "context"
	reflect "reflect"

	instance "github.com/juju/juju/core/instance"
	machine "github.com/juju/juju/core/machine"
	model "github.com/juju/juju/core/model"
	network "github.com/juju/juju/core/network"
	watcher "github.com/juju/juju/core/watcher"
	config "github.com/juju/juju/environs/config"
	gomock "go.uber.org/mock/gomock"
)

// MockModelConfigService is a mock of ModelConfigService interface.
type MockModelConfigService struct {
	ctrl     *gomock.Controller
	recorder *MockModelConfigServiceMockRecorder
}

// MockModelConfigServiceMockRecorder is the mock recorder for MockModelConfigService.
type MockModelConfigServiceMockRecorder struct {
	mock *MockModelConfigService
}

// NewMockModelConfigService creates a new mock instance.
func NewMockModelConfigService(ctrl *gomock.Controller) *MockModelConfigService {
	mock := &MockModelConfigService{ctrl: ctrl}
	mock.recorder = &MockModelConfigServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockModelConfigService) EXPECT() *MockModelConfigServiceMockRecorder {
	return m.recorder
}

// ModelConfig mocks base method.
func (m *MockModelConfigService) ModelConfig(arg0 context.Context) (*config.Config, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ModelConfig", arg0)
	ret0, _ := ret[0].(*config.Config)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ModelConfig indicates an expected call of ModelConfig.
func (mr *MockModelConfigServiceMockRecorder) ModelConfig(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ModelConfig", reflect.TypeOf((*MockModelConfigService)(nil).ModelConfig), arg0)
}

// Watch mocks base method.
func (m *MockModelConfigService) Watch() (watcher.Watcher[[]string], error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Watch")
	ret0, _ := ret[0].(watcher.Watcher[[]string])
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Watch indicates an expected call of Watch.
func (mr *MockModelConfigServiceMockRecorder) Watch() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Watch", reflect.TypeOf((*MockModelConfigService)(nil).Watch))
}

// MockModelInfoService is a mock of ModelInfoService interface.
type MockModelInfoService struct {
	ctrl     *gomock.Controller
	recorder *MockModelInfoServiceMockRecorder
}

// MockModelInfoServiceMockRecorder is the mock recorder for MockModelInfoService.
type MockModelInfoServiceMockRecorder struct {
	mock *MockModelInfoService
}

// NewMockModelInfoService creates a new mock instance.
func NewMockModelInfoService(ctrl *gomock.Controller) *MockModelInfoService {
	mock := &MockModelInfoService{ctrl: ctrl}
	mock.recorder = &MockModelInfoServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockModelInfoService) EXPECT() *MockModelInfoServiceMockRecorder {
	return m.recorder
}

// GetModelInfo mocks base method.
func (m *MockModelInfoService) GetModelInfo(arg0 context.Context) (model.ReadOnlyModel, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetModelInfo", arg0)
	ret0, _ := ret[0].(model.ReadOnlyModel)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetModelInfo indicates an expected call of GetModelInfo.
func (mr *MockModelInfoServiceMockRecorder) GetModelInfo(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetModelInfo", reflect.TypeOf((*MockModelInfoService)(nil).GetModelInfo), arg0)
}

// MockNetworkService is a mock of NetworkService interface.
type MockNetworkService struct {
	ctrl     *gomock.Controller
	recorder *MockNetworkServiceMockRecorder
}

// MockNetworkServiceMockRecorder is the mock recorder for MockNetworkService.
type MockNetworkServiceMockRecorder struct {
	mock *MockNetworkService
}

// NewMockNetworkService creates a new mock instance.
func NewMockNetworkService(ctrl *gomock.Controller) *MockNetworkService {
	mock := &MockNetworkService{ctrl: ctrl}
	mock.recorder = &MockNetworkServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockNetworkService) EXPECT() *MockNetworkServiceMockRecorder {
	return m.recorder
}

// GetAllSubnets mocks base method.
func (m *MockNetworkService) GetAllSubnets(arg0 context.Context) (network.SubnetInfos, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllSubnets", arg0)
	ret0, _ := ret[0].(network.SubnetInfos)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllSubnets indicates an expected call of GetAllSubnets.
func (mr *MockNetworkServiceMockRecorder) GetAllSubnets(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllSubnets", reflect.TypeOf((*MockNetworkService)(nil).GetAllSubnets), arg0)
}

// SpaceByName mocks base method.
func (m *MockNetworkService) SpaceByName(arg0 context.Context, arg1 string) (*network.SpaceInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SpaceByName", arg0, arg1)
	ret0, _ := ret[0].(*network.SpaceInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SpaceByName indicates an expected call of SpaceByName.
func (mr *MockNetworkServiceMockRecorder) SpaceByName(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SpaceByName", reflect.TypeOf((*MockNetworkService)(nil).SpaceByName), arg0, arg1)
}

// MockMachineService is a mock of MachineService interface.
type MockMachineService struct {
	ctrl     *gomock.Controller
	recorder *MockMachineServiceMockRecorder
}

// MockMachineServiceMockRecorder is the mock recorder for MockMachineService.
type MockMachineServiceMockRecorder struct {
	mock *MockMachineService
}

// NewMockMachineService creates a new mock instance.
func NewMockMachineService(ctrl *gomock.Controller) *MockMachineService {
	mock := &MockMachineService{ctrl: ctrl}
	mock.recorder = &MockMachineServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMachineService) EXPECT() *MockMachineServiceMockRecorder {
	return m.recorder
}

// AppliedLXDProfileNames mocks base method.
func (m *MockMachineService) AppliedLXDProfileNames(arg0 context.Context, arg1 string) ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AppliedLXDProfileNames", arg0, arg1)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AppliedLXDProfileNames indicates an expected call of AppliedLXDProfileNames.
func (mr *MockMachineServiceMockRecorder) AppliedLXDProfileNames(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AppliedLXDProfileNames", reflect.TypeOf((*MockMachineService)(nil).AppliedLXDProfileNames), arg0, arg1)
}

// ClearMachineReboot mocks base method.
func (m *MockMachineService) ClearMachineReboot(arg0 context.Context, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ClearMachineReboot", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// ClearMachineReboot indicates an expected call of ClearMachineReboot.
func (mr *MockMachineServiceMockRecorder) ClearMachineReboot(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ClearMachineReboot", reflect.TypeOf((*MockMachineService)(nil).ClearMachineReboot), arg0, arg1)
}

// GetMachineUUID mocks base method.
func (m *MockMachineService) GetMachineUUID(arg0 context.Context, arg1 machine.Name) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMachineUUID", arg0, arg1)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMachineUUID indicates an expected call of GetMachineUUID.
func (mr *MockMachineServiceMockRecorder) GetMachineUUID(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMachineUUID", reflect.TypeOf((*MockMachineService)(nil).GetMachineUUID), arg0, arg1)
}

// HardwareCharacteristics mocks base method.
func (m *MockMachineService) HardwareCharacteristics(arg0 context.Context, arg1 string) (*instance.HardwareCharacteristics, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HardwareCharacteristics", arg0, arg1)
	ret0, _ := ret[0].(*instance.HardwareCharacteristics)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// HardwareCharacteristics indicates an expected call of HardwareCharacteristics.
func (mr *MockMachineServiceMockRecorder) HardwareCharacteristics(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HardwareCharacteristics", reflect.TypeOf((*MockMachineService)(nil).HardwareCharacteristics), arg0, arg1)
}

// IsMachineRebootRequired mocks base method.
func (m *MockMachineService) IsMachineRebootRequired(arg0 context.Context, arg1 string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsMachineRebootRequired", arg0, arg1)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IsMachineRebootRequired indicates an expected call of IsMachineRebootRequired.
func (mr *MockMachineServiceMockRecorder) IsMachineRebootRequired(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsMachineRebootRequired", reflect.TypeOf((*MockMachineService)(nil).IsMachineRebootRequired), arg0, arg1)
}

// RequireMachineReboot mocks base method.
func (m *MockMachineService) RequireMachineReboot(arg0 context.Context, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RequireMachineReboot", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// RequireMachineReboot indicates an expected call of RequireMachineReboot.
func (mr *MockMachineServiceMockRecorder) RequireMachineReboot(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RequireMachineReboot", reflect.TypeOf((*MockMachineService)(nil).RequireMachineReboot), arg0, arg1)
}

// ShouldRebootOrShutdown mocks base method.
func (m *MockMachineService) ShouldRebootOrShutdown(arg0 context.Context, arg1 string) (machine.RebootAction, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ShouldRebootOrShutdown", arg0, arg1)
	ret0, _ := ret[0].(machine.RebootAction)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ShouldRebootOrShutdown indicates an expected call of ShouldRebootOrShutdown.
func (mr *MockMachineServiceMockRecorder) ShouldRebootOrShutdown(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ShouldRebootOrShutdown", reflect.TypeOf((*MockMachineService)(nil).ShouldRebootOrShutdown), arg0, arg1)
}

// WatchLXDProfiles mocks base method.
func (m *MockMachineService) WatchLXDProfiles(arg0 context.Context, arg1 string) (watcher.Watcher[struct{}], error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WatchLXDProfiles", arg0, arg1)
	ret0, _ := ret[0].(watcher.Watcher[struct{}])
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// WatchLXDProfiles indicates an expected call of WatchLXDProfiles.
func (mr *MockMachineServiceMockRecorder) WatchLXDProfiles(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WatchLXDProfiles", reflect.TypeOf((*MockMachineService)(nil).WatchLXDProfiles), arg0, arg1)
}
