// Code generated by mockery v2.21.4. DO NOT EDIT.

package mock

import (
	mock "github.com/stretchr/testify/mock"

	crypto "github.com/onflow/flow-go/crypto"
)

// SafeBeaconKeys is an autogenerated mock type for the SafeBeaconKeys type
type SafeBeaconKeys struct {
	mock.Mock
}

// RetrieveMyBeaconPrivateKey provides a mock function with given fields: epochCounter
func (_m *SafeBeaconKeys) RetrieveMyBeaconPrivateKey(epochCounter uint64) (crypto.PrivateKey, bool, error) {
	ret := _m.Called(epochCounter)

	var r0 crypto.PrivateKey
	var r1 bool
	var r2 error
	if rf, ok := ret.Get(0).(func(uint64) (crypto.PrivateKey, bool, error)); ok {
		return rf(epochCounter)
	}
	if rf, ok := ret.Get(0).(func(uint64) crypto.PrivateKey); ok {
		r0 = rf(epochCounter)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(crypto.PrivateKey)
		}
	}

	if rf, ok := ret.Get(1).(func(uint64) bool); ok {
		r1 = rf(epochCounter)
	} else {
		r1 = ret.Get(1).(bool)
	}

	if rf, ok := ret.Get(2).(func(uint64) error); ok {
		r2 = rf(epochCounter)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

type mockConstructorTestingTNewSafeBeaconKeys interface {
	mock.TestingT
	Cleanup(func())
}

// NewSafeBeaconKeys creates a new instance of SafeBeaconKeys. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewSafeBeaconKeys(t mockConstructorTestingTNewSafeBeaconKeys) *SafeBeaconKeys {
	mock := &SafeBeaconKeys{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
