// Code generated by mockery v2.20.0. DO NOT EDIT.

package mocks

import (
	context "context"

	cache "github.com/DataDog/KubeHound/pkg/kubehound/storage/cache"

	mock "github.com/stretchr/testify/mock"
)

// CacheProvider is an autogenerated mock type for the CacheProvider type
type CacheProvider struct {
	mock.Mock
}

type CacheProvider_Expecter struct {
	mock *mock.Mock
}

func (_m *CacheProvider) EXPECT() *CacheProvider_Expecter {
	return &CacheProvider_Expecter{mock: &_m.Mock}
}

// BulkWriter provides a mock function with given fields: ctx
func (_m *CacheProvider) BulkWriter(ctx context.Context) (cache.AsyncWriter, error) {
	ret := _m.Called(ctx)

	var r0 cache.AsyncWriter
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) (cache.AsyncWriter, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) cache.AsyncWriter); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(cache.AsyncWriter)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CacheProvider_BulkWriter_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'BulkWriter'
type CacheProvider_BulkWriter_Call struct {
	*mock.Call
}

// BulkWriter is a helper method to define mock.On call
//   - ctx context.Context
func (_e *CacheProvider_Expecter) BulkWriter(ctx interface{}) *CacheProvider_BulkWriter_Call {
	return &CacheProvider_BulkWriter_Call{Call: _e.mock.On("BulkWriter", ctx)}
}

func (_c *CacheProvider_BulkWriter_Call) Run(run func(ctx context.Context)) *CacheProvider_BulkWriter_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *CacheProvider_BulkWriter_Call) Return(_a0 cache.AsyncWriter, _a1 error) *CacheProvider_BulkWriter_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *CacheProvider_BulkWriter_Call) RunAndReturn(run func(context.Context) (cache.AsyncWriter, error)) *CacheProvider_BulkWriter_Call {
	_c.Call.Return(run)
	return _c
}

// Close provides a mock function with given fields: ctx
func (_m *CacheProvider) Close(ctx context.Context) error {
	ret := _m.Called(ctx)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context) error); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CacheProvider_Close_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Close'
type CacheProvider_Close_Call struct {
	*mock.Call
}

// Close is a helper method to define mock.On call
//   - ctx context.Context
func (_e *CacheProvider_Expecter) Close(ctx interface{}) *CacheProvider_Close_Call {
	return &CacheProvider_Close_Call{Call: _e.mock.On("Close", ctx)}
}

func (_c *CacheProvider_Close_Call) Run(run func(ctx context.Context)) *CacheProvider_Close_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *CacheProvider_Close_Call) Return(_a0 error) *CacheProvider_Close_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *CacheProvider_Close_Call) RunAndReturn(run func(context.Context) error) *CacheProvider_Close_Call {
	_c.Call.Return(run)
	return _c
}

// Get provides a mock function with given fields: ctx, key
func (_m *CacheProvider) Get(ctx context.Context, key cache.CacheKey) (string, error) {
	ret := _m.Called(ctx, key)

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, cache.CacheKey) (string, error)); ok {
		return rf(ctx, key)
	}
	if rf, ok := ret.Get(0).(func(context.Context, cache.CacheKey) string); ok {
		r0 = rf(ctx, key)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(context.Context, cache.CacheKey) error); ok {
		r1 = rf(ctx, key)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CacheProvider_Get_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Get'
type CacheProvider_Get_Call struct {
	*mock.Call
}

// Get is a helper method to define mock.On call
//   - ctx context.Context
//   - key cache.CacheKey
func (_e *CacheProvider_Expecter) Get(ctx interface{}, key interface{}) *CacheProvider_Get_Call {
	return &CacheProvider_Get_Call{Call: _e.mock.On("Get", ctx, key)}
}

func (_c *CacheProvider_Get_Call) Run(run func(ctx context.Context, key cache.CacheKey)) *CacheProvider_Get_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(cache.CacheKey))
	})
	return _c
}

func (_c *CacheProvider_Get_Call) Return(_a0 string, _a1 error) *CacheProvider_Get_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *CacheProvider_Get_Call) RunAndReturn(run func(context.Context, cache.CacheKey) (string, error)) *CacheProvider_Get_Call {
	_c.Call.Return(run)
	return _c
}

// HealthCheck provides a mock function with given fields: ctx
func (_m *CacheProvider) HealthCheck(ctx context.Context) (bool, error) {
	ret := _m.Called(ctx)

	var r0 bool
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) (bool, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) bool); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Get(0).(bool)
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CacheProvider_HealthCheck_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'HealthCheck'
type CacheProvider_HealthCheck_Call struct {
	*mock.Call
}

// HealthCheck is a helper method to define mock.On call
//   - ctx context.Context
func (_e *CacheProvider_Expecter) HealthCheck(ctx interface{}) *CacheProvider_HealthCheck_Call {
	return &CacheProvider_HealthCheck_Call{Call: _e.mock.On("HealthCheck", ctx)}
}

func (_c *CacheProvider_HealthCheck_Call) Run(run func(ctx context.Context)) *CacheProvider_HealthCheck_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *CacheProvider_HealthCheck_Call) Return(_a0 bool, _a1 error) *CacheProvider_HealthCheck_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *CacheProvider_HealthCheck_Call) RunAndReturn(run func(context.Context) (bool, error)) *CacheProvider_HealthCheck_Call {
	_c.Call.Return(run)
	return _c
}

// Name provides a mock function with given fields:
func (_m *CacheProvider) Name() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// CacheProvider_Name_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Name'
type CacheProvider_Name_Call struct {
	*mock.Call
}

// Name is a helper method to define mock.On call
func (_e *CacheProvider_Expecter) Name() *CacheProvider_Name_Call {
	return &CacheProvider_Name_Call{Call: _e.mock.On("Name")}
}

func (_c *CacheProvider_Name_Call) Run(run func()) *CacheProvider_Name_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *CacheProvider_Name_Call) Return(_a0 string) *CacheProvider_Name_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *CacheProvider_Name_Call) RunAndReturn(run func() string) *CacheProvider_Name_Call {
	_c.Call.Return(run)
	return _c
}

type mockConstructorTestingTNewCacheProvider interface {
	mock.TestingT
	Cleanup(func())
}

// NewCacheProvider creates a new instance of CacheProvider. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewCacheProvider(t mockConstructorTestingTNewCacheProvider) *CacheProvider {
	mock := &CacheProvider{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}