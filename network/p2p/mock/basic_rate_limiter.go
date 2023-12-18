// Code generated by mockery v2.21.4. DO NOT EDIT.

package mockp2p

import (
	mock "github.com/stretchr/testify/mock"

	irrecoverable "github.com/onflow/flow-go/module/irrecoverable"

	peer "github.com/libp2p/go-libp2p/core/peer"
)

// BasicRateLimiter is an autogenerated mock type for the BasicRateLimiter type
type BasicRateLimiter struct {
	mock.Mock
}

// Allow provides a mock function with given fields: peerID, msgSize
func (_m *BasicRateLimiter) Allow(peerID peer.ID, msgSize int) bool {
	ret := _m.Called(peerID, msgSize)

	var r0 bool
	if rf, ok := ret.Get(0).(func(peer.ID, int) bool); ok {
		r0 = rf(peerID, msgSize)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// Done provides a mock function with given fields:
func (_m *BasicRateLimiter) Done() <-chan struct{} {
	ret := _m.Called()

	var r0 <-chan struct{}
	if rf, ok := ret.Get(0).(func() <-chan struct{}); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(<-chan struct{})
		}
	}

	return r0
}

// Ready provides a mock function with given fields:
func (_m *BasicRateLimiter) Ready() <-chan struct{} {
	ret := _m.Called()

	var r0 <-chan struct{}
	if rf, ok := ret.Get(0).(func() <-chan struct{}); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(<-chan struct{})
		}
	}

	return r0
}

// Start provides a mock function with given fields: _a0
func (_m *BasicRateLimiter) Start(_a0 irrecoverable.SignalerContext) {
	_m.Called(_a0)
}

type mockConstructorTestingTNewBasicRateLimiter interface {
	mock.TestingT
	Cleanup(func())
}

// NewBasicRateLimiter creates a new instance of BasicRateLimiter. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewBasicRateLimiter(t mockConstructorTestingTNewBasicRateLimiter) *BasicRateLimiter {
	mock := &BasicRateLimiter{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
