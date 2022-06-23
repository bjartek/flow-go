// Code generated by mockery v2.13.0. DO NOT EDIT.

package mock

import (
	mock "github.com/stretchr/testify/mock"

	flow "github.com/onflow/flow-go/model/flow"

	module "github.com/onflow/flow-go/module"
)

// DKGControllerFactory is an autogenerated mock type for the DKGControllerFactory type
type DKGControllerFactory struct {
	mock.Mock
}

// Create provides a mock function with given fields: dkgInstanceID, participants, seed
func (_m *DKGControllerFactory) Create(dkgInstanceID string, participants flow.IdentityList, seed []byte) (module.DKGController, error) {
	ret := _m.Called(dkgInstanceID, participants, seed)

	var r0 module.DKGController
	if rf, ok := ret.Get(0).(func(string, flow.IdentityList, []byte) module.DKGController); ok {
		r0 = rf(dkgInstanceID, participants, seed)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(module.DKGController)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, flow.IdentityList, []byte) error); ok {
		r1 = rf(dkgInstanceID, participants, seed)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type NewDKGControllerFactoryT interface {
	mock.TestingT
	Cleanup(func())
}

// NewDKGControllerFactory creates a new instance of DKGControllerFactory. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewDKGControllerFactory(t NewDKGControllerFactoryT) *DKGControllerFactory {
	mock := &DKGControllerFactory{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
