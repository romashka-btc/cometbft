// Code generated by mockery. DO NOT EDIT.

package mocks

import (
	context "context"

	abcicli "github.com/cometbft/cometbft/abci/client"

	mock "github.com/stretchr/testify/mock"

	v1beta4 "github.com/cometbft/cometbft/api/cometbft/abci/v1beta4"
)

// AppConnMempool is an autogenerated mock type for the AppConnMempool type
type AppConnMempool struct {
	mock.Mock
}

// CheckTx provides a mock function with given fields: _a0, _a1
func (_m *AppConnMempool) CheckTx(_a0 context.Context, _a1 *v1beta4.CheckTxRequest) (*v1beta4.CheckTxResponse, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *v1beta4.CheckTxResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *v1beta4.CheckTxRequest) (*v1beta4.CheckTxResponse, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *v1beta4.CheckTxRequest) *v1beta4.CheckTxResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v1beta4.CheckTxResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *v1beta4.CheckTxRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CheckTxAsync provides a mock function with given fields: _a0, _a1
func (_m *AppConnMempool) CheckTxAsync(_a0 context.Context, _a1 *v1beta4.CheckTxRequest) (*abcicli.ReqRes, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *abcicli.ReqRes
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *v1beta4.CheckTxRequest) (*abcicli.ReqRes, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *v1beta4.CheckTxRequest) *abcicli.ReqRes); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*abcicli.ReqRes)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *v1beta4.CheckTxRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Error provides a mock function with given fields:
func (_m *AppConnMempool) Error() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Flush provides a mock function with given fields: _a0
func (_m *AppConnMempool) Flush(_a0 context.Context) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SetResponseCallback provides a mock function with given fields: _a0
func (_m *AppConnMempool) SetResponseCallback(_a0 abcicli.Callback) {
	_m.Called(_a0)
}

// NewAppConnMempool creates a new instance of AppConnMempool. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewAppConnMempool(t interface {
	mock.TestingT
	Cleanup(func())
}) *AppConnMempool {
	mock := &AppConnMempool{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
