// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/juju/juju/core/database (interfaces: DBGetter,TxnRunner)
//
// Generated by this command:
//
//	mockgen -package modelmigration -destination getter_mock_test.go github.com/juju/juju/core/database DBGetter,TxnRunner
//

// Package modelmigration is a generated GoMock package.
package modelmigration

import (
	context "context"
	sql "database/sql"
	reflect "reflect"

	sqlair "github.com/canonical/sqlair"
	database "github.com/juju/juju/core/database"
	gomock "go.uber.org/mock/gomock"
)

// MockDBGetter is a mock of DBGetter interface.
type MockDBGetter struct {
	ctrl     *gomock.Controller
	recorder *MockDBGetterMockRecorder
}

// MockDBGetterMockRecorder is the mock recorder for MockDBGetter.
type MockDBGetterMockRecorder struct {
	mock *MockDBGetter
}

// NewMockDBGetter creates a new mock instance.
func NewMockDBGetter(ctrl *gomock.Controller) *MockDBGetter {
	mock := &MockDBGetter{ctrl: ctrl}
	mock.recorder = &MockDBGetterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDBGetter) EXPECT() *MockDBGetterMockRecorder {
	return m.recorder
}

// GetDB mocks base method.
func (m *MockDBGetter) GetDB(arg0 string) (database.TxnRunner, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDB", arg0)
	ret0, _ := ret[0].(database.TxnRunner)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDB indicates an expected call of GetDB.
func (mr *MockDBGetterMockRecorder) GetDB(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDB", reflect.TypeOf((*MockDBGetter)(nil).GetDB), arg0)
}

// MockTxnRunner is a mock of TxnRunner interface.
type MockTxnRunner struct {
	ctrl     *gomock.Controller
	recorder *MockTxnRunnerMockRecorder
}

// MockTxnRunnerMockRecorder is the mock recorder for MockTxnRunner.
type MockTxnRunnerMockRecorder struct {
	mock *MockTxnRunner
}

// NewMockTxnRunner creates a new mock instance.
func NewMockTxnRunner(ctrl *gomock.Controller) *MockTxnRunner {
	mock := &MockTxnRunner{ctrl: ctrl}
	mock.recorder = &MockTxnRunnerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTxnRunner) EXPECT() *MockTxnRunnerMockRecorder {
	return m.recorder
}

// StdTxn mocks base method.
func (m *MockTxnRunner) StdTxn(arg0 context.Context, arg1 func(context.Context, *sql.Tx) error) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StdTxn", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// StdTxn indicates an expected call of StdTxn.
func (mr *MockTxnRunnerMockRecorder) StdTxn(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StdTxn", reflect.TypeOf((*MockTxnRunner)(nil).StdTxn), arg0, arg1)
}

// Txn mocks base method.
func (m *MockTxnRunner) Txn(arg0 context.Context, arg1 func(context.Context, *sqlair.TX) error) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Txn", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Txn indicates an expected call of Txn.
func (mr *MockTxnRunnerMockRecorder) Txn(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Txn", reflect.TypeOf((*MockTxnRunner)(nil).Txn), arg0, arg1)
}
