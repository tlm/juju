// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/juju/juju/domain/controllernode/service (interfaces: State)
//
// Generated by this command:
//
//	mockgen -package service -destination package_mock_test.go github.com/juju/juju/domain/controllernode/service State
//

// Package service is a generated GoMock package.
package service

import (
	context "context"
	reflect "reflect"

	model "github.com/juju/juju/core/model"
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

// CurateNodes mocks base method.
func (m *MockState) CurateNodes(arg0 context.Context, arg1, arg2 []string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CurateNodes", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// CurateNodes indicates an expected call of CurateNodes.
func (mr *MockStateMockRecorder) CurateNodes(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CurateNodes", reflect.TypeOf((*MockState)(nil).CurateNodes), arg0, arg1, arg2)
}

// SelectModelUUID mocks base method.
func (m *MockState) SelectModelUUID(arg0 context.Context, arg1 model.UUID) (model.UUID, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SelectModelUUID", arg0, arg1)
	ret0, _ := ret[0].(model.UUID)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SelectModelUUID indicates an expected call of SelectModelUUID.
func (mr *MockStateMockRecorder) SelectModelUUID(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SelectModelUUID", reflect.TypeOf((*MockState)(nil).SelectModelUUID), arg0, arg1)
}

// UpdateDqliteNode mocks base method.
func (m *MockState) UpdateDqliteNode(arg0 context.Context, arg1 string, arg2 uint64, arg3 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateDqliteNode", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateDqliteNode indicates an expected call of UpdateDqliteNode.
func (mr *MockStateMockRecorder) UpdateDqliteNode(arg0, arg1, arg2, arg3 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateDqliteNode", reflect.TypeOf((*MockState)(nil).UpdateDqliteNode), arg0, arg1, arg2, arg3)
}
