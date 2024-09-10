// Code generated by mockery v2.20.0. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
)

// API is an autogenerated mock type for the API type
type API struct {
	mock.Mock
}

type API_Expecter struct {
	mock *mock.Mock
}

func (_m *API) EXPECT() *API_Expecter {
	return &API_Expecter{mock: &_m.Mock}
}

// Ingest provides a mock function with given fields: ctx, clusterName, runID
func (_m *API) Ingest(ctx context.Context, clusterName string, runID string) error {
	ret := _m.Called(ctx, clusterName, runID)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string) error); ok {
		r0 = rf(ctx, clusterName, runID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// API_Ingest_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Ingest'
type API_Ingest_Call struct {
	*mock.Call
}

// Ingest is a helper method to define mock.On call
//   - ctx context.Context
//   - clusterName string
//   - runID string
func (_e *API_Expecter) Ingest(ctx interface{}, clusterName interface{}, runID interface{}) *API_Ingest_Call {
	return &API_Ingest_Call{Call: _e.mock.On("Ingest", ctx, clusterName, runID)}
}

func (_c *API_Ingest_Call) Run(run func(ctx context.Context, clusterName string, runID string)) *API_Ingest_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(string))
	})
	return _c
}

func (_c *API_Ingest_Call) Return(_a0 error) *API_Ingest_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *API_Ingest_Call) RunAndReturn(run func(context.Context, string, string) error) *API_Ingest_Call {
	_c.Call.Return(run)
	return _c
}

// Notify provides a mock function with given fields: ctx, clusterName, runID
func (_m *API) Notify(ctx context.Context, clusterName string, runID string) error {
	ret := _m.Called(ctx, clusterName, runID)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string) error); ok {
		r0 = rf(ctx, clusterName, runID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// API_Notify_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Notify'
type API_Notify_Call struct {
	*mock.Call
}

// Notify is a helper method to define mock.On call
//   - ctx context.Context
//   - clusterName string
//   - runID string
func (_e *API_Expecter) Notify(ctx interface{}, clusterName interface{}, runID interface{}) *API_Notify_Call {
	return &API_Notify_Call{Call: _e.mock.On("Notify", ctx, clusterName, runID)}
}

func (_c *API_Notify_Call) Run(run func(ctx context.Context, clusterName string, runID string)) *API_Notify_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(string))
	})
	return _c
}

func (_c *API_Notify_Call) Return(_a0 error) *API_Notify_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *API_Notify_Call) RunAndReturn(run func(context.Context, string, string) error) *API_Notify_Call {
	_c.Call.Return(run)
	return _c
}

// isAlreadyIngestedInGraph provides a mock function with given fields: ctx, clusterName, runID
func (_m *API) isAlreadyIngestedInGraph(ctx context.Context, clusterName string, runID string) (bool, error) {
	ret := _m.Called(ctx, clusterName, runID)

	var r0 bool
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string) (bool, error)); ok {
		return rf(ctx, clusterName, runID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, string) bool); ok {
		r0 = rf(ctx, clusterName, runID)
	} else {
		r0 = ret.Get(0).(bool)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, string) error); ok {
		r1 = rf(ctx, clusterName, runID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// API_isAlreadyIngestedInGraph_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'isAlreadyIngestedInGraph'
type API_isAlreadyIngestedInGraph_Call struct {
	*mock.Call
}

// isAlreadyIngestedInGraph is a helper method to define mock.On call
//   - ctx context.Context
//   - clusterName string
//   - runID string
func (_e *API_Expecter) isAlreadyIngestedInGraph(ctx interface{}, clusterName interface{}, runID interface{}) *API_isAlreadyIngestedInGraph_Call {
	return &API_isAlreadyIngestedInGraph_Call{Call: _e.mock.On("isAlreadyIngestedInGraph", ctx, clusterName, runID)}
}

func (_c *API_isAlreadyIngestedInGraph_Call) Run(run func(ctx context.Context, clusterName string, runID string)) *API_isAlreadyIngestedInGraph_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(string))
	})
	return _c
}

func (_c *API_isAlreadyIngestedInGraph_Call) Return(_a0 bool, _a1 error) *API_isAlreadyIngestedInGraph_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *API_isAlreadyIngestedInGraph_Call) RunAndReturn(run func(context.Context, string, string) (bool, error)) *API_isAlreadyIngestedInGraph_Call {
	_c.Call.Return(run)
	return _c
}

type mockConstructorTestingTNewAPI interface {
	mock.TestingT
	Cleanup(func())
}

// NewAPI creates a new instance of API. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewAPI(t mockConstructorTestingTNewAPI) *API {
	mock := &API{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
