// Code generated by mockery v2.21.4. DO NOT EDIT.

package mock

import (
	mock "github.com/stretchr/testify/mock"

	flow "github.com/onflow/flow-go/model/flow"
)

// Guarantees is an autogenerated mock type for the Guarantees type
type Guarantees struct {
	mock.Mock
}

// ByCollectionID provides a mock function with given fields: collID
func (_m *Guarantees) ByCollectionID(collID flow.Identifier) (*flow.CollectionGuarantee, error) {
	ret := _m.Called(collID)

	var r0 *flow.CollectionGuarantee
	var r1 error
	if rf, ok := ret.Get(0).(func(flow.Identifier) (*flow.CollectionGuarantee, error)); ok {
		return rf(collID)
	}
	if rf, ok := ret.Get(0).(func(flow.Identifier) *flow.CollectionGuarantee); ok {
		r0 = rf(collID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*flow.CollectionGuarantee)
		}
	}

	if rf, ok := ret.Get(1).(func(flow.Identifier) error); ok {
		r1 = rf(collID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Store provides a mock function with given fields: guarantee
func (_m *Guarantees) Store(guarantee *flow.CollectionGuarantee) error {
	ret := _m.Called(guarantee)

	var r0 error
	if rf, ok := ret.Get(0).(func(*flow.CollectionGuarantee) error); ok {
		r0 = rf(guarantee)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewGuarantees interface {
	mock.TestingT
	Cleanup(func())
}

// NewGuarantees creates a new instance of Guarantees. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewGuarantees(t mockConstructorTestingTNewGuarantees) *Guarantees {
	mock := &Guarantees{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
