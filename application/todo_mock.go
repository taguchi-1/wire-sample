// Code generated by MockGen. DO NOT EDIT.
// Source: /Users/b09872/dev/src/github.com/taguchi-1/wire-sample/application/todo.go

// Package application is a generated GoMock package.
package application

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	entity "github.com/taguchi-1/wire-sample/domain/entity"
)

// MockTodo is a mock of Todo interface
type MockTodo struct {
	ctrl     *gomock.Controller
	recorder *MockTodoMockRecorder
}

// MockTodoMockRecorder is the mock recorder for MockTodo
type MockTodoMockRecorder struct {
	mock *MockTodo
}

// NewMockTodo creates a new mock instance
func NewMockTodo(ctrl *gomock.Controller) *MockTodo {
	mock := &MockTodo{ctrl: ctrl}
	mock.recorder = &MockTodoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockTodo) EXPECT() *MockTodoMockRecorder {
	return m.recorder
}

// Get mocks base method
func (m *MockTodo) Get(ctx context.Context, req *entity.TodoGetRequest) (*entity.TodoResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", ctx, req)
	ret0, _ := ret[0].(*entity.TodoResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get
func (mr *MockTodoMockRecorder) Get(ctx, req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockTodo)(nil).Get), ctx, req)
}
