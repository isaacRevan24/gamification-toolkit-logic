// Code generated by MockGen. DO NOT EDIT.
// Source: connection.go

// Package mock is a generated GoMock package.
package mock

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockUserRepository is a mock of UserRepository interface.
type MockUserRepository struct {
	ctrl     *gomock.Controller
	recorder *MockUserRepositoryMockRecorder
}

// MockUserRepositoryMockRecorder is the mock recorder for MockUserRepository.
type MockUserRepositoryMockRecorder struct {
	mock *MockUserRepository
}

// NewMockUserRepository creates a new mock instance.
func NewMockUserRepository(ctrl *gomock.Controller) *MockUserRepository {
	mock := &MockUserRepository{ctrl: ctrl}
	mock.recorder = &MockUserRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserRepository) EXPECT() *MockUserRepositoryMockRecorder {
	return m.recorder
}

// SignUpRepository mocks base method.
func (m *MockUserRepository) SignUpRepository(userId string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SignUpRepository", userId)
	ret0, _ := ret[0].(error)
	return ret0
}

// SignUpRepository indicates an expected call of SignUpRepository.
func (mr *MockUserRepositoryMockRecorder) SignUpRepository(userId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SignUpRepository", reflect.TypeOf((*MockUserRepository)(nil).SignUpRepository), userId)
}

// MockHabitRepository is a mock of HabitRepository interface.
type MockHabitRepository struct {
	ctrl     *gomock.Controller
	recorder *MockHabitRepositoryMockRecorder
}

// MockHabitRepositoryMockRecorder is the mock recorder for MockHabitRepository.
type MockHabitRepositoryMockRecorder struct {
	mock *MockHabitRepository
}

// NewMockHabitRepository creates a new mock instance.
func NewMockHabitRepository(ctrl *gomock.Controller) *MockHabitRepository {
	mock := &MockHabitRepository{ctrl: ctrl}
	mock.recorder = &MockHabitRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockHabitRepository) EXPECT() *MockHabitRepositoryMockRecorder {
	return m.recorder
}

// AddNewHabitRepository mocks base method.
func (m *MockHabitRepository) AddNewHabitRepository(userId, name, description, condition string, repetition int) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddNewHabitRepository", userId, name, description, condition, repetition)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddNewHabitRepository indicates an expected call of AddNewHabitRepository.
func (mr *MockHabitRepositoryMockRecorder) AddNewHabitRepository(userId, name, description, condition, repetition interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddNewHabitRepository", reflect.TypeOf((*MockHabitRepository)(nil).AddNewHabitRepository), userId, name, description, condition, repetition)
}

// DeleteHabitRepository mocks base method.
func (m *MockHabitRepository) DeleteHabitRepository(userId, habitId string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteHabitRepository", userId, habitId)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteHabitRepository indicates an expected call of DeleteHabitRepository.
func (mr *MockHabitRepositoryMockRecorder) DeleteHabitRepository(userId, habitId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteHabitRepository", reflect.TypeOf((*MockHabitRepository)(nil).DeleteHabitRepository), userId, habitId)
}
