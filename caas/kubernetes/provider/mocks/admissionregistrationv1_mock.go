// Code generated by MockGen. DO NOT EDIT.
// Source: k8s.io/client-go/kubernetes/typed/admissionregistration/v1 (interfaces: AdmissionregistrationV1Interface,MutatingWebhookConfigurationInterface,ValidatingWebhookConfigurationInterface)
//
// Generated by this command:
//
//	mockgen -package mocks -destination mocks/admissionregistrationv1_mock.go -mock_names=MutatingWebhookConfigurationInterface=MockMutatingWebhookConfigurationV1Interface,ValidatingWebhookConfigurationInterface=MockValidatingWebhookConfigurationV1Interface k8s.io/client-go/kubernetes/typed/admissionregistration/v1 AdmissionregistrationV1Interface,MutatingWebhookConfigurationInterface,ValidatingWebhookConfigurationInterface
//

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
	v1 "k8s.io/api/admissionregistration/v1"
	v10 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	v11 "k8s.io/client-go/applyconfigurations/admissionregistration/v1"
	v12 "k8s.io/client-go/kubernetes/typed/admissionregistration/v1"
	rest "k8s.io/client-go/rest"
)

// MockAdmissionregistrationV1Interface is a mock of AdmissionregistrationV1Interface interface.
type MockAdmissionregistrationV1Interface struct {
	ctrl     *gomock.Controller
	recorder *MockAdmissionregistrationV1InterfaceMockRecorder
}

// MockAdmissionregistrationV1InterfaceMockRecorder is the mock recorder for MockAdmissionregistrationV1Interface.
type MockAdmissionregistrationV1InterfaceMockRecorder struct {
	mock *MockAdmissionregistrationV1Interface
}

// NewMockAdmissionregistrationV1Interface creates a new mock instance.
func NewMockAdmissionregistrationV1Interface(ctrl *gomock.Controller) *MockAdmissionregistrationV1Interface {
	mock := &MockAdmissionregistrationV1Interface{ctrl: ctrl}
	mock.recorder = &MockAdmissionregistrationV1InterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAdmissionregistrationV1Interface) EXPECT() *MockAdmissionregistrationV1InterfaceMockRecorder {
	return m.recorder
}

// MutatingWebhookConfigurations mocks base method.
func (m *MockAdmissionregistrationV1Interface) MutatingWebhookConfigurations() v12.MutatingWebhookConfigurationInterface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MutatingWebhookConfigurations")
	ret0, _ := ret[0].(v12.MutatingWebhookConfigurationInterface)
	return ret0
}

// MutatingWebhookConfigurations indicates an expected call of MutatingWebhookConfigurations.
func (mr *MockAdmissionregistrationV1InterfaceMockRecorder) MutatingWebhookConfigurations() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MutatingWebhookConfigurations", reflect.TypeOf((*MockAdmissionregistrationV1Interface)(nil).MutatingWebhookConfigurations))
}

// RESTClient mocks base method.
func (m *MockAdmissionregistrationV1Interface) RESTClient() rest.Interface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RESTClient")
	ret0, _ := ret[0].(rest.Interface)
	return ret0
}

// RESTClient indicates an expected call of RESTClient.
func (mr *MockAdmissionregistrationV1InterfaceMockRecorder) RESTClient() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RESTClient", reflect.TypeOf((*MockAdmissionregistrationV1Interface)(nil).RESTClient))
}

// ValidatingWebhookConfigurations mocks base method.
func (m *MockAdmissionregistrationV1Interface) ValidatingWebhookConfigurations() v12.ValidatingWebhookConfigurationInterface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ValidatingWebhookConfigurations")
	ret0, _ := ret[0].(v12.ValidatingWebhookConfigurationInterface)
	return ret0
}

// ValidatingWebhookConfigurations indicates an expected call of ValidatingWebhookConfigurations.
func (mr *MockAdmissionregistrationV1InterfaceMockRecorder) ValidatingWebhookConfigurations() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ValidatingWebhookConfigurations", reflect.TypeOf((*MockAdmissionregistrationV1Interface)(nil).ValidatingWebhookConfigurations))
}

// MockMutatingWebhookConfigurationV1Interface is a mock of MutatingWebhookConfigurationInterface interface.
type MockMutatingWebhookConfigurationV1Interface struct {
	ctrl     *gomock.Controller
	recorder *MockMutatingWebhookConfigurationV1InterfaceMockRecorder
}

// MockMutatingWebhookConfigurationV1InterfaceMockRecorder is the mock recorder for MockMutatingWebhookConfigurationV1Interface.
type MockMutatingWebhookConfigurationV1InterfaceMockRecorder struct {
	mock *MockMutatingWebhookConfigurationV1Interface
}

// NewMockMutatingWebhookConfigurationV1Interface creates a new mock instance.
func NewMockMutatingWebhookConfigurationV1Interface(ctrl *gomock.Controller) *MockMutatingWebhookConfigurationV1Interface {
	mock := &MockMutatingWebhookConfigurationV1Interface{ctrl: ctrl}
	mock.recorder = &MockMutatingWebhookConfigurationV1InterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMutatingWebhookConfigurationV1Interface) EXPECT() *MockMutatingWebhookConfigurationV1InterfaceMockRecorder {
	return m.recorder
}

// Apply mocks base method.
func (m *MockMutatingWebhookConfigurationV1Interface) Apply(arg0 context.Context, arg1 *v11.MutatingWebhookConfigurationApplyConfiguration, arg2 v10.ApplyOptions) (*v1.MutatingWebhookConfiguration, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Apply", arg0, arg1, arg2)
	ret0, _ := ret[0].(*v1.MutatingWebhookConfiguration)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Apply indicates an expected call of Apply.
func (mr *MockMutatingWebhookConfigurationV1InterfaceMockRecorder) Apply(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Apply", reflect.TypeOf((*MockMutatingWebhookConfigurationV1Interface)(nil).Apply), arg0, arg1, arg2)
}

// Create mocks base method.
func (m *MockMutatingWebhookConfigurationV1Interface) Create(arg0 context.Context, arg1 *v1.MutatingWebhookConfiguration, arg2 v10.CreateOptions) (*v1.MutatingWebhookConfiguration, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0, arg1, arg2)
	ret0, _ := ret[0].(*v1.MutatingWebhookConfiguration)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockMutatingWebhookConfigurationV1InterfaceMockRecorder) Create(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockMutatingWebhookConfigurationV1Interface)(nil).Create), arg0, arg1, arg2)
}

// Delete mocks base method.
func (m *MockMutatingWebhookConfigurationV1Interface) Delete(arg0 context.Context, arg1 string, arg2 v10.DeleteOptions) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockMutatingWebhookConfigurationV1InterfaceMockRecorder) Delete(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockMutatingWebhookConfigurationV1Interface)(nil).Delete), arg0, arg1, arg2)
}

// DeleteCollection mocks base method.
func (m *MockMutatingWebhookConfigurationV1Interface) DeleteCollection(arg0 context.Context, arg1 v10.DeleteOptions, arg2 v10.ListOptions) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteCollection", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteCollection indicates an expected call of DeleteCollection.
func (mr *MockMutatingWebhookConfigurationV1InterfaceMockRecorder) DeleteCollection(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteCollection", reflect.TypeOf((*MockMutatingWebhookConfigurationV1Interface)(nil).DeleteCollection), arg0, arg1, arg2)
}

// Get mocks base method.
func (m *MockMutatingWebhookConfigurationV1Interface) Get(arg0 context.Context, arg1 string, arg2 v10.GetOptions) (*v1.MutatingWebhookConfiguration, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0, arg1, arg2)
	ret0, _ := ret[0].(*v1.MutatingWebhookConfiguration)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockMutatingWebhookConfigurationV1InterfaceMockRecorder) Get(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockMutatingWebhookConfigurationV1Interface)(nil).Get), arg0, arg1, arg2)
}

// List mocks base method.
func (m *MockMutatingWebhookConfigurationV1Interface) List(arg0 context.Context, arg1 v10.ListOptions) (*v1.MutatingWebhookConfigurationList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", arg0, arg1)
	ret0, _ := ret[0].(*v1.MutatingWebhookConfigurationList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockMutatingWebhookConfigurationV1InterfaceMockRecorder) List(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockMutatingWebhookConfigurationV1Interface)(nil).List), arg0, arg1)
}

// Patch mocks base method.
func (m *MockMutatingWebhookConfigurationV1Interface) Patch(arg0 context.Context, arg1 string, arg2 types.PatchType, arg3 []byte, arg4 v10.PatchOptions, arg5 ...string) (*v1.MutatingWebhookConfiguration, error) {
	m.ctrl.T.Helper()
	varargs := []any{arg0, arg1, arg2, arg3, arg4}
	for _, a := range arg5 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Patch", varargs...)
	ret0, _ := ret[0].(*v1.MutatingWebhookConfiguration)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Patch indicates an expected call of Patch.
func (mr *MockMutatingWebhookConfigurationV1InterfaceMockRecorder) Patch(arg0, arg1, arg2, arg3, arg4 any, arg5 ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{arg0, arg1, arg2, arg3, arg4}, arg5...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Patch", reflect.TypeOf((*MockMutatingWebhookConfigurationV1Interface)(nil).Patch), varargs...)
}

// Update mocks base method.
func (m *MockMutatingWebhookConfigurationV1Interface) Update(arg0 context.Context, arg1 *v1.MutatingWebhookConfiguration, arg2 v10.UpdateOptions) (*v1.MutatingWebhookConfiguration, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", arg0, arg1, arg2)
	ret0, _ := ret[0].(*v1.MutatingWebhookConfiguration)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockMutatingWebhookConfigurationV1InterfaceMockRecorder) Update(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockMutatingWebhookConfigurationV1Interface)(nil).Update), arg0, arg1, arg2)
}

// Watch mocks base method.
func (m *MockMutatingWebhookConfigurationV1Interface) Watch(arg0 context.Context, arg1 v10.ListOptions) (watch.Interface, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Watch", arg0, arg1)
	ret0, _ := ret[0].(watch.Interface)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Watch indicates an expected call of Watch.
func (mr *MockMutatingWebhookConfigurationV1InterfaceMockRecorder) Watch(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Watch", reflect.TypeOf((*MockMutatingWebhookConfigurationV1Interface)(nil).Watch), arg0, arg1)
}

// MockValidatingWebhookConfigurationV1Interface is a mock of ValidatingWebhookConfigurationInterface interface.
type MockValidatingWebhookConfigurationV1Interface struct {
	ctrl     *gomock.Controller
	recorder *MockValidatingWebhookConfigurationV1InterfaceMockRecorder
}

// MockValidatingWebhookConfigurationV1InterfaceMockRecorder is the mock recorder for MockValidatingWebhookConfigurationV1Interface.
type MockValidatingWebhookConfigurationV1InterfaceMockRecorder struct {
	mock *MockValidatingWebhookConfigurationV1Interface
}

// NewMockValidatingWebhookConfigurationV1Interface creates a new mock instance.
func NewMockValidatingWebhookConfigurationV1Interface(ctrl *gomock.Controller) *MockValidatingWebhookConfigurationV1Interface {
	mock := &MockValidatingWebhookConfigurationV1Interface{ctrl: ctrl}
	mock.recorder = &MockValidatingWebhookConfigurationV1InterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockValidatingWebhookConfigurationV1Interface) EXPECT() *MockValidatingWebhookConfigurationV1InterfaceMockRecorder {
	return m.recorder
}

// Apply mocks base method.
func (m *MockValidatingWebhookConfigurationV1Interface) Apply(arg0 context.Context, arg1 *v11.ValidatingWebhookConfigurationApplyConfiguration, arg2 v10.ApplyOptions) (*v1.ValidatingWebhookConfiguration, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Apply", arg0, arg1, arg2)
	ret0, _ := ret[0].(*v1.ValidatingWebhookConfiguration)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Apply indicates an expected call of Apply.
func (mr *MockValidatingWebhookConfigurationV1InterfaceMockRecorder) Apply(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Apply", reflect.TypeOf((*MockValidatingWebhookConfigurationV1Interface)(nil).Apply), arg0, arg1, arg2)
}

// Create mocks base method.
func (m *MockValidatingWebhookConfigurationV1Interface) Create(arg0 context.Context, arg1 *v1.ValidatingWebhookConfiguration, arg2 v10.CreateOptions) (*v1.ValidatingWebhookConfiguration, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0, arg1, arg2)
	ret0, _ := ret[0].(*v1.ValidatingWebhookConfiguration)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockValidatingWebhookConfigurationV1InterfaceMockRecorder) Create(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockValidatingWebhookConfigurationV1Interface)(nil).Create), arg0, arg1, arg2)
}

// Delete mocks base method.
func (m *MockValidatingWebhookConfigurationV1Interface) Delete(arg0 context.Context, arg1 string, arg2 v10.DeleteOptions) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockValidatingWebhookConfigurationV1InterfaceMockRecorder) Delete(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockValidatingWebhookConfigurationV1Interface)(nil).Delete), arg0, arg1, arg2)
}

// DeleteCollection mocks base method.
func (m *MockValidatingWebhookConfigurationV1Interface) DeleteCollection(arg0 context.Context, arg1 v10.DeleteOptions, arg2 v10.ListOptions) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteCollection", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteCollection indicates an expected call of DeleteCollection.
func (mr *MockValidatingWebhookConfigurationV1InterfaceMockRecorder) DeleteCollection(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteCollection", reflect.TypeOf((*MockValidatingWebhookConfigurationV1Interface)(nil).DeleteCollection), arg0, arg1, arg2)
}

// Get mocks base method.
func (m *MockValidatingWebhookConfigurationV1Interface) Get(arg0 context.Context, arg1 string, arg2 v10.GetOptions) (*v1.ValidatingWebhookConfiguration, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0, arg1, arg2)
	ret0, _ := ret[0].(*v1.ValidatingWebhookConfiguration)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockValidatingWebhookConfigurationV1InterfaceMockRecorder) Get(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockValidatingWebhookConfigurationV1Interface)(nil).Get), arg0, arg1, arg2)
}

// List mocks base method.
func (m *MockValidatingWebhookConfigurationV1Interface) List(arg0 context.Context, arg1 v10.ListOptions) (*v1.ValidatingWebhookConfigurationList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", arg0, arg1)
	ret0, _ := ret[0].(*v1.ValidatingWebhookConfigurationList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockValidatingWebhookConfigurationV1InterfaceMockRecorder) List(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockValidatingWebhookConfigurationV1Interface)(nil).List), arg0, arg1)
}

// Patch mocks base method.
func (m *MockValidatingWebhookConfigurationV1Interface) Patch(arg0 context.Context, arg1 string, arg2 types.PatchType, arg3 []byte, arg4 v10.PatchOptions, arg5 ...string) (*v1.ValidatingWebhookConfiguration, error) {
	m.ctrl.T.Helper()
	varargs := []any{arg0, arg1, arg2, arg3, arg4}
	for _, a := range arg5 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Patch", varargs...)
	ret0, _ := ret[0].(*v1.ValidatingWebhookConfiguration)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Patch indicates an expected call of Patch.
func (mr *MockValidatingWebhookConfigurationV1InterfaceMockRecorder) Patch(arg0, arg1, arg2, arg3, arg4 any, arg5 ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{arg0, arg1, arg2, arg3, arg4}, arg5...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Patch", reflect.TypeOf((*MockValidatingWebhookConfigurationV1Interface)(nil).Patch), varargs...)
}

// Update mocks base method.
func (m *MockValidatingWebhookConfigurationV1Interface) Update(arg0 context.Context, arg1 *v1.ValidatingWebhookConfiguration, arg2 v10.UpdateOptions) (*v1.ValidatingWebhookConfiguration, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", arg0, arg1, arg2)
	ret0, _ := ret[0].(*v1.ValidatingWebhookConfiguration)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockValidatingWebhookConfigurationV1InterfaceMockRecorder) Update(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockValidatingWebhookConfigurationV1Interface)(nil).Update), arg0, arg1, arg2)
}

// Watch mocks base method.
func (m *MockValidatingWebhookConfigurationV1Interface) Watch(arg0 context.Context, arg1 v10.ListOptions) (watch.Interface, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Watch", arg0, arg1)
	ret0, _ := ret[0].(watch.Interface)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Watch indicates an expected call of Watch.
func (mr *MockValidatingWebhookConfigurationV1InterfaceMockRecorder) Watch(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Watch", reflect.TypeOf((*MockValidatingWebhookConfigurationV1Interface)(nil).Watch), arg0, arg1)
}
