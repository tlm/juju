// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/juju/juju/core/charm/downloader (interfaces: CharmArchive)
//
// Generated by this command:
//
//	mockgen -package mocks -destination mocks/charm_mocks.go github.com/juju/juju/core/charm/downloader CharmArchive
//

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	charm "github.com/juju/charm/v12"
	gomock "go.uber.org/mock/gomock"
)

// MockCharmArchive is a mock of CharmArchive interface.
type MockCharmArchive struct {
	ctrl     *gomock.Controller
	recorder *MockCharmArchiveMockRecorder
}

// MockCharmArchiveMockRecorder is the mock recorder for MockCharmArchive.
type MockCharmArchiveMockRecorder struct {
	mock *MockCharmArchive
}

// NewMockCharmArchive creates a new mock instance.
func NewMockCharmArchive(ctrl *gomock.Controller) *MockCharmArchive {
	mock := &MockCharmArchive{ctrl: ctrl}
	mock.recorder = &MockCharmArchiveMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCharmArchive) EXPECT() *MockCharmArchiveMockRecorder {
	return m.recorder
}

// Actions mocks base method.
func (m *MockCharmArchive) Actions() *charm.Actions {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Actions")
	ret0, _ := ret[0].(*charm.Actions)
	return ret0
}

// Actions indicates an expected call of Actions.
func (mr *MockCharmArchiveMockRecorder) Actions() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Actions", reflect.TypeOf((*MockCharmArchive)(nil).Actions))
}

// Config mocks base method.
func (m *MockCharmArchive) Config() *charm.Config {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Config")
	ret0, _ := ret[0].(*charm.Config)
	return ret0
}

// Config indicates an expected call of Config.
func (mr *MockCharmArchiveMockRecorder) Config() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Config", reflect.TypeOf((*MockCharmArchive)(nil).Config))
}

// LXDProfile mocks base method.
func (m *MockCharmArchive) LXDProfile() *charm.LXDProfile {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LXDProfile")
	ret0, _ := ret[0].(*charm.LXDProfile)
	return ret0
}

// LXDProfile indicates an expected call of LXDProfile.
func (mr *MockCharmArchiveMockRecorder) LXDProfile() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LXDProfile", reflect.TypeOf((*MockCharmArchive)(nil).LXDProfile))
}

// Manifest mocks base method.
func (m *MockCharmArchive) Manifest() *charm.Manifest {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Manifest")
	ret0, _ := ret[0].(*charm.Manifest)
	return ret0
}

// Manifest indicates an expected call of Manifest.
func (mr *MockCharmArchiveMockRecorder) Manifest() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Manifest", reflect.TypeOf((*MockCharmArchive)(nil).Manifest))
}

// Meta mocks base method.
func (m *MockCharmArchive) Meta() *charm.Meta {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Meta")
	ret0, _ := ret[0].(*charm.Meta)
	return ret0
}

// Meta indicates an expected call of Meta.
func (mr *MockCharmArchiveMockRecorder) Meta() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Meta", reflect.TypeOf((*MockCharmArchive)(nil).Meta))
}

// Metrics mocks base method.
func (m *MockCharmArchive) Metrics() *charm.Metrics {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Metrics")
	ret0, _ := ret[0].(*charm.Metrics)
	return ret0
}

// Metrics indicates an expected call of Metrics.
func (mr *MockCharmArchiveMockRecorder) Metrics() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Metrics", reflect.TypeOf((*MockCharmArchive)(nil).Metrics))
}

// Revision mocks base method.
func (m *MockCharmArchive) Revision() int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Revision")
	ret0, _ := ret[0].(int)
	return ret0
}

// Revision indicates an expected call of Revision.
func (mr *MockCharmArchiveMockRecorder) Revision() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Revision", reflect.TypeOf((*MockCharmArchive)(nil).Revision))
}

// Version mocks base method.
func (m *MockCharmArchive) Version() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Version")
	ret0, _ := ret[0].(string)
	return ret0
}

// Version indicates an expected call of Version.
func (mr *MockCharmArchiveMockRecorder) Version() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Version", reflect.TypeOf((*MockCharmArchive)(nil).Version))
}
