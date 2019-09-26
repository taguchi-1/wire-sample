// Code generated by MockGen. DO NOT EDIT.
// Source: /Users/b09872/dev/src/github.com/taguchi-1/wire-sample/pkg/domain/service/user.go

// Package service is a generated GoMock package.
package service

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	entity "github.com/taguchi-1/wire-sample/pkg/domain/entity"
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
func (m *MockTodo) Get(ctx context.Context, id string) (*entity.Todo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", ctx, id)
	ret0, _ := ret[0].(*entity.Todo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get
func (mr *MockTodoMockRecorder) Get(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockTodo)(nil).Get), ctx, id)
}