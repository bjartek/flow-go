// Code generated by mockery v2.21.4. DO NOT EDIT.

package mockp2p

import (
	mock "github.com/stretchr/testify/mock"

	p2p "github.com/onflow/flow-go/network/p2p"

	peer "github.com/libp2p/go-libp2p/core/peer"
)

// PeerScoreExposer is an autogenerated mock type for the PeerScoreExposer type
type PeerScoreExposer struct {
	mock.Mock
}

// GetAppScore provides a mock function with given fields: peerID
func (_m *PeerScoreExposer) GetAppScore(peerID peer.ID) (float64, bool) {
	ret := _m.Called(peerID)

	var r0 float64
	var r1 bool
	if rf, ok := ret.Get(0).(func(peer.ID) (float64, bool)); ok {
		return rf(peerID)
	}
	if rf, ok := ret.Get(0).(func(peer.ID) float64); ok {
		r0 = rf(peerID)
	} else {
		r0 = ret.Get(0).(float64)
	}

	if rf, ok := ret.Get(1).(func(peer.ID) bool); ok {
		r1 = rf(peerID)
	} else {
		r1 = ret.Get(1).(bool)
	}

	return r0, r1
}

// GetBehaviourPenalty provides a mock function with given fields: peerID
func (_m *PeerScoreExposer) GetBehaviourPenalty(peerID peer.ID) (float64, bool) {
	ret := _m.Called(peerID)

	var r0 float64
	var r1 bool
	if rf, ok := ret.Get(0).(func(peer.ID) (float64, bool)); ok {
		return rf(peerID)
	}
	if rf, ok := ret.Get(0).(func(peer.ID) float64); ok {
		r0 = rf(peerID)
	} else {
		r0 = ret.Get(0).(float64)
	}

	if rf, ok := ret.Get(1).(func(peer.ID) bool); ok {
		r1 = rf(peerID)
	} else {
		r1 = ret.Get(1).(bool)
	}

	return r0, r1
}

// GetIPColocationFactor provides a mock function with given fields: peerID
func (_m *PeerScoreExposer) GetIPColocationFactor(peerID peer.ID) (float64, bool) {
	ret := _m.Called(peerID)

	var r0 float64
	var r1 bool
	if rf, ok := ret.Get(0).(func(peer.ID) (float64, bool)); ok {
		return rf(peerID)
	}
	if rf, ok := ret.Get(0).(func(peer.ID) float64); ok {
		r0 = rf(peerID)
	} else {
		r0 = ret.Get(0).(float64)
	}

	if rf, ok := ret.Get(1).(func(peer.ID) bool); ok {
		r1 = rf(peerID)
	} else {
		r1 = ret.Get(1).(bool)
	}

	return r0, r1
}

// GetScore provides a mock function with given fields: peerID
func (_m *PeerScoreExposer) GetScore(peerID peer.ID) (float64, bool) {
	ret := _m.Called(peerID)

	var r0 float64
	var r1 bool
	if rf, ok := ret.Get(0).(func(peer.ID) (float64, bool)); ok {
		return rf(peerID)
	}
	if rf, ok := ret.Get(0).(func(peer.ID) float64); ok {
		r0 = rf(peerID)
	} else {
		r0 = ret.Get(0).(float64)
	}

	if rf, ok := ret.Get(1).(func(peer.ID) bool); ok {
		r1 = rf(peerID)
	} else {
		r1 = ret.Get(1).(bool)
	}

	return r0, r1
}

// GetTopicScores provides a mock function with given fields: peerID
func (_m *PeerScoreExposer) GetTopicScores(peerID peer.ID) (map[string]p2p.TopicScoreSnapshot, bool) {
	ret := _m.Called(peerID)

	var r0 map[string]p2p.TopicScoreSnapshot
	var r1 bool
	if rf, ok := ret.Get(0).(func(peer.ID) (map[string]p2p.TopicScoreSnapshot, bool)); ok {
		return rf(peerID)
	}
	if rf, ok := ret.Get(0).(func(peer.ID) map[string]p2p.TopicScoreSnapshot); ok {
		r0 = rf(peerID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[string]p2p.TopicScoreSnapshot)
		}
	}

	if rf, ok := ret.Get(1).(func(peer.ID) bool); ok {
		r1 = rf(peerID)
	} else {
		r1 = ret.Get(1).(bool)
	}

	return r0, r1
}

type mockConstructorTestingTNewPeerScoreExposer interface {
	mock.TestingT
	Cleanup(func())
}

// NewPeerScoreExposer creates a new instance of PeerScoreExposer. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewPeerScoreExposer(t mockConstructorTestingTNewPeerScoreExposer) *PeerScoreExposer {
	mock := &PeerScoreExposer{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
