// Code generated by mockery v2.13.1. DO NOT EDIT.

package mock

import (
	common "github.com/onflow/cadence/runtime/common"

	meter "github.com/onflow/flow-go/fvm/meter"

	mock "github.com/stretchr/testify/mock"
)

// Meter is an autogenerated mock type for the Meter type
type Meter struct {
	mock.Mock
}

// ComputationIntensities provides a mock function with given fields:
func (_m *Meter) ComputationIntensities() meter.MeteredComputationIntensities {
	ret := _m.Called()

	var r0 meter.MeteredComputationIntensities
	if rf, ok := ret.Get(0).(func() meter.MeteredComputationIntensities); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(meter.MeteredComputationIntensities)
		}
	}

	return r0
}

// ComputationUsed provides a mock function with given fields:
func (_m *Meter) ComputationUsed() uint64 {
	ret := _m.Called()

	var r0 uint64
	if rf, ok := ret.Get(0).(func() uint64); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(uint64)
	}

	return r0
}

// MemoryEstimate provides a mock function with given fields:
func (_m *Meter) MemoryEstimate() uint64 {
	ret := _m.Called()

	var r0 uint64
	if rf, ok := ret.Get(0).(func() uint64); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(uint64)
	}

	return r0
}

// MeterComputation provides a mock function with given fields: _a0, _a1
func (_m *Meter) MeterComputation(_a0 common.ComputationKind, _a1 uint) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(common.ComputationKind, uint) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MeterEmittedEvent provides a mock function with given fields: byteSize
func (_m *Meter) MeterEmittedEvent(byteSize uint64) error {
	ret := _m.Called(byteSize)

	var r0 error
	if rf, ok := ret.Get(0).(func(uint64) error); ok {
		r0 = rf(byteSize)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MeterMemory provides a mock function with given fields: usage
func (_m *Meter) MeterMemory(usage common.MemoryUsage) error {
	ret := _m.Called(usage)

	var r0 error
	if rf, ok := ret.Get(0).(func(common.MemoryUsage) error); ok {
		r0 = rf(usage)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// TotalEmittedEventBytes provides a mock function with given fields:
func (_m *Meter) TotalEmittedEventBytes() uint64 {
	ret := _m.Called()

	var r0 uint64
	if rf, ok := ret.Get(0).(func() uint64); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(uint64)
	}

	return r0
}

type mockConstructorTestingTNewMeter interface {
	mock.TestingT
	Cleanup(func())
}

// NewMeter creates a new instance of Meter. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMeter(t mockConstructorTestingTNewMeter) *Meter {
	mock := &Meter{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
