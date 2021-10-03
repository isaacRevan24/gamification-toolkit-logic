// Code generated by MockGen. DO NOT EDIT.
// Source: router.go

// Package mock is a generated GoMock package.
package mock

import (
	reflect "reflect"

	gin "github.com/gin-gonic/gin"
	gomock "github.com/golang/mock/gomock"
)

// MockRouterRegistersInterface is a mock of RouterRegistersInterface interface.
type MockRouterRegistersInterface struct {
	ctrl     *gomock.Controller
	recorder *MockRouterRegistersInterfaceMockRecorder
}

// MockRouterRegistersInterfaceMockRecorder is the mock recorder for MockRouterRegistersInterface.
type MockRouterRegistersInterfaceMockRecorder struct {
	mock *MockRouterRegistersInterface
}

// NewMockRouterRegistersInterface creates a new mock instance.
func NewMockRouterRegistersInterface(ctrl *gomock.Controller) *MockRouterRegistersInterface {
	mock := &MockRouterRegistersInterface{ctrl: ctrl}
	mock.recorder = &MockRouterRegistersInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRouterRegistersInterface) EXPECT() *MockRouterRegistersInterfaceMockRecorder {
	return m.recorder
}

// HabitRegister mocks base method.
func (m *MockRouterRegistersInterface) HabitRegister(router *gin.RouterGroup) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "HabitRegister", router)
}

// HabitRegister indicates an expected call of HabitRegister.
func (mr *MockRouterRegistersInterfaceMockRecorder) HabitRegister(router interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HabitRegister", reflect.TypeOf((*MockRouterRegistersInterface)(nil).HabitRegister), router)
}

// UserRegister mocks base method.
func (m *MockRouterRegistersInterface) UserRegister(router *gin.RouterGroup) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "UserRegister", router)
}

// UserRegister indicates an expected call of UserRegister.
func (mr *MockRouterRegistersInterfaceMockRecorder) UserRegister(router interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UserRegister", reflect.TypeOf((*MockRouterRegistersInterface)(nil).UserRegister), router)
}