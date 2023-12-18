// Code generated by mockery v2.21.4. DO NOT EDIT.

package mockp2p

import (
	mock "github.com/stretchr/testify/mock"

	p2p "github.com/onflow/flow-go/network/p2p"

	peer "github.com/libp2p/go-libp2p/core/peer"
)

// GossipSubSpamRecordCache is an autogenerated mock type for the GossipSubSpamRecordCache type
type GossipSubSpamRecordCache struct {
	mock.Mock
}

// Add provides a mock function with given fields: peerId, record
func (_m *GossipSubSpamRecordCache) Add(peerId peer.ID, record p2p.GossipSubSpamRecord) bool {
	ret := _m.Called(peerId, record)

	var r0 bool
	if rf, ok := ret.Get(0).(func(peer.ID, p2p.GossipSubSpamRecord) bool); ok {
		r0 = rf(peerId, record)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// Get provides a mock function with given fields: peerID
func (_m *GossipSubSpamRecordCache) Get(peerID peer.ID) (*p2p.GossipSubSpamRecord, error, bool) {
	ret := _m.Called(peerID)

	var r0 *p2p.GossipSubSpamRecord
	var r1 error
	var r2 bool
	if rf, ok := ret.Get(0).(func(peer.ID) (*p2p.GossipSubSpamRecord, error, bool)); ok {
		return rf(peerID)
	}
	if rf, ok := ret.Get(0).(func(peer.ID) *p2p.GossipSubSpamRecord); ok {
		r0 = rf(peerID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*p2p.GossipSubSpamRecord)
		}
	}

	if rf, ok := ret.Get(1).(func(peer.ID) error); ok {
		r1 = rf(peerID)
	} else {
		r1 = ret.Error(1)
	}

	if rf, ok := ret.Get(2).(func(peer.ID) bool); ok {
		r2 = rf(peerID)
	} else {
		r2 = ret.Get(2).(bool)
	}

	return r0, r1, r2
}

// Has provides a mock function with given fields: peerID
func (_m *GossipSubSpamRecordCache) Has(peerID peer.ID) bool {
	ret := _m.Called(peerID)

	var r0 bool
	if rf, ok := ret.Get(0).(func(peer.ID) bool); ok {
		r0 = rf(peerID)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// Update provides a mock function with given fields: peerID, updateFunc
func (_m *GossipSubSpamRecordCache) Update(peerID peer.ID, updateFunc p2p.UpdateFunction) (*p2p.GossipSubSpamRecord, error) {
	ret := _m.Called(peerID, updateFunc)

	var r0 *p2p.GossipSubSpamRecord
	var r1 error
	if rf, ok := ret.Get(0).(func(peer.ID, p2p.UpdateFunction) (*p2p.GossipSubSpamRecord, error)); ok {
		return rf(peerID, updateFunc)
	}
	if rf, ok := ret.Get(0).(func(peer.ID, p2p.UpdateFunction) *p2p.GossipSubSpamRecord); ok {
		r0 = rf(peerID, updateFunc)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*p2p.GossipSubSpamRecord)
		}
	}

	if rf, ok := ret.Get(1).(func(peer.ID, p2p.UpdateFunction) error); ok {
		r1 = rf(peerID, updateFunc)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewGossipSubSpamRecordCache interface {
	mock.TestingT
	Cleanup(func())
}

// NewGossipSubSpamRecordCache creates a new instance of GossipSubSpamRecordCache. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewGossipSubSpamRecordCache(t mockConstructorTestingTNewGossipSubSpamRecordCache) *GossipSubSpamRecordCache {
	mock := &GossipSubSpamRecordCache{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
