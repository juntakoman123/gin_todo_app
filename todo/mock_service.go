// Code generated by MockGen. DO NOT EDIT.
// Source: C:\Users\PC20004\gin_todo_app\todo\service.go

// Package todo is a generated GoMock package.
package todo

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockService is a mock of Service interface.
type MockService struct {
	ctrl     *gomock.Controller
	recorder *MockServiceMockRecorder
}

// MockServiceMockRecorder is the mock recorder for MockService.
type MockServiceMockRecorder struct {
	mock *MockService
}

// NewMockService creates a new mock instance.
func NewMockService(ctrl *gomock.Controller) *MockService {
	mock := &MockService{ctrl: ctrl}
	mock.recorder = &MockServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockService) EXPECT() *MockServiceMockRecorder {
	return m.recorder
}

// AddTask mocks base method.
func (m *MockService) AddTask(task Task) (Task, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddTask", task)
	ret0, _ := ret[0].(Task)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddTask indicates an expected call of AddTask.
func (mr *MockServiceMockRecorder) AddTask(task interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddTask", reflect.TypeOf((*MockService)(nil).AddTask), task)
}

// DeleteTask mocks base method.
func (m *MockService) DeleteTask(id TaskID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteTask", id)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteTask indicates an expected call of DeleteTask.
func (mr *MockServiceMockRecorder) DeleteTask(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteTask", reflect.TypeOf((*MockService)(nil).DeleteTask), id)
}

// GetTasks mocks base method.
func (m *MockService) GetTasks() (Tasks, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTasks")
	ret0, _ := ret[0].(Tasks)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTasks indicates an expected call of GetTasks.
func (mr *MockServiceMockRecorder) GetTasks() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTasks", reflect.TypeOf((*MockService)(nil).GetTasks))
}
