// Code generated by mockery v2.13.1. DO NOT EDIT.

package mock

import (
	state "github.com/onflow/flow-go/fvm/state"
	mock "github.com/stretchr/testify/mock"
)

// StateOption is an autogenerated mock type for the StateOption type
type StateOption struct {
	mock.Mock
}

// Execute provides a mock function with given fields: st
func (_m *StateOption) Execute(st *state.State) *state.State {
	ret := _m.Called(st)

	var r0 *state.State
	if rf, ok := ret.Get(0).(func(*state.State) *state.State); ok {
		r0 = rf(st)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*state.State)
		}
	}

	return r0
}

type mockConstructorTestingTNewStateOption interface {
	mock.TestingT
	Cleanup(func())
}

// NewStateOption creates a new instance of StateOption. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewStateOption(t mockConstructorTestingTNewStateOption) *StateOption {
	mock := &StateOption{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
