// Code generated by MockGen. DO NOT EDIT.
// Source: handler.go

// Package mock is a generated GoMock package.
package mock

import (
	reflect "reflect"

	gin "github.com/gin-gonic/gin"
	gomock "github.com/golang/mock/gomock"
)

// MockUserHandlerInterface is a mock of UserHandlerInterface interface.
type MockUserHandlerInterface struct {
	ctrl     *gomock.Controller
	recorder *MockUserHandlerInterfaceMockRecorder
}

// MockUserHandlerInterfaceMockRecorder is the mock recorder for MockUserHandlerInterface.
type MockUserHandlerInterfaceMockRecorder struct {
	mock *MockUserHandlerInterface
}

// NewMockUserHandlerInterface creates a new mock instance.
func NewMockUserHandlerInterface(ctrl *gomock.Controller) *MockUserHandlerInterface {
	mock := &MockUserHandlerInterface{ctrl: ctrl}
	mock.recorder = &MockUserHandlerInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserHandlerInterface) EXPECT() *MockUserHandlerInterfaceMockRecorder {
	return m.recorder
}

// UserRegister mocks base method.
func (m *MockUserHandlerInterface) UserRegister(router *gin.RouterGroup) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "UserRegister", router)
}

// UserRegister indicates an expected call of UserRegister.
func (mr *MockUserHandlerInterfaceMockRecorder) UserRegister(router interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UserRegister", reflect.TypeOf((*MockUserHandlerInterface)(nil).UserRegister), router)
}

// MockHabitHandlerInterface is a mock of HabitHandlerInterface interface.
type MockHabitHandlerInterface struct {
	ctrl     *gomock.Controller
	recorder *MockHabitHandlerInterfaceMockRecorder
}

// MockHabitHandlerInterfaceMockRecorder is the mock recorder for MockHabitHandlerInterface.
type MockHabitHandlerInterfaceMockRecorder struct {
	mock *MockHabitHandlerInterface
}

// NewMockHabitHandlerInterface creates a new mock instance.
func NewMockHabitHandlerInterface(ctrl *gomock.Controller) *MockHabitHandlerInterface {
	mock := &MockHabitHandlerInterface{ctrl: ctrl}
	mock.recorder = &MockHabitHandlerInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockHabitHandlerInterface) EXPECT() *MockHabitHandlerInterfaceMockRecorder {
	return m.recorder
}

// HabitRegister mocks base method.
func (m *MockHabitHandlerInterface) HabitRegister(router *gin.RouterGroup) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "HabitRegister", router)
}

// HabitRegister indicates an expected call of HabitRegister.
func (mr *MockHabitHandlerInterfaceMockRecorder) HabitRegister(router interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HabitRegister", reflect.TypeOf((*MockHabitHandlerInterface)(nil).HabitRegister), router)
}