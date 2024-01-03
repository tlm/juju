// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/juju/juju/internal/worker/charmdownloader (interfaces: CharmDownloaderAPI,Logger)
//
// Generated by this command:
//
//	mockgen -package mocks -destination mocks/mocks.go github.com/juju/juju/internal/worker/charmdownloader CharmDownloaderAPI,Logger
//

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	watcher "github.com/juju/juju/core/watcher"
	names "github.com/juju/names/v5"
	gomock "go.uber.org/mock/gomock"
)

// MockCharmDownloaderAPI is a mock of CharmDownloaderAPI interface.
type MockCharmDownloaderAPI struct {
	ctrl     *gomock.Controller
	recorder *MockCharmDownloaderAPIMockRecorder
}

// MockCharmDownloaderAPIMockRecorder is the mock recorder for MockCharmDownloaderAPI.
type MockCharmDownloaderAPIMockRecorder struct {
	mock *MockCharmDownloaderAPI
}

// NewMockCharmDownloaderAPI creates a new mock instance.
func NewMockCharmDownloaderAPI(ctrl *gomock.Controller) *MockCharmDownloaderAPI {
	mock := &MockCharmDownloaderAPI{ctrl: ctrl}
	mock.recorder = &MockCharmDownloaderAPIMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCharmDownloaderAPI) EXPECT() *MockCharmDownloaderAPIMockRecorder {
	return m.recorder
}

// DownloadApplicationCharms mocks base method.
func (m *MockCharmDownloaderAPI) DownloadApplicationCharms(arg0 []names.ApplicationTag) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DownloadApplicationCharms", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// DownloadApplicationCharms indicates an expected call of DownloadApplicationCharms.
func (mr *MockCharmDownloaderAPIMockRecorder) DownloadApplicationCharms(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DownloadApplicationCharms", reflect.TypeOf((*MockCharmDownloaderAPI)(nil).DownloadApplicationCharms), arg0)
}

// WatchApplicationsWithPendingCharms mocks base method.
func (m *MockCharmDownloaderAPI) WatchApplicationsWithPendingCharms() (watcher.Watcher[[]string], error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WatchApplicationsWithPendingCharms")
	ret0, _ := ret[0].(watcher.Watcher[[]string])
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// WatchApplicationsWithPendingCharms indicates an expected call of WatchApplicationsWithPendingCharms.
func (mr *MockCharmDownloaderAPIMockRecorder) WatchApplicationsWithPendingCharms() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WatchApplicationsWithPendingCharms", reflect.TypeOf((*MockCharmDownloaderAPI)(nil).WatchApplicationsWithPendingCharms))
}

// MockLogger is a mock of Logger interface.
type MockLogger struct {
	ctrl     *gomock.Controller
	recorder *MockLoggerMockRecorder
}

// MockLoggerMockRecorder is the mock recorder for MockLogger.
type MockLoggerMockRecorder struct {
	mock *MockLogger
}

// NewMockLogger creates a new mock instance.
func NewMockLogger(ctrl *gomock.Controller) *MockLogger {
	mock := &MockLogger{ctrl: ctrl}
	mock.recorder = &MockLoggerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockLogger) EXPECT() *MockLoggerMockRecorder {
	return m.recorder
}

// Criticalf mocks base method.
func (m *MockLogger) Criticalf(arg0 string, arg1 ...any) {
	m.ctrl.T.Helper()
	varargs := []any{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Criticalf", varargs...)
}

// Criticalf indicates an expected call of Criticalf.
func (mr *MockLoggerMockRecorder) Criticalf(arg0 any, arg1 ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Criticalf", reflect.TypeOf((*MockLogger)(nil).Criticalf), varargs...)
}

// Debugf mocks base method.
func (m *MockLogger) Debugf(arg0 string, arg1 ...any) {
	m.ctrl.T.Helper()
	varargs := []any{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Debugf", varargs...)
}

// Debugf indicates an expected call of Debugf.
func (mr *MockLoggerMockRecorder) Debugf(arg0 any, arg1 ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Debugf", reflect.TypeOf((*MockLogger)(nil).Debugf), varargs...)
}

// Errorf mocks base method.
func (m *MockLogger) Errorf(arg0 string, arg1 ...any) {
	m.ctrl.T.Helper()
	varargs := []any{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Errorf", varargs...)
}

// Errorf indicates an expected call of Errorf.
func (mr *MockLoggerMockRecorder) Errorf(arg0 any, arg1 ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Errorf", reflect.TypeOf((*MockLogger)(nil).Errorf), varargs...)
}

// Infof mocks base method.
func (m *MockLogger) Infof(arg0 string, arg1 ...any) {
	m.ctrl.T.Helper()
	varargs := []any{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Infof", varargs...)
}

// Infof indicates an expected call of Infof.
func (mr *MockLoggerMockRecorder) Infof(arg0 any, arg1 ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Infof", reflect.TypeOf((*MockLogger)(nil).Infof), varargs...)
}

// Tracef mocks base method.
func (m *MockLogger) Tracef(arg0 string, arg1 ...any) {
	m.ctrl.T.Helper()
	varargs := []any{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Tracef", varargs...)
}

// Tracef indicates an expected call of Tracef.
func (mr *MockLoggerMockRecorder) Tracef(arg0 any, arg1 ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Tracef", reflect.TypeOf((*MockLogger)(nil).Tracef), varargs...)
}

// Warningf mocks base method.
func (m *MockLogger) Warningf(arg0 string, arg1 ...any) {
	m.ctrl.T.Helper()
	varargs := []any{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Warningf", varargs...)
}

// Warningf indicates an expected call of Warningf.
func (mr *MockLoggerMockRecorder) Warningf(arg0 any, arg1 ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Warningf", reflect.TypeOf((*MockLogger)(nil).Warningf), varargs...)
}
