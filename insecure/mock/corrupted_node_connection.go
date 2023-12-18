// Code generated by mockery v2.21.4. DO NOT EDIT.

package mockinsecure

import (
	mock "github.com/stretchr/testify/mock"

	insecure "github.com/onflow/flow-go/insecure"
)

// CorruptedNodeConnection is an autogenerated mock type for the CorruptedNodeConnection type
type CorruptedNodeConnection struct {
	mock.Mock
}

// CloseConnection provides a mock function with given fields:
func (_m *CorruptedNodeConnection) CloseConnection() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SendMessage provides a mock function with given fields: _a0
func (_m *CorruptedNodeConnection) SendMessage(_a0 *insecure.Message) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(*insecure.Message) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewCorruptedNodeConnection interface {
	mock.TestingT
	Cleanup(func())
}

// NewCorruptedNodeConnection creates a new instance of CorruptedNodeConnection. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewCorruptedNodeConnection(t mockConstructorTestingTNewCorruptedNodeConnection) *CorruptedNodeConnection {
	mock := &CorruptedNodeConnection{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
