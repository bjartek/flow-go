// Code generated by mockery v2.21.4. DO NOT EDIT.

package mock

import (
	mock "github.com/stretchr/testify/mock"

	flow "github.com/onflow/flow-go/model/flow"

	transaction "github.com/onflow/flow-go/storage/badger/transaction"
)

// QuorumCertificates is an autogenerated mock type for the QuorumCertificates type
type QuorumCertificates struct {
	mock.Mock
}

// ByBlockID provides a mock function with given fields: blockID
func (_m *QuorumCertificates) ByBlockID(blockID flow.Identifier) (*flow.QuorumCertificate, error) {
	ret := _m.Called(blockID)

	var r0 *flow.QuorumCertificate
	var r1 error
	if rf, ok := ret.Get(0).(func(flow.Identifier) (*flow.QuorumCertificate, error)); ok {
		return rf(blockID)
	}
	if rf, ok := ret.Get(0).(func(flow.Identifier) *flow.QuorumCertificate); ok {
		r0 = rf(blockID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*flow.QuorumCertificate)
		}
	}

	if rf, ok := ret.Get(1).(func(flow.Identifier) error); ok {
		r1 = rf(blockID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// StoreTx provides a mock function with given fields: qc
func (_m *QuorumCertificates) StoreTx(qc *flow.QuorumCertificate) func(*transaction.Tx) error {
	ret := _m.Called(qc)

	var r0 func(*transaction.Tx) error
	if rf, ok := ret.Get(0).(func(*flow.QuorumCertificate) func(*transaction.Tx) error); ok {
		r0 = rf(qc)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(func(*transaction.Tx) error)
		}
	}

	return r0
}

type mockConstructorTestingTNewQuorumCertificates interface {
	mock.TestingT
	Cleanup(func())
}

// NewQuorumCertificates creates a new instance of QuorumCertificates. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewQuorumCertificates(t mockConstructorTestingTNewQuorumCertificates) *QuorumCertificates {
	mock := &QuorumCertificates{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
