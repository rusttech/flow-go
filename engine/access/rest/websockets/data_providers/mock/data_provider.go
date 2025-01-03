// Code generated by mockery v2.43.2. DO NOT EDIT.

package mock

import (
	uuid "github.com/google/uuid"
	mock "github.com/stretchr/testify/mock"
)

// DataProvider is an autogenerated mock type for the DataProvider type
type DataProvider struct {
	mock.Mock
}

// Close provides a mock function with given fields:
func (_m *DataProvider) Close() {
	_m.Called()
}

// ID provides a mock function with given fields:
func (_m *DataProvider) ID() uuid.UUID {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for ID")
	}

	var r0 uuid.UUID
	if rf, ok := ret.Get(0).(func() uuid.UUID); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(uuid.UUID)
		}
	}

	return r0
}

// Run provides a mock function with given fields:
func (_m *DataProvider) Run() error {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Run")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Topic provides a mock function with given fields:
func (_m *DataProvider) Topic() string {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Topic")
	}

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// NewDataProvider creates a new instance of DataProvider. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewDataProvider(t interface {
	mock.TestingT
	Cleanup(func())
}) *DataProvider {
	mock := &DataProvider{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
