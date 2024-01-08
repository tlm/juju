// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/juju/juju/api/base (interfaces: APICallCloser,ClientFacade)
//
// Generated by this command:
//
//	mockgen -package mocks -destination mocks/clientfacade_mock.go github.com/juju/juju/api/base APICallCloser,ClientFacade
//

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	http "net/http"
	url "net/url"
	reflect "reflect"

	base "github.com/juju/juju/api/base"
	names "github.com/juju/names/v5"
	gomock "go.uber.org/mock/gomock"
	httprequest "gopkg.in/httprequest.v1"
)

// MockAPICallCloser is a mock of APICallCloser interface.
type MockAPICallCloser struct {
	ctrl     *gomock.Controller
	recorder *MockAPICallCloserMockRecorder
}

// MockAPICallCloserMockRecorder is the mock recorder for MockAPICallCloser.
type MockAPICallCloserMockRecorder struct {
	mock *MockAPICallCloser
}

// NewMockAPICallCloser creates a new mock instance.
func NewMockAPICallCloser(ctrl *gomock.Controller) *MockAPICallCloser {
	mock := &MockAPICallCloser{ctrl: ctrl}
	mock.recorder = &MockAPICallCloserMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAPICallCloser) EXPECT() *MockAPICallCloserMockRecorder {
	return m.recorder
}

// APICall mocks base method.
func (m *MockAPICallCloser) APICall(arg0 context.Context, arg1 string, arg2 int, arg3, arg4 string, arg5, arg6 any) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "APICall", arg0, arg1, arg2, arg3, arg4, arg5, arg6)
	ret0, _ := ret[0].(error)
	return ret0
}

// APICall indicates an expected call of APICall.
func (mr *MockAPICallCloserMockRecorder) APICall(arg0, arg1, arg2, arg3, arg4, arg5, arg6 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "APICall", reflect.TypeOf((*MockAPICallCloser)(nil).APICall), arg0, arg1, arg2, arg3, arg4, arg5, arg6)
}

// BakeryClient mocks base method.
func (m *MockAPICallCloser) BakeryClient() base.MacaroonDischarger {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BakeryClient")
	ret0, _ := ret[0].(base.MacaroonDischarger)
	return ret0
}

// BakeryClient indicates an expected call of BakeryClient.
func (mr *MockAPICallCloserMockRecorder) BakeryClient() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BakeryClient", reflect.TypeOf((*MockAPICallCloser)(nil).BakeryClient))
}

// BestFacadeVersion mocks base method.
func (m *MockAPICallCloser) BestFacadeVersion(arg0 string) int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BestFacadeVersion", arg0)
	ret0, _ := ret[0].(int)
	return ret0
}

// BestFacadeVersion indicates an expected call of BestFacadeVersion.
func (mr *MockAPICallCloserMockRecorder) BestFacadeVersion(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BestFacadeVersion", reflect.TypeOf((*MockAPICallCloser)(nil).BestFacadeVersion), arg0)
}

// Close mocks base method.
func (m *MockAPICallCloser) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close.
func (mr *MockAPICallCloserMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockAPICallCloser)(nil).Close))
}

// ConnectControllerStream mocks base method.
func (m *MockAPICallCloser) ConnectControllerStream(arg0 string, arg1 url.Values, arg2 http.Header) (base.Stream, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ConnectControllerStream", arg0, arg1, arg2)
	ret0, _ := ret[0].(base.Stream)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ConnectControllerStream indicates an expected call of ConnectControllerStream.
func (mr *MockAPICallCloserMockRecorder) ConnectControllerStream(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ConnectControllerStream", reflect.TypeOf((*MockAPICallCloser)(nil).ConnectControllerStream), arg0, arg1, arg2)
}

// ConnectStream mocks base method.
func (m *MockAPICallCloser) ConnectStream(arg0 string, arg1 url.Values) (base.Stream, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ConnectStream", arg0, arg1)
	ret0, _ := ret[0].(base.Stream)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ConnectStream indicates an expected call of ConnectStream.
func (mr *MockAPICallCloserMockRecorder) ConnectStream(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ConnectStream", reflect.TypeOf((*MockAPICallCloser)(nil).ConnectStream), arg0, arg1)
}

// Context mocks base method.
func (m *MockAPICallCloser) Context() context.Context {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Context")
	ret0, _ := ret[0].(context.Context)
	return ret0
}

// Context indicates an expected call of Context.
func (mr *MockAPICallCloserMockRecorder) Context() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Context", reflect.TypeOf((*MockAPICallCloser)(nil).Context))
}

// HTTPClient mocks base method.
func (m *MockAPICallCloser) HTTPClient() (*httprequest.Client, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HTTPClient")
	ret0, _ := ret[0].(*httprequest.Client)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// HTTPClient indicates an expected call of HTTPClient.
func (mr *MockAPICallCloserMockRecorder) HTTPClient() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HTTPClient", reflect.TypeOf((*MockAPICallCloser)(nil).HTTPClient))
}

// ModelTag mocks base method.
func (m *MockAPICallCloser) ModelTag() (names.ModelTag, bool) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ModelTag")
	ret0, _ := ret[0].(names.ModelTag)
	ret1, _ := ret[1].(bool)
	return ret0, ret1
}

// ModelTag indicates an expected call of ModelTag.
func (mr *MockAPICallCloserMockRecorder) ModelTag() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ModelTag", reflect.TypeOf((*MockAPICallCloser)(nil).ModelTag))
}

// RootHTTPClient mocks base method.
func (m *MockAPICallCloser) RootHTTPClient() (*httprequest.Client, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RootHTTPClient")
	ret0, _ := ret[0].(*httprequest.Client)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RootHTTPClient indicates an expected call of RootHTTPClient.
func (mr *MockAPICallCloserMockRecorder) RootHTTPClient() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RootHTTPClient", reflect.TypeOf((*MockAPICallCloser)(nil).RootHTTPClient))
}

// MockClientFacade is a mock of ClientFacade interface.
type MockClientFacade struct {
	ctrl     *gomock.Controller
	recorder *MockClientFacadeMockRecorder
}

// MockClientFacadeMockRecorder is the mock recorder for MockClientFacade.
type MockClientFacadeMockRecorder struct {
	mock *MockClientFacade
}

// NewMockClientFacade creates a new mock instance.
func NewMockClientFacade(ctrl *gomock.Controller) *MockClientFacade {
	mock := &MockClientFacade{ctrl: ctrl}
	mock.recorder = &MockClientFacadeMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockClientFacade) EXPECT() *MockClientFacadeMockRecorder {
	return m.recorder
}

// BestAPIVersion mocks base method.
func (m *MockClientFacade) BestAPIVersion() int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BestAPIVersion")
	ret0, _ := ret[0].(int)
	return ret0
}

// BestAPIVersion indicates an expected call of BestAPIVersion.
func (mr *MockClientFacadeMockRecorder) BestAPIVersion() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BestAPIVersion", reflect.TypeOf((*MockClientFacade)(nil).BestAPIVersion))
}

// Close mocks base method.
func (m *MockClientFacade) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close.
func (mr *MockClientFacadeMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockClientFacade)(nil).Close))
}
