// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/juju/juju/domain/access/service (interfaces: State,UserState,PermissionState)
//
// Generated by this command:
//
//	mockgen -package service -destination state_mock_test.go github.com/juju/juju/domain/access/service State,UserState,PermissionState
//

// Package service is a generated GoMock package.
package service

import (
	context "context"
	reflect "reflect"

	credential "github.com/juju/juju/core/credential"
	permission "github.com/juju/juju/core/permission"
	user "github.com/juju/juju/core/user"
	access "github.com/juju/juju/domain/access"
	auth "github.com/juju/juju/internal/auth"
	uuid "github.com/juju/juju/internal/uuid"
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

// AddUser mocks base method.
func (m *MockState) AddUser(arg0 context.Context, arg1 user.UUID, arg2, arg3 string, arg4 user.UUID, arg5 permission.AccessSpec) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddUser", arg0, arg1, arg2, arg3, arg4, arg5)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddUser indicates an expected call of AddUser.
func (mr *MockStateMockRecorder) AddUser(arg0, arg1, arg2, arg3, arg4, arg5 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddUser", reflect.TypeOf((*MockState)(nil).AddUser), arg0, arg1, arg2, arg3, arg4, arg5)
}

// AddUserWithActivationKey mocks base method.
func (m *MockState) AddUserWithActivationKey(arg0 context.Context, arg1 user.UUID, arg2, arg3 string, arg4 user.UUID, arg5 permission.AccessSpec, arg6 []byte) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddUserWithActivationKey", arg0, arg1, arg2, arg3, arg4, arg5, arg6)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddUserWithActivationKey indicates an expected call of AddUserWithActivationKey.
func (mr *MockStateMockRecorder) AddUserWithActivationKey(arg0, arg1, arg2, arg3, arg4, arg5, arg6 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddUserWithActivationKey", reflect.TypeOf((*MockState)(nil).AddUserWithActivationKey), arg0, arg1, arg2, arg3, arg4, arg5, arg6)
}

// AddUserWithPasswordHash mocks base method.
func (m *MockState) AddUserWithPasswordHash(arg0 context.Context, arg1 user.UUID, arg2, arg3 string, arg4 user.UUID, arg5 permission.AccessSpec, arg6 string, arg7 []byte) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddUserWithPasswordHash", arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddUserWithPasswordHash indicates an expected call of AddUserWithPasswordHash.
func (mr *MockStateMockRecorder) AddUserWithPasswordHash(arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddUserWithPasswordHash", reflect.TypeOf((*MockState)(nil).AddUserWithPasswordHash), arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7)
}

// AllModelAccessForCloudCredential mocks base method.
func (m *MockState) AllModelAccessForCloudCredential(arg0 context.Context, arg1 credential.Key) ([]access.CredentialOwnerModelAccess, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AllModelAccessForCloudCredential", arg0, arg1)
	ret0, _ := ret[0].([]access.CredentialOwnerModelAccess)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AllModelAccessForCloudCredential indicates an expected call of AllModelAccessForCloudCredential.
func (mr *MockStateMockRecorder) AllModelAccessForCloudCredential(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AllModelAccessForCloudCredential", reflect.TypeOf((*MockState)(nil).AllModelAccessForCloudCredential), arg0, arg1)
}

// CreatePermission mocks base method.
func (m *MockState) CreatePermission(arg0 context.Context, arg1 uuid.UUID, arg2 permission.UserAccessSpec) (permission.UserAccess, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreatePermission", arg0, arg1, arg2)
	ret0, _ := ret[0].(permission.UserAccess)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreatePermission indicates an expected call of CreatePermission.
func (mr *MockStateMockRecorder) CreatePermission(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreatePermission", reflect.TypeOf((*MockState)(nil).CreatePermission), arg0, arg1, arg2)
}

// DeletePermission mocks base method.
func (m *MockState) DeletePermission(arg0 context.Context, arg1 string, arg2 permission.ID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeletePermission", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeletePermission indicates an expected call of DeletePermission.
func (mr *MockStateMockRecorder) DeletePermission(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeletePermission", reflect.TypeOf((*MockState)(nil).DeletePermission), arg0, arg1, arg2)
}

// DisableUserAuthentication mocks base method.
func (m *MockState) DisableUserAuthentication(arg0 context.Context, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DisableUserAuthentication", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DisableUserAuthentication indicates an expected call of DisableUserAuthentication.
func (mr *MockStateMockRecorder) DisableUserAuthentication(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DisableUserAuthentication", reflect.TypeOf((*MockState)(nil).DisableUserAuthentication), arg0, arg1)
}

// EnableUserAuthentication mocks base method.
func (m *MockState) EnableUserAuthentication(arg0 context.Context, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EnableUserAuthentication", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// EnableUserAuthentication indicates an expected call of EnableUserAuthentication.
func (mr *MockStateMockRecorder) EnableUserAuthentication(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EnableUserAuthentication", reflect.TypeOf((*MockState)(nil).EnableUserAuthentication), arg0, arg1)
}

// GetActivationKey mocks base method.
func (m *MockState) GetActivationKey(arg0 context.Context, arg1 string) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetActivationKey", arg0, arg1)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetActivationKey indicates an expected call of GetActivationKey.
func (mr *MockStateMockRecorder) GetActivationKey(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetActivationKey", reflect.TypeOf((*MockState)(nil).GetActivationKey), arg0, arg1)
}

// GetAllUsers mocks base method.
func (m *MockState) GetAllUsers(arg0 context.Context) ([]user.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllUsers", arg0)
	ret0, _ := ret[0].([]user.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllUsers indicates an expected call of GetAllUsers.
func (mr *MockStateMockRecorder) GetAllUsers(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllUsers", reflect.TypeOf((*MockState)(nil).GetAllUsers), arg0)
}

// GetUser mocks base method.
func (m *MockState) GetUser(arg0 context.Context, arg1 user.UUID) (user.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUser", arg0, arg1)
	ret0, _ := ret[0].(user.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUser indicates an expected call of GetUser.
func (mr *MockStateMockRecorder) GetUser(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUser", reflect.TypeOf((*MockState)(nil).GetUser), arg0, arg1)
}

// GetUserByAuth mocks base method.
func (m *MockState) GetUserByAuth(arg0 context.Context, arg1 string, arg2 auth.Password) (user.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserByAuth", arg0, arg1, arg2)
	ret0, _ := ret[0].(user.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserByAuth indicates an expected call of GetUserByAuth.
func (mr *MockStateMockRecorder) GetUserByAuth(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserByAuth", reflect.TypeOf((*MockState)(nil).GetUserByAuth), arg0, arg1, arg2)
}

// GetUserByName mocks base method.
func (m *MockState) GetUserByName(arg0 context.Context, arg1 string) (user.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserByName", arg0, arg1)
	ret0, _ := ret[0].(user.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserByName indicates an expected call of GetUserByName.
func (mr *MockStateMockRecorder) GetUserByName(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserByName", reflect.TypeOf((*MockState)(nil).GetUserByName), arg0, arg1)
}

// ReadAllAccessForUserAndObjectType mocks base method.
func (m *MockState) ReadAllAccessForUserAndObjectType(arg0 context.Context, arg1 string, arg2 permission.ObjectType) ([]permission.UserAccess, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReadAllAccessForUserAndObjectType", arg0, arg1, arg2)
	ret0, _ := ret[0].([]permission.UserAccess)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReadAllAccessForUserAndObjectType indicates an expected call of ReadAllAccessForUserAndObjectType.
func (mr *MockStateMockRecorder) ReadAllAccessForUserAndObjectType(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadAllAccessForUserAndObjectType", reflect.TypeOf((*MockState)(nil).ReadAllAccessForUserAndObjectType), arg0, arg1, arg2)
}

// ReadAllUserAccessForTarget mocks base method.
func (m *MockState) ReadAllUserAccessForTarget(arg0 context.Context, arg1 permission.ID) ([]permission.UserAccess, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReadAllUserAccessForTarget", arg0, arg1)
	ret0, _ := ret[0].([]permission.UserAccess)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReadAllUserAccessForTarget indicates an expected call of ReadAllUserAccessForTarget.
func (mr *MockStateMockRecorder) ReadAllUserAccessForTarget(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadAllUserAccessForTarget", reflect.TypeOf((*MockState)(nil).ReadAllUserAccessForTarget), arg0, arg1)
}

// ReadAllUserAccessForUser mocks base method.
func (m *MockState) ReadAllUserAccessForUser(arg0 context.Context, arg1 string) ([]permission.UserAccess, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReadAllUserAccessForUser", arg0, arg1)
	ret0, _ := ret[0].([]permission.UserAccess)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReadAllUserAccessForUser indicates an expected call of ReadAllUserAccessForUser.
func (mr *MockStateMockRecorder) ReadAllUserAccessForUser(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadAllUserAccessForUser", reflect.TypeOf((*MockState)(nil).ReadAllUserAccessForUser), arg0, arg1)
}

// ReadUserAccessForTarget mocks base method.
func (m *MockState) ReadUserAccessForTarget(arg0 context.Context, arg1 string, arg2 permission.ID) (permission.UserAccess, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReadUserAccessForTarget", arg0, arg1, arg2)
	ret0, _ := ret[0].(permission.UserAccess)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReadUserAccessForTarget indicates an expected call of ReadUserAccessForTarget.
func (mr *MockStateMockRecorder) ReadUserAccessForTarget(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadUserAccessForTarget", reflect.TypeOf((*MockState)(nil).ReadUserAccessForTarget), arg0, arg1, arg2)
}

// ReadUserAccessLevelForTarget mocks base method.
func (m *MockState) ReadUserAccessLevelForTarget(arg0 context.Context, arg1 string, arg2 permission.ID) (permission.Access, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReadUserAccessLevelForTarget", arg0, arg1, arg2)
	ret0, _ := ret[0].(permission.Access)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReadUserAccessLevelForTarget indicates an expected call of ReadUserAccessLevelForTarget.
func (mr *MockStateMockRecorder) ReadUserAccessLevelForTarget(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadUserAccessLevelForTarget", reflect.TypeOf((*MockState)(nil).ReadUserAccessLevelForTarget), arg0, arg1, arg2)
}

// RemoveUser mocks base method.
func (m *MockState) RemoveUser(arg0 context.Context, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveUser", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveUser indicates an expected call of RemoveUser.
func (mr *MockStateMockRecorder) RemoveUser(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveUser", reflect.TypeOf((*MockState)(nil).RemoveUser), arg0, arg1)
}

// SetActivationKey mocks base method.
func (m *MockState) SetActivationKey(arg0 context.Context, arg1 string, arg2 []byte) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetActivationKey", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetActivationKey indicates an expected call of SetActivationKey.
func (mr *MockStateMockRecorder) SetActivationKey(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetActivationKey", reflect.TypeOf((*MockState)(nil).SetActivationKey), arg0, arg1, arg2)
}

// SetPasswordHash mocks base method.
func (m *MockState) SetPasswordHash(arg0 context.Context, arg1, arg2 string, arg3 []byte) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetPasswordHash", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetPasswordHash indicates an expected call of SetPasswordHash.
func (mr *MockStateMockRecorder) SetPasswordHash(arg0, arg1, arg2, arg3 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetPasswordHash", reflect.TypeOf((*MockState)(nil).SetPasswordHash), arg0, arg1, arg2, arg3)
}

// UpdateLastLogin mocks base method.
func (m *MockState) UpdateLastLogin(arg0 context.Context, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateLastLogin", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateLastLogin indicates an expected call of UpdateLastLogin.
func (mr *MockStateMockRecorder) UpdateLastLogin(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateLastLogin", reflect.TypeOf((*MockState)(nil).UpdateLastLogin), arg0, arg1)
}

// UpsertPermission mocks base method.
func (m *MockState) UpsertPermission(arg0 context.Context, arg1 access.UpdatePermissionArgs) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpsertPermission", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpsertPermission indicates an expected call of UpsertPermission.
func (mr *MockStateMockRecorder) UpsertPermission(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpsertPermission", reflect.TypeOf((*MockState)(nil).UpsertPermission), arg0, arg1)
}

// MockUserState is a mock of UserState interface.
type MockUserState struct {
	ctrl     *gomock.Controller
	recorder *MockUserStateMockRecorder
}

// MockUserStateMockRecorder is the mock recorder for MockUserState.
type MockUserStateMockRecorder struct {
	mock *MockUserState
}

// NewMockUserState creates a new mock instance.
func NewMockUserState(ctrl *gomock.Controller) *MockUserState {
	mock := &MockUserState{ctrl: ctrl}
	mock.recorder = &MockUserStateMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserState) EXPECT() *MockUserStateMockRecorder {
	return m.recorder
}

// AddUser mocks base method.
func (m *MockUserState) AddUser(arg0 context.Context, arg1 user.UUID, arg2, arg3 string, arg4 user.UUID, arg5 permission.AccessSpec) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddUser", arg0, arg1, arg2, arg3, arg4, arg5)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddUser indicates an expected call of AddUser.
func (mr *MockUserStateMockRecorder) AddUser(arg0, arg1, arg2, arg3, arg4, arg5 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddUser", reflect.TypeOf((*MockUserState)(nil).AddUser), arg0, arg1, arg2, arg3, arg4, arg5)
}

// AddUserWithActivationKey mocks base method.
func (m *MockUserState) AddUserWithActivationKey(arg0 context.Context, arg1 user.UUID, arg2, arg3 string, arg4 user.UUID, arg5 permission.AccessSpec, arg6 []byte) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddUserWithActivationKey", arg0, arg1, arg2, arg3, arg4, arg5, arg6)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddUserWithActivationKey indicates an expected call of AddUserWithActivationKey.
func (mr *MockUserStateMockRecorder) AddUserWithActivationKey(arg0, arg1, arg2, arg3, arg4, arg5, arg6 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddUserWithActivationKey", reflect.TypeOf((*MockUserState)(nil).AddUserWithActivationKey), arg0, arg1, arg2, arg3, arg4, arg5, arg6)
}

// AddUserWithPasswordHash mocks base method.
func (m *MockUserState) AddUserWithPasswordHash(arg0 context.Context, arg1 user.UUID, arg2, arg3 string, arg4 user.UUID, arg5 permission.AccessSpec, arg6 string, arg7 []byte) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddUserWithPasswordHash", arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddUserWithPasswordHash indicates an expected call of AddUserWithPasswordHash.
func (mr *MockUserStateMockRecorder) AddUserWithPasswordHash(arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddUserWithPasswordHash", reflect.TypeOf((*MockUserState)(nil).AddUserWithPasswordHash), arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7)
}

// DisableUserAuthentication mocks base method.
func (m *MockUserState) DisableUserAuthentication(arg0 context.Context, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DisableUserAuthentication", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DisableUserAuthentication indicates an expected call of DisableUserAuthentication.
func (mr *MockUserStateMockRecorder) DisableUserAuthentication(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DisableUserAuthentication", reflect.TypeOf((*MockUserState)(nil).DisableUserAuthentication), arg0, arg1)
}

// EnableUserAuthentication mocks base method.
func (m *MockUserState) EnableUserAuthentication(arg0 context.Context, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EnableUserAuthentication", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// EnableUserAuthentication indicates an expected call of EnableUserAuthentication.
func (mr *MockUserStateMockRecorder) EnableUserAuthentication(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EnableUserAuthentication", reflect.TypeOf((*MockUserState)(nil).EnableUserAuthentication), arg0, arg1)
}

// GetActivationKey mocks base method.
func (m *MockUserState) GetActivationKey(arg0 context.Context, arg1 string) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetActivationKey", arg0, arg1)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetActivationKey indicates an expected call of GetActivationKey.
func (mr *MockUserStateMockRecorder) GetActivationKey(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetActivationKey", reflect.TypeOf((*MockUserState)(nil).GetActivationKey), arg0, arg1)
}

// GetAllUsers mocks base method.
func (m *MockUserState) GetAllUsers(arg0 context.Context) ([]user.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllUsers", arg0)
	ret0, _ := ret[0].([]user.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllUsers indicates an expected call of GetAllUsers.
func (mr *MockUserStateMockRecorder) GetAllUsers(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllUsers", reflect.TypeOf((*MockUserState)(nil).GetAllUsers), arg0)
}

// GetUser mocks base method.
func (m *MockUserState) GetUser(arg0 context.Context, arg1 user.UUID) (user.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUser", arg0, arg1)
	ret0, _ := ret[0].(user.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUser indicates an expected call of GetUser.
func (mr *MockUserStateMockRecorder) GetUser(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUser", reflect.TypeOf((*MockUserState)(nil).GetUser), arg0, arg1)
}

// GetUserByAuth mocks base method.
func (m *MockUserState) GetUserByAuth(arg0 context.Context, arg1 string, arg2 auth.Password) (user.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserByAuth", arg0, arg1, arg2)
	ret0, _ := ret[0].(user.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserByAuth indicates an expected call of GetUserByAuth.
func (mr *MockUserStateMockRecorder) GetUserByAuth(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserByAuth", reflect.TypeOf((*MockUserState)(nil).GetUserByAuth), arg0, arg1, arg2)
}

// GetUserByName mocks base method.
func (m *MockUserState) GetUserByName(arg0 context.Context, arg1 string) (user.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserByName", arg0, arg1)
	ret0, _ := ret[0].(user.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserByName indicates an expected call of GetUserByName.
func (mr *MockUserStateMockRecorder) GetUserByName(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserByName", reflect.TypeOf((*MockUserState)(nil).GetUserByName), arg0, arg1)
}

// RemoveUser mocks base method.
func (m *MockUserState) RemoveUser(arg0 context.Context, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveUser", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveUser indicates an expected call of RemoveUser.
func (mr *MockUserStateMockRecorder) RemoveUser(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveUser", reflect.TypeOf((*MockUserState)(nil).RemoveUser), arg0, arg1)
}

// SetActivationKey mocks base method.
func (m *MockUserState) SetActivationKey(arg0 context.Context, arg1 string, arg2 []byte) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetActivationKey", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetActivationKey indicates an expected call of SetActivationKey.
func (mr *MockUserStateMockRecorder) SetActivationKey(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetActivationKey", reflect.TypeOf((*MockUserState)(nil).SetActivationKey), arg0, arg1, arg2)
}

// SetPasswordHash mocks base method.
func (m *MockUserState) SetPasswordHash(arg0 context.Context, arg1, arg2 string, arg3 []byte) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetPasswordHash", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetPasswordHash indicates an expected call of SetPasswordHash.
func (mr *MockUserStateMockRecorder) SetPasswordHash(arg0, arg1, arg2, arg3 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetPasswordHash", reflect.TypeOf((*MockUserState)(nil).SetPasswordHash), arg0, arg1, arg2, arg3)
}

// UpdateLastLogin mocks base method.
func (m *MockUserState) UpdateLastLogin(arg0 context.Context, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateLastLogin", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateLastLogin indicates an expected call of UpdateLastLogin.
func (mr *MockUserStateMockRecorder) UpdateLastLogin(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateLastLogin", reflect.TypeOf((*MockUserState)(nil).UpdateLastLogin), arg0, arg1)
}

// MockPermissionState is a mock of PermissionState interface.
type MockPermissionState struct {
	ctrl     *gomock.Controller
	recorder *MockPermissionStateMockRecorder
}

// MockPermissionStateMockRecorder is the mock recorder for MockPermissionState.
type MockPermissionStateMockRecorder struct {
	mock *MockPermissionState
}

// NewMockPermissionState creates a new mock instance.
func NewMockPermissionState(ctrl *gomock.Controller) *MockPermissionState {
	mock := &MockPermissionState{ctrl: ctrl}
	mock.recorder = &MockPermissionStateMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPermissionState) EXPECT() *MockPermissionStateMockRecorder {
	return m.recorder
}

// AllModelAccessForCloudCredential mocks base method.
func (m *MockPermissionState) AllModelAccessForCloudCredential(arg0 context.Context, arg1 credential.Key) ([]access.CredentialOwnerModelAccess, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AllModelAccessForCloudCredential", arg0, arg1)
	ret0, _ := ret[0].([]access.CredentialOwnerModelAccess)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AllModelAccessForCloudCredential indicates an expected call of AllModelAccessForCloudCredential.
func (mr *MockPermissionStateMockRecorder) AllModelAccessForCloudCredential(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AllModelAccessForCloudCredential", reflect.TypeOf((*MockPermissionState)(nil).AllModelAccessForCloudCredential), arg0, arg1)
}

// CreatePermission mocks base method.
func (m *MockPermissionState) CreatePermission(arg0 context.Context, arg1 uuid.UUID, arg2 permission.UserAccessSpec) (permission.UserAccess, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreatePermission", arg0, arg1, arg2)
	ret0, _ := ret[0].(permission.UserAccess)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreatePermission indicates an expected call of CreatePermission.
func (mr *MockPermissionStateMockRecorder) CreatePermission(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreatePermission", reflect.TypeOf((*MockPermissionState)(nil).CreatePermission), arg0, arg1, arg2)
}

// DeletePermission mocks base method.
func (m *MockPermissionState) DeletePermission(arg0 context.Context, arg1 string, arg2 permission.ID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeletePermission", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeletePermission indicates an expected call of DeletePermission.
func (mr *MockPermissionStateMockRecorder) DeletePermission(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeletePermission", reflect.TypeOf((*MockPermissionState)(nil).DeletePermission), arg0, arg1, arg2)
}

// ReadAllAccessForUserAndObjectType mocks base method.
func (m *MockPermissionState) ReadAllAccessForUserAndObjectType(arg0 context.Context, arg1 string, arg2 permission.ObjectType) ([]permission.UserAccess, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReadAllAccessForUserAndObjectType", arg0, arg1, arg2)
	ret0, _ := ret[0].([]permission.UserAccess)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReadAllAccessForUserAndObjectType indicates an expected call of ReadAllAccessForUserAndObjectType.
func (mr *MockPermissionStateMockRecorder) ReadAllAccessForUserAndObjectType(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadAllAccessForUserAndObjectType", reflect.TypeOf((*MockPermissionState)(nil).ReadAllAccessForUserAndObjectType), arg0, arg1, arg2)
}

// ReadAllUserAccessForTarget mocks base method.
func (m *MockPermissionState) ReadAllUserAccessForTarget(arg0 context.Context, arg1 permission.ID) ([]permission.UserAccess, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReadAllUserAccessForTarget", arg0, arg1)
	ret0, _ := ret[0].([]permission.UserAccess)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReadAllUserAccessForTarget indicates an expected call of ReadAllUserAccessForTarget.
func (mr *MockPermissionStateMockRecorder) ReadAllUserAccessForTarget(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadAllUserAccessForTarget", reflect.TypeOf((*MockPermissionState)(nil).ReadAllUserAccessForTarget), arg0, arg1)
}

// ReadAllUserAccessForUser mocks base method.
func (m *MockPermissionState) ReadAllUserAccessForUser(arg0 context.Context, arg1 string) ([]permission.UserAccess, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReadAllUserAccessForUser", arg0, arg1)
	ret0, _ := ret[0].([]permission.UserAccess)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReadAllUserAccessForUser indicates an expected call of ReadAllUserAccessForUser.
func (mr *MockPermissionStateMockRecorder) ReadAllUserAccessForUser(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadAllUserAccessForUser", reflect.TypeOf((*MockPermissionState)(nil).ReadAllUserAccessForUser), arg0, arg1)
}

// ReadUserAccessForTarget mocks base method.
func (m *MockPermissionState) ReadUserAccessForTarget(arg0 context.Context, arg1 string, arg2 permission.ID) (permission.UserAccess, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReadUserAccessForTarget", arg0, arg1, arg2)
	ret0, _ := ret[0].(permission.UserAccess)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReadUserAccessForTarget indicates an expected call of ReadUserAccessForTarget.
func (mr *MockPermissionStateMockRecorder) ReadUserAccessForTarget(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadUserAccessForTarget", reflect.TypeOf((*MockPermissionState)(nil).ReadUserAccessForTarget), arg0, arg1, arg2)
}

// ReadUserAccessLevelForTarget mocks base method.
func (m *MockPermissionState) ReadUserAccessLevelForTarget(arg0 context.Context, arg1 string, arg2 permission.ID) (permission.Access, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReadUserAccessLevelForTarget", arg0, arg1, arg2)
	ret0, _ := ret[0].(permission.Access)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReadUserAccessLevelForTarget indicates an expected call of ReadUserAccessLevelForTarget.
func (mr *MockPermissionStateMockRecorder) ReadUserAccessLevelForTarget(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadUserAccessLevelForTarget", reflect.TypeOf((*MockPermissionState)(nil).ReadUserAccessLevelForTarget), arg0, arg1, arg2)
}

// UpsertPermission mocks base method.
func (m *MockPermissionState) UpsertPermission(arg0 context.Context, arg1 access.UpdatePermissionArgs) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpsertPermission", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpsertPermission indicates an expected call of UpsertPermission.
func (mr *MockPermissionStateMockRecorder) UpsertPermission(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpsertPermission", reflect.TypeOf((*MockPermissionState)(nil).UpsertPermission), arg0, arg1)
}
