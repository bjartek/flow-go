// Code generated by mockery v2.21.4. DO NOT EDIT.

package mockp2p

import (
	mock "github.com/stretchr/testify/mock"

	flow "github.com/onflow/flow-go/model/flow"

	peer "github.com/libp2p/go-libp2p/core/peer"
)

// IDTranslator is an autogenerated mock type for the IDTranslator type
type IDTranslator struct {
	mock.Mock
}

// GetFlowID provides a mock function with given fields: _a0
func (_m *IDTranslator) GetFlowID(_a0 peer.ID) (flow.Identifier, error) {
	ret := _m.Called(_a0)

	var r0 flow.Identifier
	var r1 error
	if rf, ok := ret.Get(0).(func(peer.ID) (flow.Identifier, error)); ok {
		return rf(_a0)
	}
	if rf, ok := ret.Get(0).(func(peer.ID) flow.Identifier); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(flow.Identifier)
		}
	}

	if rf, ok := ret.Get(1).(func(peer.ID) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetPeerID provides a mock function with given fields: _a0
func (_m *IDTranslator) GetPeerID(_a0 flow.Identifier) (peer.ID, error) {
	ret := _m.Called(_a0)

	var r0 peer.ID
	var r1 error
	if rf, ok := ret.Get(0).(func(flow.Identifier) (peer.ID, error)); ok {
		return rf(_a0)
	}
	if rf, ok := ret.Get(0).(func(flow.Identifier) peer.ID); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(peer.ID)
	}

	if rf, ok := ret.Get(1).(func(flow.Identifier) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewIDTranslator interface {
	mock.TestingT
	Cleanup(func())
}

// NewIDTranslator creates a new instance of IDTranslator. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewIDTranslator(t mockConstructorTestingTNewIDTranslator) *IDTranslator {
	mock := &IDTranslator{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
