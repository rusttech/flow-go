// Code generated by mockery v2.21.4. DO NOT EDIT.

package mocks

import (
	model "github.com/onflow/flow-go/consensus/hotstuff/model"
	"github.com/onflow/flow-go/model/flow"
	mock "github.com/stretchr/testify/mock"
)

// ProposalViolationConsumer is an autogenerated mock type for the ProposalViolationConsumer type
type ProposalViolationConsumer struct {
	mock.Mock
}

// OnDoubleProposeDetected provides a mock function with given fields: _a0, _a1
func (_m *ProposalViolationConsumer) OnDoubleProposeDetected(_a0 *model.Block, _a1 *model.Block) {
	_m.Called(_a0, _a1)
}

// OnInvalidBlockDetected provides a mock function with given fields: err
func (_m *ProposalViolationConsumer) OnInvalidBlockDetected(err flow.Slashable[model.InvalidProposalError]) {
	_m.Called(err)
}

type mockConstructorTestingTNewProposalViolationConsumer interface {
	mock.TestingT
	Cleanup(func())
}

// NewProposalViolationConsumer creates a new instance of ProposalViolationConsumer. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewProposalViolationConsumer(t mockConstructorTestingTNewProposalViolationConsumer) *ProposalViolationConsumer {
	mock := &ProposalViolationConsumer{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
