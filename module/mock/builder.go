// Code generated by mockery v2.21.4. DO NOT EDIT.

package mock

import (
	mock "github.com/stretchr/testify/mock"

	flow "github.com/onflow/flow-go/model/flow"
)

// Builder is an autogenerated mock type for the Builder type
type Builder struct {
	mock.Mock
}

// BuildOn provides a mock function with given fields: parentID, setter
func (_m *Builder) BuildOn(parentID flow.Identifier, setter func(*flow.Header) error) (*flow.Header, error) {
	ret := _m.Called(parentID, setter)

	var r0 *flow.Header
	var r1 error
	if rf, ok := ret.Get(0).(func(flow.Identifier, func(*flow.Header) error) (*flow.Header, error)); ok {
		return rf(parentID, setter)
	}
	if rf, ok := ret.Get(0).(func(flow.Identifier, func(*flow.Header) error) *flow.Header); ok {
		r0 = rf(parentID, setter)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*flow.Header)
		}
	}

	if rf, ok := ret.Get(1).(func(flow.Identifier, func(*flow.Header) error) error); ok {
		r1 = rf(parentID, setter)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewBuilder interface {
	mock.TestingT
	Cleanup(func())
}

// NewBuilder creates a new instance of Builder. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewBuilder(t mockConstructorTestingTNewBuilder) *Builder {
	mock := &Builder{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
