// Code generated by mockery v2.21.4. DO NOT EDIT.

package mock

import (
	flow "github.com/onflow/flow-go/model/flow"
	mock "github.com/stretchr/testify/mock"

	protocol "github.com/onflow/flow-go/state/protocol"
)

// DynamicProtocolState is an autogenerated mock type for the DynamicProtocolState type
type DynamicProtocolState struct {
	mock.Mock
}

// Clustering provides a mock function with given fields:
func (_m *DynamicProtocolState) Clustering() (flow.ClusterList, error) {
	ret := _m.Called()

	var r0 flow.ClusterList
	var r1 error
	if rf, ok := ret.Get(0).(func() (flow.ClusterList, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() flow.ClusterList); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(flow.ClusterList)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DKG provides a mock function with given fields:
func (_m *DynamicProtocolState) DKG() (protocol.DKG, error) {
	ret := _m.Called()

	var r0 protocol.DKG
	var r1 error
	if rf, ok := ret.Get(0).(func() (protocol.DKG, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() protocol.DKG); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(protocol.DKG)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Entry provides a mock function with given fields:
func (_m *DynamicProtocolState) Entry() *flow.RichProtocolStateEntry {
	ret := _m.Called()

	var r0 *flow.RichProtocolStateEntry
	if rf, ok := ret.Get(0).(func() *flow.RichProtocolStateEntry); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*flow.RichProtocolStateEntry)
		}
	}

	return r0
}

// Epoch provides a mock function with given fields:
func (_m *DynamicProtocolState) Epoch() uint64 {
	ret := _m.Called()

	var r0 uint64
	if rf, ok := ret.Get(0).(func() uint64); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(uint64)
	}

	return r0
}

// EpochCommit provides a mock function with given fields:
func (_m *DynamicProtocolState) EpochCommit() *flow.EpochCommit {
	ret := _m.Called()

	var r0 *flow.EpochCommit
	if rf, ok := ret.Get(0).(func() *flow.EpochCommit); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*flow.EpochCommit)
		}
	}

	return r0
}

// EpochPhase provides a mock function with given fields:
func (_m *DynamicProtocolState) EpochPhase() flow.EpochPhase {
	ret := _m.Called()

	var r0 flow.EpochPhase
	if rf, ok := ret.Get(0).(func() flow.EpochPhase); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(flow.EpochPhase)
	}

	return r0
}

// EpochSetup provides a mock function with given fields:
func (_m *DynamicProtocolState) EpochSetup() *flow.EpochSetup {
	ret := _m.Called()

	var r0 *flow.EpochSetup
	if rf, ok := ret.Get(0).(func() *flow.EpochSetup); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*flow.EpochSetup)
		}
	}

	return r0
}

// GlobalParams provides a mock function with given fields:
func (_m *DynamicProtocolState) GlobalParams() protocol.GlobalParams {
	ret := _m.Called()

	var r0 protocol.GlobalParams
	if rf, ok := ret.Get(0).(func() protocol.GlobalParams); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(protocol.GlobalParams)
		}
	}

	return r0
}

// Identities provides a mock function with given fields:
func (_m *DynamicProtocolState) Identities() flow.GenericIdentityList[flow.Identity] {
	ret := _m.Called()

	var r0 flow.GenericIdentityList[flow.Identity]
	if rf, ok := ret.Get(0).(func() flow.GenericIdentityList[flow.Identity]); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(flow.GenericIdentityList[flow.Identity])
		}
	}

	return r0
}

// InvalidEpochTransitionAttempted provides a mock function with given fields:
func (_m *DynamicProtocolState) InvalidEpochTransitionAttempted() bool {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// PreviousEpochExists provides a mock function with given fields:
func (_m *DynamicProtocolState) PreviousEpochExists() bool {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

type mockConstructorTestingTNewDynamicProtocolState interface {
	mock.TestingT
	Cleanup(func())
}

// NewDynamicProtocolState creates a new instance of DynamicProtocolState. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewDynamicProtocolState(t mockConstructorTestingTNewDynamicProtocolState) *DynamicProtocolState {
	mock := &DynamicProtocolState{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
