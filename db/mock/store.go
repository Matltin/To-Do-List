// Code generated by MockGen. DO NOT EDIT.
// Source: to_do_list/db/sqlc (interfaces: Store)

// Package mock_sqlc is a generated GoMock package.
package mock_sqlc

import (
	context "context"
	reflect "reflect"
	sqlc "to_do_list/db/sqlc"

	gomock "github.com/golang/mock/gomock"
)

// MockStore is a mock of Store interface.
type MockStore struct {
	ctrl     *gomock.Controller
	recorder *MockStoreMockRecorder
}

// MockStoreMockRecorder is the mock recorder for MockStore.
type MockStoreMockRecorder struct {
	mock *MockStore
}

// NewMockStore creates a new mock instance.
func NewMockStore(ctrl *gomock.Controller) *MockStore {
	mock := &MockStore{ctrl: ctrl}
	mock.recorder = &MockStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStore) EXPECT() *MockStoreMockRecorder {
	return m.recorder
}

// CheckUserExists mocks base method.
func (m *MockStore) CheckUserExists(arg0 context.Context, arg1 string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckUserExists", arg0, arg1)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CheckUserExists indicates an expected call of CheckUserExists.
func (mr *MockStoreMockRecorder) CheckUserExists(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckUserExists", reflect.TypeOf((*MockStore)(nil).CheckUserExists), arg0, arg1)
}

// CreateTodo mocks base method.
func (m *MockStore) CreateTodo(arg0 context.Context, arg1 sqlc.CreateTodoParams) (sqlc.Todo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateTodo", arg0, arg1)
	ret0, _ := ret[0].(sqlc.Todo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateTodo indicates an expected call of CreateTodo.
func (mr *MockStoreMockRecorder) CreateTodo(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateTodo", reflect.TypeOf((*MockStore)(nil).CreateTodo), arg0, arg1)
}

// CreateUser mocks base method.
func (m *MockStore) CreateUser(arg0 context.Context, arg1 sqlc.CreateUserParams) (sqlc.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUser", arg0, arg1)
	ret0, _ := ret[0].(sqlc.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateUser indicates an expected call of CreateUser.
func (mr *MockStoreMockRecorder) CreateUser(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser", reflect.TypeOf((*MockStore)(nil).CreateUser), arg0, arg1)
}

// DeleteTodo mocks base method.
func (m *MockStore) DeleteTodo(arg0 context.Context, arg1 sqlc.DeleteTodoParams) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteTodo", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteTodo indicates an expected call of DeleteTodo.
func (mr *MockStoreMockRecorder) DeleteTodo(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteTodo", reflect.TypeOf((*MockStore)(nil).DeleteTodo), arg0, arg1)
}

// GetTodoByID mocks base method.
func (m *MockStore) GetTodoByID(arg0 context.Context, arg1 int64) (sqlc.Todo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTodoByID", arg0, arg1)
	ret0, _ := ret[0].(sqlc.Todo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTodoByID indicates an expected call of GetTodoByID.
func (mr *MockStoreMockRecorder) GetTodoByID(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTodoByID", reflect.TypeOf((*MockStore)(nil).GetTodoByID), arg0, arg1)
}

// GetTodosByID mocks base method.
func (m *MockStore) GetTodosByID(arg0 context.Context, arg1 sqlc.GetTodosByIDParams) ([]sqlc.Todo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTodosByID", arg0, arg1)
	ret0, _ := ret[0].([]sqlc.Todo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTodosByID indicates an expected call of GetTodosByID.
func (mr *MockStoreMockRecorder) GetTodosByID(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTodosByID", reflect.TypeOf((*MockStore)(nil).GetTodosByID), arg0, arg1)
}

// GetUser mocks base method.
func (m *MockStore) GetUser(arg0 context.Context, arg1 string) (sqlc.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUser", arg0, arg1)
	ret0, _ := ret[0].(sqlc.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUser indicates an expected call of GetUser.
func (mr *MockStoreMockRecorder) GetUser(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUser", reflect.TypeOf((*MockStore)(nil).GetUser), arg0, arg1)
}

// UpdateTodo mocks base method.
func (m *MockStore) UpdateTodo(arg0 context.Context, arg1 sqlc.UpdateTodoParams) (sqlc.Todo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateTodo", arg0, arg1)
	ret0, _ := ret[0].(sqlc.Todo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateTodo indicates an expected call of UpdateTodo.
func (mr *MockStoreMockRecorder) UpdateTodo(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateTodo", reflect.TypeOf((*MockStore)(nil).UpdateTodo), arg0, arg1)
}
