// Code generated by mockery v2.21.4. DO NOT EDIT.

package mocknetwork

import (
	mock "github.com/stretchr/testify/mock"

	network "github.com/onflow/flow-go/network"
)

// DisallowListNotificationConsumer is an autogenerated mock type for the DisallowListNotificationConsumer type
type DisallowListNotificationConsumer struct {
	mock.Mock
}

// OnAllowListNotification provides a mock function with given fields: _a0
func (_m *DisallowListNotificationConsumer) OnAllowListNotification(_a0 *network.AllowListingUpdate) {
	_m.Called(_a0)
}

// OnDisallowListNotification provides a mock function with given fields: _a0
func (_m *DisallowListNotificationConsumer) OnDisallowListNotification(_a0 *network.DisallowListingUpdate) {
	_m.Called(_a0)
}

type mockConstructorTestingTNewDisallowListNotificationConsumer interface {
	mock.TestingT
	Cleanup(func())
}

// NewDisallowListNotificationConsumer creates a new instance of DisallowListNotificationConsumer. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewDisallowListNotificationConsumer(t mockConstructorTestingTNewDisallowListNotificationConsumer) *DisallowListNotificationConsumer {
	mock := &DisallowListNotificationConsumer{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
