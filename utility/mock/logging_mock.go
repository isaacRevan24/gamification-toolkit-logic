// Code generated by MockGen. DO NOT EDIT.
// Source: logging.go

// Package mock is a generated GoMock package.
package mock

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockLoggingInterface is a mock of LoggingInterface interface.
type MockLoggingInterface struct {
	ctrl     *gomock.Controller
	recorder *MockLoggingInterfaceMockRecorder
}

// MockLoggingInterfaceMockRecorder is the mock recorder for MockLoggingInterface.
type MockLoggingInterfaceMockRecorder struct {
	mock *MockLoggingInterface
}

// NewMockLoggingInterface creates a new mock instance.
func NewMockLoggingInterface(ctrl *gomock.Controller) *MockLoggingInterface {
	mock := &MockLoggingInterface{ctrl: ctrl}
	mock.recorder = &MockLoggingInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockLoggingInterface) EXPECT() *MockLoggingInterfaceMockRecorder {
	return m.recorder
}

// LogDebug mocks base method.
func (m *MockLoggingInterface) LogDebug(message string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "LogDebug", message)
}

// LogDebug indicates an expected call of LogDebug.
func (mr *MockLoggingInterfaceMockRecorder) LogDebug(message interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LogDebug", reflect.TypeOf((*MockLoggingInterface)(nil).LogDebug), message)
}

// LogError mocks base method.
func (m *MockLoggingInterface) LogError(err error) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "LogError", err)
}

// LogError indicates an expected call of LogError.
func (mr *MockLoggingInterfaceMockRecorder) LogError(err interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LogError", reflect.TypeOf((*MockLoggingInterface)(nil).LogError), err)
}

// LogInfo mocks base method.
func (m *MockLoggingInterface) LogInfo(message string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "LogInfo", message)
}

// LogInfo indicates an expected call of LogInfo.
func (mr *MockLoggingInterfaceMockRecorder) LogInfo(message interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LogInfo", reflect.TypeOf((*MockLoggingInterface)(nil).LogInfo), message)
}

// LogWarning mocks base method.
func (m *MockLoggingInterface) LogWarning(message interface{}) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "LogWarning", message)
}

// LogWarning indicates an expected call of LogWarning.
func (mr *MockLoggingInterfaceMockRecorder) LogWarning(message interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LogWarning", reflect.TypeOf((*MockLoggingInterface)(nil).LogWarning), message)
}
