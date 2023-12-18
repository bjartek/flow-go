// Code generated by mockery v2.21.4. DO NOT EDIT.

package mock

import (
	mock "github.com/stretchr/testify/mock"

	flow "github.com/onflow/flow-go/model/flow"
)

// Requester is an autogenerated mock type for the Requester type
type Requester struct {
	mock.Mock
}

// EntityByID provides a mock function with given fields: entityID, selector
func (_m *Requester) EntityByID(entityID flow.Identifier, selector flow.IdentityFilter) {
	_m.Called(entityID, selector)
}

// Force provides a mock function with given fields:
func (_m *Requester) Force() {
	_m.Called()
}

// Query provides a mock function with given fields: key, selector
func (_m *Requester) Query(key flow.Identifier, selector flow.IdentityFilter) {
	_m.Called(key, selector)
}

type mockConstructorTestingTNewRequester interface {
	mock.TestingT
	Cleanup(func())
}

// NewRequester creates a new instance of Requester. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewRequester(t mockConstructorTestingTNewRequester) *Requester {
	mock := &Requester{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
