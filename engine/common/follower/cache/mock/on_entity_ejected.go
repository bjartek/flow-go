// Code generated by mockery v2.21.4. DO NOT EDIT.

package mock

import (
	mock "github.com/stretchr/testify/mock"

	flow "github.com/onflow/flow-go/model/flow"
)

// OnEntityEjected is an autogenerated mock type for the OnEntityEjected type
type OnEntityEjected struct {
	mock.Mock
}

// Execute provides a mock function with given fields: ejectedEntity
func (_m *OnEntityEjected) Execute(ejectedEntity flow.Entity) {
	_m.Called(ejectedEntity)
}

type mockConstructorTestingTNewOnEntityEjected interface {
	mock.TestingT
	Cleanup(func())
}

// NewOnEntityEjected creates a new instance of OnEntityEjected. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewOnEntityEjected(t mockConstructorTestingTNewOnEntityEjected) *OnEntityEjected {
	mock := &OnEntityEjected{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
