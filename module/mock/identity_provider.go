// Code generated by mockery v2.21.4. DO NOT EDIT.

package mock

import (
	mock "github.com/stretchr/testify/mock"

	flow "github.com/onflow/flow-go/model/flow"

	peer "github.com/libp2p/go-libp2p/core/peer"
)

// IdentityProvider is an autogenerated mock type for the IdentityProvider type
type IdentityProvider struct {
	mock.Mock
}

// ByNodeID provides a mock function with given fields: _a0
func (_m *IdentityProvider) ByNodeID(_a0 flow.Identifier) (*flow.Identity, bool) {
	ret := _m.Called(_a0)

	var r0 *flow.Identity
	var r1 bool
	if rf, ok := ret.Get(0).(func(flow.Identifier) (*flow.Identity, bool)); ok {
		return rf(_a0)
	}
	if rf, ok := ret.Get(0).(func(flow.Identifier) *flow.Identity); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*flow.Identity)
		}
	}

	if rf, ok := ret.Get(1).(func(flow.Identifier) bool); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Get(1).(bool)
	}

	return r0, r1
}

// ByPeerID provides a mock function with given fields: _a0
func (_m *IdentityProvider) ByPeerID(_a0 peer.ID) (*flow.Identity, bool) {
	ret := _m.Called(_a0)

	var r0 *flow.Identity
	var r1 bool
	if rf, ok := ret.Get(0).(func(peer.ID) (*flow.Identity, bool)); ok {
		return rf(_a0)
	}
	if rf, ok := ret.Get(0).(func(peer.ID) *flow.Identity); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*flow.Identity)
		}
	}

	if rf, ok := ret.Get(1).(func(peer.ID) bool); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Get(1).(bool)
	}

	return r0, r1
}

// Identities provides a mock function with given fields: _a0
func (_m *IdentityProvider) Identities(_a0 flow.IdentityFilter) flow.IdentityList {
	ret := _m.Called(_a0)

	var r0 flow.IdentityList
	if rf, ok := ret.Get(0).(func(flow.IdentityFilter) flow.IdentityList); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(flow.IdentityList)
		}
	}

	return r0
}

type mockConstructorTestingTNewIdentityProvider interface {
	mock.TestingT
	Cleanup(func())
}

// NewIdentityProvider creates a new instance of IdentityProvider. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewIdentityProvider(t mockConstructorTestingTNewIdentityProvider) *IdentityProvider {
	mock := &IdentityProvider{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
