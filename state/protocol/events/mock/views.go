// Code generated by mockery v2.21.4. DO NOT EDIT.

package mock

import (
	mock "github.com/stretchr/testify/mock"

	events "github.com/onflow/flow-go/state/protocol/events"
)

// Views is an autogenerated mock type for the Views type
type Views struct {
	mock.Mock
}

// OnView provides a mock function with given fields: view, callback
func (_m *Views) OnView(view uint64, callback events.OnViewCallback) {
	_m.Called(view, callback)
}

type mockConstructorTestingTNewViews interface {
	mock.TestingT
	Cleanup(func())
}

// NewViews creates a new instance of Views. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewViews(t mockConstructorTestingTNewViews) *Views {
	mock := &Views{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
