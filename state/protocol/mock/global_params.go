// Code generated by mockery v2.43.2. DO NOT EDIT.

package mock

import (
	flow "github.com/onflow/flow-go/model/flow"
	mock "github.com/stretchr/testify/mock"
)

// GlobalParams is an autogenerated mock type for the GlobalParams type
type GlobalParams struct {
	mock.Mock
}

// ChainID provides a mock function with given fields:
func (_m *GlobalParams) ChainID() flow.ChainID {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for ChainID")
	}

	var r0 flow.ChainID
	if rf, ok := ret.Get(0).(func() flow.ChainID); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(flow.ChainID)
	}

	return r0
}

// ProtocolVersion provides a mock function with given fields:
func (_m *GlobalParams) ProtocolVersion() uint {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for ProtocolVersion")
	}

	var r0 uint
	if rf, ok := ret.Get(0).(func() uint); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(uint)
	}

	return r0
}

// SporkID provides a mock function with given fields:
func (_m *GlobalParams) SporkID() flow.Identifier {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for SporkID")
	}

	var r0 flow.Identifier
	if rf, ok := ret.Get(0).(func() flow.Identifier); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(flow.Identifier)
		}
	}

	return r0
}

// SporkRootBlockHeight provides a mock function with given fields:
func (_m *GlobalParams) SporkRootBlockHeight() uint64 {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for SporkRootBlockHeight")
	}

	var r0 uint64
	if rf, ok := ret.Get(0).(func() uint64); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(uint64)
	}

	return r0
}

// NewGlobalParams creates a new instance of GlobalParams. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewGlobalParams(t interface {
	mock.TestingT
	Cleanup(func())
}) *GlobalParams {
	mock := &GlobalParams{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
