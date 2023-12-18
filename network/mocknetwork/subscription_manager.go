// Code generated by mockery v2.21.4. DO NOT EDIT.

package mocknetwork

import (
	mock "github.com/stretchr/testify/mock"

	channels "github.com/onflow/flow-go/network/channels"

	network "github.com/onflow/flow-go/network"
)

// SubscriptionManager is an autogenerated mock type for the SubscriptionManager type
type SubscriptionManager struct {
	mock.Mock
}

// Channels provides a mock function with given fields:
func (_m *SubscriptionManager) Channels() channels.ChannelList {
	ret := _m.Called()

	var r0 channels.ChannelList
	if rf, ok := ret.Get(0).(func() channels.ChannelList); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(channels.ChannelList)
		}
	}

	return r0
}

// GetEngine provides a mock function with given fields: channel
func (_m *SubscriptionManager) GetEngine(channel channels.Channel) (network.MessageProcessor, error) {
	ret := _m.Called(channel)

	var r0 network.MessageProcessor
	var r1 error
	if rf, ok := ret.Get(0).(func(channels.Channel) (network.MessageProcessor, error)); ok {
		return rf(channel)
	}
	if rf, ok := ret.Get(0).(func(channels.Channel) network.MessageProcessor); ok {
		r0 = rf(channel)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(network.MessageProcessor)
		}
	}

	if rf, ok := ret.Get(1).(func(channels.Channel) error); ok {
		r1 = rf(channel)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Register provides a mock function with given fields: channel, engine
func (_m *SubscriptionManager) Register(channel channels.Channel, engine network.MessageProcessor) error {
	ret := _m.Called(channel, engine)

	var r0 error
	if rf, ok := ret.Get(0).(func(channels.Channel, network.MessageProcessor) error); ok {
		r0 = rf(channel, engine)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Unregister provides a mock function with given fields: channel
func (_m *SubscriptionManager) Unregister(channel channels.Channel) error {
	ret := _m.Called(channel)

	var r0 error
	if rf, ok := ret.Get(0).(func(channels.Channel) error); ok {
		r0 = rf(channel)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewSubscriptionManager interface {
	mock.TestingT
	Cleanup(func())
}

// NewSubscriptionManager creates a new instance of SubscriptionManager. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewSubscriptionManager(t mockConstructorTestingTNewSubscriptionManager) *SubscriptionManager {
	mock := &SubscriptionManager{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
