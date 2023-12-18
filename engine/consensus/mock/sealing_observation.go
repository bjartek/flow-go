// Code generated by mockery v2.21.4. DO NOT EDIT.

package mock

import (
	mock "github.com/stretchr/testify/mock"

	flow "github.com/onflow/flow-go/model/flow"
)

// SealingObservation is an autogenerated mock type for the SealingObservation type
type SealingObservation struct {
	mock.Mock
}

// ApprovalsMissing provides a mock function with given fields: ir, chunksWithMissingApprovals
func (_m *SealingObservation) ApprovalsMissing(ir *flow.IncorporatedResult, chunksWithMissingApprovals map[uint64]flow.IdentifierList) {
	_m.Called(ir, chunksWithMissingApprovals)
}

// ApprovalsRequested provides a mock function with given fields: ir, requestCount
func (_m *SealingObservation) ApprovalsRequested(ir *flow.IncorporatedResult, requestCount uint) {
	_m.Called(ir, requestCount)
}

// Complete provides a mock function with given fields:
func (_m *SealingObservation) Complete() {
	_m.Called()
}

// QualifiesForEmergencySealing provides a mock function with given fields: ir, emergencySealable
func (_m *SealingObservation) QualifiesForEmergencySealing(ir *flow.IncorporatedResult, emergencySealable bool) {
	_m.Called(ir, emergencySealable)
}

type mockConstructorTestingTNewSealingObservation interface {
	mock.TestingT
	Cleanup(func())
}

// NewSealingObservation creates a new instance of SealingObservation. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewSealingObservation(t mockConstructorTestingTNewSealingObservation) *SealingObservation {
	mock := &SealingObservation{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
