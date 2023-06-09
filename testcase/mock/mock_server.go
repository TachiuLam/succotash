// Code generated by MockGen. DO NOT EDIT.
// Source: server.go

// Package mock is a generated GoMock package.
package mock

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockWebServer is a mock of WebServer interface.
type MockWebServer struct {
	ctrl     *gomock.Controller
	recorder *MockWebServerMockRecorder
}

// MockWebServerMockRecorder is the mock recorder for MockWebServer.
type MockWebServerMockRecorder struct {
	mock *MockWebServer
}

// NewMockWebServer creates a new mock instance.
func NewMockWebServer(ctrl *gomock.Controller) *MockWebServer {
	mock := &MockWebServer{ctrl: ctrl}
	mock.recorder = &MockWebServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockWebServer) EXPECT() *MockWebServerMockRecorder {
	return m.recorder
}

// Login mocks base method.
func (m *MockWebServer) Login(username, password string) int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Login", username, password)
	ret0, _ := ret[0].(int)
	return ret0
}

// Login indicates an expected call of Login.
func (mr *MockWebServerMockRecorder) Login(username, password interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Login", reflect.TypeOf((*MockWebServer)(nil).Login), username, password)
}

// Logout mocks base method.
func (m *MockWebServer) Logout() int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Logout")
	ret0, _ := ret[0].(int)
	return ret0
}

// Logout indicates an expected call of Logout.
func (mr *MockWebServerMockRecorder) Logout() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Logout", reflect.TypeOf((*MockWebServer)(nil).Logout))
}
