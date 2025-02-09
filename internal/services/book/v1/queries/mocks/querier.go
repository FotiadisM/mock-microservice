// Code generated by mockery v2.50.1. DO NOT EDIT.

package mocks

import (
	context "context"

	uuid "github.com/google/uuid"
	mock "github.com/stretchr/testify/mock"

	queries "github.com/FotiadisM/mock-microservice/internal/services/book/v1/queries"
)

// MockQuerier is an autogenerated mock type for the Querier type
type MockQuerier struct {
	mock.Mock
}

type MockQuerier_Expecter struct {
	mock *mock.Mock
}

func (_m *MockQuerier) EXPECT() *MockQuerier_Expecter {
	return &MockQuerier_Expecter{mock: &_m.Mock}
}

// CreateAuthor provides a mock function with given fields: ctx, arg
func (_m *MockQuerier) CreateAuthor(ctx context.Context, arg queries.CreateAuthorParams) error {
	ret := _m.Called(ctx, arg)

	if len(ret) == 0 {
		panic("no return value specified for CreateAuthor")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, queries.CreateAuthorParams) error); ok {
		r0 = rf(ctx, arg)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockQuerier_CreateAuthor_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateAuthor'
type MockQuerier_CreateAuthor_Call struct {
	*mock.Call
}

// CreateAuthor is a helper method to define mock.On call
//   - ctx context.Context
//   - arg queries.CreateAuthorParams
func (_e *MockQuerier_Expecter) CreateAuthor(ctx interface{}, arg interface{}) *MockQuerier_CreateAuthor_Call {
	return &MockQuerier_CreateAuthor_Call{Call: _e.mock.On("CreateAuthor", ctx, arg)}
}

func (_c *MockQuerier_CreateAuthor_Call) Run(run func(ctx context.Context, arg queries.CreateAuthorParams)) *MockQuerier_CreateAuthor_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(queries.CreateAuthorParams))
	})
	return _c
}

func (_c *MockQuerier_CreateAuthor_Call) Return(_a0 error) *MockQuerier_CreateAuthor_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockQuerier_CreateAuthor_Call) RunAndReturn(run func(context.Context, queries.CreateAuthorParams) error) *MockQuerier_CreateAuthor_Call {
	_c.Call.Return(run)
	return _c
}

// CreateBook provides a mock function with given fields: ctx, arg
func (_m *MockQuerier) CreateBook(ctx context.Context, arg queries.CreateBookParams) error {
	ret := _m.Called(ctx, arg)

	if len(ret) == 0 {
		panic("no return value specified for CreateBook")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, queries.CreateBookParams) error); ok {
		r0 = rf(ctx, arg)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockQuerier_CreateBook_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateBook'
type MockQuerier_CreateBook_Call struct {
	*mock.Call
}

// CreateBook is a helper method to define mock.On call
//   - ctx context.Context
//   - arg queries.CreateBookParams
func (_e *MockQuerier_Expecter) CreateBook(ctx interface{}, arg interface{}) *MockQuerier_CreateBook_Call {
	return &MockQuerier_CreateBook_Call{Call: _e.mock.On("CreateBook", ctx, arg)}
}

func (_c *MockQuerier_CreateBook_Call) Run(run func(ctx context.Context, arg queries.CreateBookParams)) *MockQuerier_CreateBook_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(queries.CreateBookParams))
	})
	return _c
}

func (_c *MockQuerier_CreateBook_Call) Return(_a0 error) *MockQuerier_CreateBook_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockQuerier_CreateBook_Call) RunAndReturn(run func(context.Context, queries.CreateBookParams) error) *MockQuerier_CreateBook_Call {
	_c.Call.Return(run)
	return _c
}

// DeleteAuthor provides a mock function with given fields: ctx, id
func (_m *MockQuerier) DeleteAuthor(ctx context.Context, id uuid.UUID) error {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for DeleteAuthor")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) error); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockQuerier_DeleteAuthor_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DeleteAuthor'
type MockQuerier_DeleteAuthor_Call struct {
	*mock.Call
}

// DeleteAuthor is a helper method to define mock.On call
//   - ctx context.Context
//   - id uuid.UUID
func (_e *MockQuerier_Expecter) DeleteAuthor(ctx interface{}, id interface{}) *MockQuerier_DeleteAuthor_Call {
	return &MockQuerier_DeleteAuthor_Call{Call: _e.mock.On("DeleteAuthor", ctx, id)}
}

func (_c *MockQuerier_DeleteAuthor_Call) Run(run func(ctx context.Context, id uuid.UUID)) *MockQuerier_DeleteAuthor_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(uuid.UUID))
	})
	return _c
}

func (_c *MockQuerier_DeleteAuthor_Call) Return(_a0 error) *MockQuerier_DeleteAuthor_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockQuerier_DeleteAuthor_Call) RunAndReturn(run func(context.Context, uuid.UUID) error) *MockQuerier_DeleteAuthor_Call {
	_c.Call.Return(run)
	return _c
}

// DeleteBook provides a mock function with given fields: ctx, id
func (_m *MockQuerier) DeleteBook(ctx context.Context, id uuid.UUID) error {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for DeleteBook")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) error); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockQuerier_DeleteBook_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DeleteBook'
type MockQuerier_DeleteBook_Call struct {
	*mock.Call
}

// DeleteBook is a helper method to define mock.On call
//   - ctx context.Context
//   - id uuid.UUID
func (_e *MockQuerier_Expecter) DeleteBook(ctx interface{}, id interface{}) *MockQuerier_DeleteBook_Call {
	return &MockQuerier_DeleteBook_Call{Call: _e.mock.On("DeleteBook", ctx, id)}
}

func (_c *MockQuerier_DeleteBook_Call) Run(run func(ctx context.Context, id uuid.UUID)) *MockQuerier_DeleteBook_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(uuid.UUID))
	})
	return _c
}

func (_c *MockQuerier_DeleteBook_Call) Return(_a0 error) *MockQuerier_DeleteBook_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockQuerier_DeleteBook_Call) RunAndReturn(run func(context.Context, uuid.UUID) error) *MockQuerier_DeleteBook_Call {
	_c.Call.Return(run)
	return _c
}

// GetAuthor provides a mock function with given fields: ctx, id
func (_m *MockQuerier) GetAuthor(ctx context.Context, id uuid.UUID) (queries.Author, error) {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for GetAuthor")
	}

	var r0 queries.Author
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) (queries.Author, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) queries.Author); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(queries.Author)
	}

	if rf, ok := ret.Get(1).(func(context.Context, uuid.UUID) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockQuerier_GetAuthor_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetAuthor'
type MockQuerier_GetAuthor_Call struct {
	*mock.Call
}

// GetAuthor is a helper method to define mock.On call
//   - ctx context.Context
//   - id uuid.UUID
func (_e *MockQuerier_Expecter) GetAuthor(ctx interface{}, id interface{}) *MockQuerier_GetAuthor_Call {
	return &MockQuerier_GetAuthor_Call{Call: _e.mock.On("GetAuthor", ctx, id)}
}

func (_c *MockQuerier_GetAuthor_Call) Run(run func(ctx context.Context, id uuid.UUID)) *MockQuerier_GetAuthor_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(uuid.UUID))
	})
	return _c
}

func (_c *MockQuerier_GetAuthor_Call) Return(_a0 queries.Author, _a1 error) *MockQuerier_GetAuthor_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockQuerier_GetAuthor_Call) RunAndReturn(run func(context.Context, uuid.UUID) (queries.Author, error)) *MockQuerier_GetAuthor_Call {
	_c.Call.Return(run)
	return _c
}

// GetBook provides a mock function with given fields: ctx, id
func (_m *MockQuerier) GetBook(ctx context.Context, id uuid.UUID) (queries.Book, error) {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for GetBook")
	}

	var r0 queries.Book
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) (queries.Book, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) queries.Book); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(queries.Book)
	}

	if rf, ok := ret.Get(1).(func(context.Context, uuid.UUID) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockQuerier_GetBook_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetBook'
type MockQuerier_GetBook_Call struct {
	*mock.Call
}

// GetBook is a helper method to define mock.On call
//   - ctx context.Context
//   - id uuid.UUID
func (_e *MockQuerier_Expecter) GetBook(ctx interface{}, id interface{}) *MockQuerier_GetBook_Call {
	return &MockQuerier_GetBook_Call{Call: _e.mock.On("GetBook", ctx, id)}
}

func (_c *MockQuerier_GetBook_Call) Run(run func(ctx context.Context, id uuid.UUID)) *MockQuerier_GetBook_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(uuid.UUID))
	})
	return _c
}

func (_c *MockQuerier_GetBook_Call) Return(_a0 queries.Book, _a1 error) *MockQuerier_GetBook_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockQuerier_GetBook_Call) RunAndReturn(run func(context.Context, uuid.UUID) (queries.Book, error)) *MockQuerier_GetBook_Call {
	_c.Call.Return(run)
	return _c
}

// ListAuthors provides a mock function with given fields: ctx
func (_m *MockQuerier) ListAuthors(ctx context.Context) ([]queries.Author, error) {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for ListAuthors")
	}

	var r0 []queries.Author
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) ([]queries.Author, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) []queries.Author); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]queries.Author)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockQuerier_ListAuthors_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ListAuthors'
type MockQuerier_ListAuthors_Call struct {
	*mock.Call
}

// ListAuthors is a helper method to define mock.On call
//   - ctx context.Context
func (_e *MockQuerier_Expecter) ListAuthors(ctx interface{}) *MockQuerier_ListAuthors_Call {
	return &MockQuerier_ListAuthors_Call{Call: _e.mock.On("ListAuthors", ctx)}
}

func (_c *MockQuerier_ListAuthors_Call) Run(run func(ctx context.Context)) *MockQuerier_ListAuthors_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *MockQuerier_ListAuthors_Call) Return(_a0 []queries.Author, _a1 error) *MockQuerier_ListAuthors_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockQuerier_ListAuthors_Call) RunAndReturn(run func(context.Context) ([]queries.Author, error)) *MockQuerier_ListAuthors_Call {
	_c.Call.Return(run)
	return _c
}

// ListBooks provides a mock function with given fields: ctx
func (_m *MockQuerier) ListBooks(ctx context.Context) ([]queries.Book, error) {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for ListBooks")
	}

	var r0 []queries.Book
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) ([]queries.Book, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) []queries.Book); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]queries.Book)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockQuerier_ListBooks_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ListBooks'
type MockQuerier_ListBooks_Call struct {
	*mock.Call
}

// ListBooks is a helper method to define mock.On call
//   - ctx context.Context
func (_e *MockQuerier_Expecter) ListBooks(ctx interface{}) *MockQuerier_ListBooks_Call {
	return &MockQuerier_ListBooks_Call{Call: _e.mock.On("ListBooks", ctx)}
}

func (_c *MockQuerier_ListBooks_Call) Run(run func(ctx context.Context)) *MockQuerier_ListBooks_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *MockQuerier_ListBooks_Call) Return(_a0 []queries.Book, _a1 error) *MockQuerier_ListBooks_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockQuerier_ListBooks_Call) RunAndReturn(run func(context.Context) ([]queries.Book, error)) *MockQuerier_ListBooks_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockQuerier creates a new instance of MockQuerier. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockQuerier(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockQuerier {
	mock := &MockQuerier{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
