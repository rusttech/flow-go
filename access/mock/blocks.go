// Code generated by mockery v2.43.2. DO NOT EDIT.

package mock

import (
	flow "github.com/onflow/flow-go/model/flow"
	mock "github.com/stretchr/testify/mock"
)

// Blocks is an autogenerated mock type for the Blocks type
type Blocks struct {
	mock.Mock
}

// FinalizedHeader provides a mock function with given fields:
func (_m *Blocks) FinalizedHeader() (*flow.Header, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for FinalizedHeader")
	}

	var r0 *flow.Header
	var r1 error
	if rf, ok := ret.Get(0).(func() (*flow.Header, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() *flow.Header); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*flow.Header)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// HeaderByID provides a mock function with given fields: id
func (_m *Blocks) HeaderByID(id flow.Identifier) (*flow.Header, error) {
	ret := _m.Called(id)

	if len(ret) == 0 {
		panic("no return value specified for HeaderByID")
	}

	var r0 *flow.Header
	var r1 error
	if rf, ok := ret.Get(0).(func(flow.Identifier) (*flow.Header, error)); ok {
		return rf(id)
	}
	if rf, ok := ret.Get(0).(func(flow.Identifier) *flow.Header); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*flow.Header)
		}
	}

	if rf, ok := ret.Get(1).(func(flow.Identifier) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SealedHeader provides a mock function with given fields:
func (_m *Blocks) SealedHeader() (*flow.Header, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for SealedHeader")
	}

	var r0 *flow.Header
	var r1 error
	if rf, ok := ret.Get(0).(func() (*flow.Header, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() *flow.Header); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*flow.Header)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewBlocks creates a new instance of Blocks. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewBlocks(t interface {
	mock.TestingT
	Cleanup(func())
}) *Blocks {
	mock := &Blocks{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}