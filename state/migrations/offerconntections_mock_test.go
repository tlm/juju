// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/juju/juju/state/migrations (interfaces: MigrationOfferConnection,AllOfferConnectionSource,OfferConnectionSource,OfferConnectionModel)
//
// Generated by this command:
//
//	mockgen -package migrations -destination offerconntections_mock_test.go github.com/juju/juju/state/migrations MigrationOfferConnection,AllOfferConnectionSource,OfferConnectionSource,OfferConnectionModel
//

// Package migrations is a generated GoMock package.
package migrations

import (
	reflect "reflect"

	description "github.com/juju/description/v5"
	gomock "go.uber.org/mock/gomock"
)

// MockMigrationOfferConnection is a mock of MigrationOfferConnection interface.
type MockMigrationOfferConnection struct {
	ctrl     *gomock.Controller
	recorder *MockMigrationOfferConnectionMockRecorder
}

// MockMigrationOfferConnectionMockRecorder is the mock recorder for MockMigrationOfferConnection.
type MockMigrationOfferConnectionMockRecorder struct {
	mock *MockMigrationOfferConnection
}

// NewMockMigrationOfferConnection creates a new mock instance.
func NewMockMigrationOfferConnection(ctrl *gomock.Controller) *MockMigrationOfferConnection {
	mock := &MockMigrationOfferConnection{ctrl: ctrl}
	mock.recorder = &MockMigrationOfferConnectionMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMigrationOfferConnection) EXPECT() *MockMigrationOfferConnectionMockRecorder {
	return m.recorder
}

// OfferUUID mocks base method.
func (m *MockMigrationOfferConnection) OfferUUID() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "OfferUUID")
	ret0, _ := ret[0].(string)
	return ret0
}

// OfferUUID indicates an expected call of OfferUUID.
func (mr *MockMigrationOfferConnectionMockRecorder) OfferUUID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OfferUUID", reflect.TypeOf((*MockMigrationOfferConnection)(nil).OfferUUID))
}

// RelationId mocks base method.
func (m *MockMigrationOfferConnection) RelationId() int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RelationId")
	ret0, _ := ret[0].(int)
	return ret0
}

// RelationId indicates an expected call of RelationId.
func (mr *MockMigrationOfferConnectionMockRecorder) RelationId() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RelationId", reflect.TypeOf((*MockMigrationOfferConnection)(nil).RelationId))
}

// RelationKey mocks base method.
func (m *MockMigrationOfferConnection) RelationKey() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RelationKey")
	ret0, _ := ret[0].(string)
	return ret0
}

// RelationKey indicates an expected call of RelationKey.
func (mr *MockMigrationOfferConnectionMockRecorder) RelationKey() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RelationKey", reflect.TypeOf((*MockMigrationOfferConnection)(nil).RelationKey))
}

// SourceModelUUID mocks base method.
func (m *MockMigrationOfferConnection) SourceModelUUID() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SourceModelUUID")
	ret0, _ := ret[0].(string)
	return ret0
}

// SourceModelUUID indicates an expected call of SourceModelUUID.
func (mr *MockMigrationOfferConnectionMockRecorder) SourceModelUUID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SourceModelUUID", reflect.TypeOf((*MockMigrationOfferConnection)(nil).SourceModelUUID))
}

// UserName mocks base method.
func (m *MockMigrationOfferConnection) UserName() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UserName")
	ret0, _ := ret[0].(string)
	return ret0
}

// UserName indicates an expected call of UserName.
func (mr *MockMigrationOfferConnectionMockRecorder) UserName() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UserName", reflect.TypeOf((*MockMigrationOfferConnection)(nil).UserName))
}

// MockAllOfferConnectionSource is a mock of AllOfferConnectionSource interface.
type MockAllOfferConnectionSource struct {
	ctrl     *gomock.Controller
	recorder *MockAllOfferConnectionSourceMockRecorder
}

// MockAllOfferConnectionSourceMockRecorder is the mock recorder for MockAllOfferConnectionSource.
type MockAllOfferConnectionSourceMockRecorder struct {
	mock *MockAllOfferConnectionSource
}

// NewMockAllOfferConnectionSource creates a new mock instance.
func NewMockAllOfferConnectionSource(ctrl *gomock.Controller) *MockAllOfferConnectionSource {
	mock := &MockAllOfferConnectionSource{ctrl: ctrl}
	mock.recorder = &MockAllOfferConnectionSourceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAllOfferConnectionSource) EXPECT() *MockAllOfferConnectionSourceMockRecorder {
	return m.recorder
}

// AllOfferConnections mocks base method.
func (m *MockAllOfferConnectionSource) AllOfferConnections() ([]MigrationOfferConnection, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AllOfferConnections")
	ret0, _ := ret[0].([]MigrationOfferConnection)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AllOfferConnections indicates an expected call of AllOfferConnections.
func (mr *MockAllOfferConnectionSourceMockRecorder) AllOfferConnections() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AllOfferConnections", reflect.TypeOf((*MockAllOfferConnectionSource)(nil).AllOfferConnections))
}

// MockOfferConnectionSource is a mock of OfferConnectionSource interface.
type MockOfferConnectionSource struct {
	ctrl     *gomock.Controller
	recorder *MockOfferConnectionSourceMockRecorder
}

// MockOfferConnectionSourceMockRecorder is the mock recorder for MockOfferConnectionSource.
type MockOfferConnectionSourceMockRecorder struct {
	mock *MockOfferConnectionSource
}

// NewMockOfferConnectionSource creates a new mock instance.
func NewMockOfferConnectionSource(ctrl *gomock.Controller) *MockOfferConnectionSource {
	mock := &MockOfferConnectionSource{ctrl: ctrl}
	mock.recorder = &MockOfferConnectionSourceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockOfferConnectionSource) EXPECT() *MockOfferConnectionSourceMockRecorder {
	return m.recorder
}

// AllOfferConnections mocks base method.
func (m *MockOfferConnectionSource) AllOfferConnections() ([]MigrationOfferConnection, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AllOfferConnections")
	ret0, _ := ret[0].([]MigrationOfferConnection)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AllOfferConnections indicates an expected call of AllOfferConnections.
func (mr *MockOfferConnectionSourceMockRecorder) AllOfferConnections() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AllOfferConnections", reflect.TypeOf((*MockOfferConnectionSource)(nil).AllOfferConnections))
}

// MockOfferConnectionModel is a mock of OfferConnectionModel interface.
type MockOfferConnectionModel struct {
	ctrl     *gomock.Controller
	recorder *MockOfferConnectionModelMockRecorder
}

// MockOfferConnectionModelMockRecorder is the mock recorder for MockOfferConnectionModel.
type MockOfferConnectionModelMockRecorder struct {
	mock *MockOfferConnectionModel
}

// NewMockOfferConnectionModel creates a new mock instance.
func NewMockOfferConnectionModel(ctrl *gomock.Controller) *MockOfferConnectionModel {
	mock := &MockOfferConnectionModel{ctrl: ctrl}
	mock.recorder = &MockOfferConnectionModelMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockOfferConnectionModel) EXPECT() *MockOfferConnectionModelMockRecorder {
	return m.recorder
}

// AddOfferConnection mocks base method.
func (m *MockOfferConnectionModel) AddOfferConnection(arg0 description.OfferConnectionArgs) description.OfferConnection {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddOfferConnection", arg0)
	ret0, _ := ret[0].(description.OfferConnection)
	return ret0
}

// AddOfferConnection indicates an expected call of AddOfferConnection.
func (mr *MockOfferConnectionModelMockRecorder) AddOfferConnection(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddOfferConnection", reflect.TypeOf((*MockOfferConnectionModel)(nil).AddOfferConnection), arg0)
}
