// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/juju/juju/domain/secretbackend/service (interfaces: State)
//
// Generated by this command:
//
//	mockgen -typed -package service -destination state_mock_test.go github.com/juju/juju/domain/secretbackend/service State
//

// Package service is a generated GoMock package.
package service

import (
	context "context"
	reflect "reflect"
	time "time"

	cloud "github.com/juju/juju/cloud"
	model "github.com/juju/juju/core/model"
	watcher "github.com/juju/juju/core/watcher"
	secretbackend "github.com/juju/juju/domain/secretbackend"
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

// CreateSecretBackend mocks base method.
func (m *MockState) CreateSecretBackend(arg0 context.Context, arg1 secretbackend.CreateSecretBackendParams) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateSecretBackend", arg0, arg1)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateSecretBackend indicates an expected call of CreateSecretBackend.
func (mr *MockStateMockRecorder) CreateSecretBackend(arg0, arg1 any) *MockStateCreateSecretBackendCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateSecretBackend", reflect.TypeOf((*MockState)(nil).CreateSecretBackend), arg0, arg1)
	return &MockStateCreateSecretBackendCall{Call: call}
}

// MockStateCreateSecretBackendCall wrap *gomock.Call
type MockStateCreateSecretBackendCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockStateCreateSecretBackendCall) Return(arg0 string, arg1 error) *MockStateCreateSecretBackendCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockStateCreateSecretBackendCall) Do(f func(context.Context, secretbackend.CreateSecretBackendParams) (string, error)) *MockStateCreateSecretBackendCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockStateCreateSecretBackendCall) DoAndReturn(f func(context.Context, secretbackend.CreateSecretBackendParams) (string, error)) *MockStateCreateSecretBackendCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// DeleteSecretBackend mocks base method.
func (m *MockState) DeleteSecretBackend(arg0 context.Context, arg1 secretbackend.BackendIdentifier, arg2 bool) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteSecretBackend", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteSecretBackend indicates an expected call of DeleteSecretBackend.
func (mr *MockStateMockRecorder) DeleteSecretBackend(arg0, arg1, arg2 any) *MockStateDeleteSecretBackendCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteSecretBackend", reflect.TypeOf((*MockState)(nil).DeleteSecretBackend), arg0, arg1, arg2)
	return &MockStateDeleteSecretBackendCall{Call: call}
}

// MockStateDeleteSecretBackendCall wrap *gomock.Call
type MockStateDeleteSecretBackendCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockStateDeleteSecretBackendCall) Return(arg0 error) *MockStateDeleteSecretBackendCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockStateDeleteSecretBackendCall) Do(f func(context.Context, secretbackend.BackendIdentifier, bool) error) *MockStateDeleteSecretBackendCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockStateDeleteSecretBackendCall) DoAndReturn(f func(context.Context, secretbackend.BackendIdentifier, bool) error) *MockStateDeleteSecretBackendCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// GetControllerModelCloudAndCredential mocks base method.
func (m *MockState) GetControllerModelCloudAndCredential(arg0 context.Context) (cloud.Cloud, cloud.Credential, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetControllerModelCloudAndCredential", arg0)
	ret0, _ := ret[0].(cloud.Cloud)
	ret1, _ := ret[1].(cloud.Credential)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetControllerModelCloudAndCredential indicates an expected call of GetControllerModelCloudAndCredential.
func (mr *MockStateMockRecorder) GetControllerModelCloudAndCredential(arg0 any) *MockStateGetControllerModelCloudAndCredentialCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetControllerModelCloudAndCredential", reflect.TypeOf((*MockState)(nil).GetControllerModelCloudAndCredential), arg0)
	return &MockStateGetControllerModelCloudAndCredentialCall{Call: call}
}

// MockStateGetControllerModelCloudAndCredentialCall wrap *gomock.Call
type MockStateGetControllerModelCloudAndCredentialCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockStateGetControllerModelCloudAndCredentialCall) Return(arg0 cloud.Cloud, arg1 cloud.Credential, arg2 error) *MockStateGetControllerModelCloudAndCredentialCall {
	c.Call = c.Call.Return(arg0, arg1, arg2)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockStateGetControllerModelCloudAndCredentialCall) Do(f func(context.Context) (cloud.Cloud, cloud.Credential, error)) *MockStateGetControllerModelCloudAndCredentialCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockStateGetControllerModelCloudAndCredentialCall) DoAndReturn(f func(context.Context) (cloud.Cloud, cloud.Credential, error)) *MockStateGetControllerModelCloudAndCredentialCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// GetModelCloudAndCredential mocks base method.
func (m *MockState) GetModelCloudAndCredential(arg0 context.Context, arg1 model.UUID) (cloud.Cloud, cloud.Credential, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetModelCloudAndCredential", arg0, arg1)
	ret0, _ := ret[0].(cloud.Cloud)
	ret1, _ := ret[1].(cloud.Credential)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetModelCloudAndCredential indicates an expected call of GetModelCloudAndCredential.
func (mr *MockStateMockRecorder) GetModelCloudAndCredential(arg0, arg1 any) *MockStateGetModelCloudAndCredentialCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetModelCloudAndCredential", reflect.TypeOf((*MockState)(nil).GetModelCloudAndCredential), arg0, arg1)
	return &MockStateGetModelCloudAndCredentialCall{Call: call}
}

// MockStateGetModelCloudAndCredentialCall wrap *gomock.Call
type MockStateGetModelCloudAndCredentialCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockStateGetModelCloudAndCredentialCall) Return(arg0 cloud.Cloud, arg1 cloud.Credential, arg2 error) *MockStateGetModelCloudAndCredentialCall {
	c.Call = c.Call.Return(arg0, arg1, arg2)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockStateGetModelCloudAndCredentialCall) Do(f func(context.Context, model.UUID) (cloud.Cloud, cloud.Credential, error)) *MockStateGetModelCloudAndCredentialCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockStateGetModelCloudAndCredentialCall) DoAndReturn(f func(context.Context, model.UUID) (cloud.Cloud, cloud.Credential, error)) *MockStateGetModelCloudAndCredentialCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// GetModelSecretBackendDetails mocks base method.
func (m *MockState) GetModelSecretBackendDetails(arg0 context.Context, arg1 model.UUID) (secretbackend.ModelSecretBackend, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetModelSecretBackendDetails", arg0, arg1)
	ret0, _ := ret[0].(secretbackend.ModelSecretBackend)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetModelSecretBackendDetails indicates an expected call of GetModelSecretBackendDetails.
func (mr *MockStateMockRecorder) GetModelSecretBackendDetails(arg0, arg1 any) *MockStateGetModelSecretBackendDetailsCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetModelSecretBackendDetails", reflect.TypeOf((*MockState)(nil).GetModelSecretBackendDetails), arg0, arg1)
	return &MockStateGetModelSecretBackendDetailsCall{Call: call}
}

// MockStateGetModelSecretBackendDetailsCall wrap *gomock.Call
type MockStateGetModelSecretBackendDetailsCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockStateGetModelSecretBackendDetailsCall) Return(arg0 secretbackend.ModelSecretBackend, arg1 error) *MockStateGetModelSecretBackendDetailsCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockStateGetModelSecretBackendDetailsCall) Do(f func(context.Context, model.UUID) (secretbackend.ModelSecretBackend, error)) *MockStateGetModelSecretBackendDetailsCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockStateGetModelSecretBackendDetailsCall) DoAndReturn(f func(context.Context, model.UUID) (secretbackend.ModelSecretBackend, error)) *MockStateGetModelSecretBackendDetailsCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// GetSecretBackend mocks base method.
func (m *MockState) GetSecretBackend(arg0 context.Context, arg1 secretbackend.BackendIdentifier) (*secretbackend.SecretBackend, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSecretBackend", arg0, arg1)
	ret0, _ := ret[0].(*secretbackend.SecretBackend)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSecretBackend indicates an expected call of GetSecretBackend.
func (mr *MockStateMockRecorder) GetSecretBackend(arg0, arg1 any) *MockStateGetSecretBackendCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSecretBackend", reflect.TypeOf((*MockState)(nil).GetSecretBackend), arg0, arg1)
	return &MockStateGetSecretBackendCall{Call: call}
}

// MockStateGetSecretBackendCall wrap *gomock.Call
type MockStateGetSecretBackendCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockStateGetSecretBackendCall) Return(arg0 *secretbackend.SecretBackend, arg1 error) *MockStateGetSecretBackendCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockStateGetSecretBackendCall) Do(f func(context.Context, secretbackend.BackendIdentifier) (*secretbackend.SecretBackend, error)) *MockStateGetSecretBackendCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockStateGetSecretBackendCall) DoAndReturn(f func(context.Context, secretbackend.BackendIdentifier) (*secretbackend.SecretBackend, error)) *MockStateGetSecretBackendCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// GetSecretBackendRotateChanges mocks base method.
func (m *MockState) GetSecretBackendRotateChanges(arg0 context.Context, arg1 ...string) ([]watcher.SecretBackendRotateChange, error) {
	m.ctrl.T.Helper()
	varargs := []any{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetSecretBackendRotateChanges", varargs...)
	ret0, _ := ret[0].([]watcher.SecretBackendRotateChange)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSecretBackendRotateChanges indicates an expected call of GetSecretBackendRotateChanges.
func (mr *MockStateMockRecorder) GetSecretBackendRotateChanges(arg0 any, arg1 ...any) *MockStateGetSecretBackendRotateChangesCall {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{arg0}, arg1...)
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSecretBackendRotateChanges", reflect.TypeOf((*MockState)(nil).GetSecretBackendRotateChanges), varargs...)
	return &MockStateGetSecretBackendRotateChangesCall{Call: call}
}

// MockStateGetSecretBackendRotateChangesCall wrap *gomock.Call
type MockStateGetSecretBackendRotateChangesCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockStateGetSecretBackendRotateChangesCall) Return(arg0 []watcher.SecretBackendRotateChange, arg1 error) *MockStateGetSecretBackendRotateChangesCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockStateGetSecretBackendRotateChangesCall) Do(f func(context.Context, ...string) ([]watcher.SecretBackendRotateChange, error)) *MockStateGetSecretBackendRotateChangesCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockStateGetSecretBackendRotateChangesCall) DoAndReturn(f func(context.Context, ...string) ([]watcher.SecretBackendRotateChange, error)) *MockStateGetSecretBackendRotateChangesCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// InitialWatchStatementForSecretBackendRotationChanges mocks base method.
func (m *MockState) InitialWatchStatementForSecretBackendRotationChanges() (string, string) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InitialWatchStatementForSecretBackendRotationChanges")
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(string)
	return ret0, ret1
}

// InitialWatchStatementForSecretBackendRotationChanges indicates an expected call of InitialWatchStatementForSecretBackendRotationChanges.
func (mr *MockStateMockRecorder) InitialWatchStatementForSecretBackendRotationChanges() *MockStateInitialWatchStatementForSecretBackendRotationChangesCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InitialWatchStatementForSecretBackendRotationChanges", reflect.TypeOf((*MockState)(nil).InitialWatchStatementForSecretBackendRotationChanges))
	return &MockStateInitialWatchStatementForSecretBackendRotationChangesCall{Call: call}
}

// MockStateInitialWatchStatementForSecretBackendRotationChangesCall wrap *gomock.Call
type MockStateInitialWatchStatementForSecretBackendRotationChangesCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockStateInitialWatchStatementForSecretBackendRotationChangesCall) Return(arg0, arg1 string) *MockStateInitialWatchStatementForSecretBackendRotationChangesCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockStateInitialWatchStatementForSecretBackendRotationChangesCall) Do(f func() (string, string)) *MockStateInitialWatchStatementForSecretBackendRotationChangesCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockStateInitialWatchStatementForSecretBackendRotationChangesCall) DoAndReturn(f func() (string, string)) *MockStateInitialWatchStatementForSecretBackendRotationChangesCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// ListInUseKubernetesSecretBackends mocks base method.
func (m *MockState) ListInUseKubernetesSecretBackends(arg0 context.Context) ([]*secretbackend.SecretBackend, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListInUseKubernetesSecretBackends", arg0)
	ret0, _ := ret[0].([]*secretbackend.SecretBackend)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListInUseKubernetesSecretBackends indicates an expected call of ListInUseKubernetesSecretBackends.
func (mr *MockStateMockRecorder) ListInUseKubernetesSecretBackends(arg0 any) *MockStateListInUseKubernetesSecretBackendsCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListInUseKubernetesSecretBackends", reflect.TypeOf((*MockState)(nil).ListInUseKubernetesSecretBackends), arg0)
	return &MockStateListInUseKubernetesSecretBackendsCall{Call: call}
}

// MockStateListInUseKubernetesSecretBackendsCall wrap *gomock.Call
type MockStateListInUseKubernetesSecretBackendsCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockStateListInUseKubernetesSecretBackendsCall) Return(arg0 []*secretbackend.SecretBackend, arg1 error) *MockStateListInUseKubernetesSecretBackendsCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockStateListInUseKubernetesSecretBackendsCall) Do(f func(context.Context) ([]*secretbackend.SecretBackend, error)) *MockStateListInUseKubernetesSecretBackendsCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockStateListInUseKubernetesSecretBackendsCall) DoAndReturn(f func(context.Context) ([]*secretbackend.SecretBackend, error)) *MockStateListInUseKubernetesSecretBackendsCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// ListSecretBackendIDs mocks base method.
func (m *MockState) ListSecretBackendIDs(arg0 context.Context) ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListSecretBackendIDs", arg0)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListSecretBackendIDs indicates an expected call of ListSecretBackendIDs.
func (mr *MockStateMockRecorder) ListSecretBackendIDs(arg0 any) *MockStateListSecretBackendIDsCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListSecretBackendIDs", reflect.TypeOf((*MockState)(nil).ListSecretBackendIDs), arg0)
	return &MockStateListSecretBackendIDsCall{Call: call}
}

// MockStateListSecretBackendIDsCall wrap *gomock.Call
type MockStateListSecretBackendIDsCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockStateListSecretBackendIDsCall) Return(arg0 []string, arg1 error) *MockStateListSecretBackendIDsCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockStateListSecretBackendIDsCall) Do(f func(context.Context) ([]string, error)) *MockStateListSecretBackendIDsCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockStateListSecretBackendIDsCall) DoAndReturn(f func(context.Context) ([]string, error)) *MockStateListSecretBackendIDsCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// ListSecretBackends mocks base method.
func (m *MockState) ListSecretBackends(arg0 context.Context) ([]*secretbackend.SecretBackend, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListSecretBackends", arg0)
	ret0, _ := ret[0].([]*secretbackend.SecretBackend)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListSecretBackends indicates an expected call of ListSecretBackends.
func (mr *MockStateMockRecorder) ListSecretBackends(arg0 any) *MockStateListSecretBackendsCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListSecretBackends", reflect.TypeOf((*MockState)(nil).ListSecretBackends), arg0)
	return &MockStateListSecretBackendsCall{Call: call}
}

// MockStateListSecretBackendsCall wrap *gomock.Call
type MockStateListSecretBackendsCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockStateListSecretBackendsCall) Return(arg0 []*secretbackend.SecretBackend, arg1 error) *MockStateListSecretBackendsCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockStateListSecretBackendsCall) Do(f func(context.Context) ([]*secretbackend.SecretBackend, error)) *MockStateListSecretBackendsCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockStateListSecretBackendsCall) DoAndReturn(f func(context.Context) ([]*secretbackend.SecretBackend, error)) *MockStateListSecretBackendsCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// ListSecretBackendsForModel mocks base method.
func (m *MockState) ListSecretBackendsForModel(arg0 context.Context, arg1 model.UUID, arg2 bool) ([]*secretbackend.SecretBackend, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListSecretBackendsForModel", arg0, arg1, arg2)
	ret0, _ := ret[0].([]*secretbackend.SecretBackend)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListSecretBackendsForModel indicates an expected call of ListSecretBackendsForModel.
func (mr *MockStateMockRecorder) ListSecretBackendsForModel(arg0, arg1, arg2 any) *MockStateListSecretBackendsForModelCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListSecretBackendsForModel", reflect.TypeOf((*MockState)(nil).ListSecretBackendsForModel), arg0, arg1, arg2)
	return &MockStateListSecretBackendsForModelCall{Call: call}
}

// MockStateListSecretBackendsForModelCall wrap *gomock.Call
type MockStateListSecretBackendsForModelCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockStateListSecretBackendsForModelCall) Return(arg0 []*secretbackend.SecretBackend, arg1 error) *MockStateListSecretBackendsForModelCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockStateListSecretBackendsForModelCall) Do(f func(context.Context, model.UUID, bool) ([]*secretbackend.SecretBackend, error)) *MockStateListSecretBackendsForModelCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockStateListSecretBackendsForModelCall) DoAndReturn(f func(context.Context, model.UUID, bool) ([]*secretbackend.SecretBackend, error)) *MockStateListSecretBackendsForModelCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// SecretBackendRotated mocks base method.
func (m *MockState) SecretBackendRotated(arg0 context.Context, arg1 string, arg2 time.Time) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SecretBackendRotated", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// SecretBackendRotated indicates an expected call of SecretBackendRotated.
func (mr *MockStateMockRecorder) SecretBackendRotated(arg0, arg1, arg2 any) *MockStateSecretBackendRotatedCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SecretBackendRotated", reflect.TypeOf((*MockState)(nil).SecretBackendRotated), arg0, arg1, arg2)
	return &MockStateSecretBackendRotatedCall{Call: call}
}

// MockStateSecretBackendRotatedCall wrap *gomock.Call
type MockStateSecretBackendRotatedCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockStateSecretBackendRotatedCall) Return(arg0 error) *MockStateSecretBackendRotatedCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockStateSecretBackendRotatedCall) Do(f func(context.Context, string, time.Time) error) *MockStateSecretBackendRotatedCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockStateSecretBackendRotatedCall) DoAndReturn(f func(context.Context, string, time.Time) error) *MockStateSecretBackendRotatedCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// SetModelSecretBackend mocks base method.
func (m *MockState) SetModelSecretBackend(arg0 context.Context, arg1 model.UUID, arg2 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetModelSecretBackend", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetModelSecretBackend indicates an expected call of SetModelSecretBackend.
func (mr *MockStateMockRecorder) SetModelSecretBackend(arg0, arg1, arg2 any) *MockStateSetModelSecretBackendCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetModelSecretBackend", reflect.TypeOf((*MockState)(nil).SetModelSecretBackend), arg0, arg1, arg2)
	return &MockStateSetModelSecretBackendCall{Call: call}
}

// MockStateSetModelSecretBackendCall wrap *gomock.Call
type MockStateSetModelSecretBackendCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockStateSetModelSecretBackendCall) Return(arg0 error) *MockStateSetModelSecretBackendCall {
	c.Call = c.Call.Return(arg0)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockStateSetModelSecretBackendCall) Do(f func(context.Context, model.UUID, string) error) *MockStateSetModelSecretBackendCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockStateSetModelSecretBackendCall) DoAndReturn(f func(context.Context, model.UUID, string) error) *MockStateSetModelSecretBackendCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}

// UpdateSecretBackend mocks base method.
func (m *MockState) UpdateSecretBackend(arg0 context.Context, arg1 secretbackend.UpdateSecretBackendParams) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateSecretBackend", arg0, arg1)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateSecretBackend indicates an expected call of UpdateSecretBackend.
func (mr *MockStateMockRecorder) UpdateSecretBackend(arg0, arg1 any) *MockStateUpdateSecretBackendCall {
	mr.mock.ctrl.T.Helper()
	call := mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateSecretBackend", reflect.TypeOf((*MockState)(nil).UpdateSecretBackend), arg0, arg1)
	return &MockStateUpdateSecretBackendCall{Call: call}
}

// MockStateUpdateSecretBackendCall wrap *gomock.Call
type MockStateUpdateSecretBackendCall struct {
	*gomock.Call
}

// Return rewrite *gomock.Call.Return
func (c *MockStateUpdateSecretBackendCall) Return(arg0 string, arg1 error) *MockStateUpdateSecretBackendCall {
	c.Call = c.Call.Return(arg0, arg1)
	return c
}

// Do rewrite *gomock.Call.Do
func (c *MockStateUpdateSecretBackendCall) Do(f func(context.Context, secretbackend.UpdateSecretBackendParams) (string, error)) *MockStateUpdateSecretBackendCall {
	c.Call = c.Call.Do(f)
	return c
}

// DoAndReturn rewrite *gomock.Call.DoAndReturn
func (c *MockStateUpdateSecretBackendCall) DoAndReturn(f func(context.Context, secretbackend.UpdateSecretBackendParams) (string, error)) *MockStateUpdateSecretBackendCall {
	c.Call = c.Call.DoAndReturn(f)
	return c
}
