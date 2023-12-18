// Code generated by mockery v2.21.4. DO NOT EDIT.

package mocknetwork

import (
	mock "github.com/stretchr/testify/mock"

	flow "github.com/onflow/flow-go/model/flow"

	network "github.com/onflow/flow-go/network"
)

// MisbehaviorReport is an autogenerated mock type for the MisbehaviorReport type
type MisbehaviorReport struct {
	mock.Mock
}

// OriginId provides a mock function with given fields:
func (_m *MisbehaviorReport) OriginId() flow.Identifier {
	ret := _m.Called()

	var r0 flow.Identifier
	if rf, ok := ret.Get(0).(func() flow.Identifier); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(flow.Identifier)
		}
	}

	return r0
}

// Penalty provides a mock function with given fields:
func (_m *MisbehaviorReport) Penalty() float64 {
	ret := _m.Called()

	var r0 float64
	if rf, ok := ret.Get(0).(func() float64); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(float64)
	}

	return r0
}

// Reason provides a mock function with given fields:
func (_m *MisbehaviorReport) Reason() network.Misbehavior {
	ret := _m.Called()

	var r0 network.Misbehavior
	if rf, ok := ret.Get(0).(func() network.Misbehavior); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(network.Misbehavior)
	}

	return r0
}

type mockConstructorTestingTNewMisbehaviorReport interface {
	mock.TestingT
	Cleanup(func())
}

// NewMisbehaviorReport creates a new instance of MisbehaviorReport. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMisbehaviorReport(t mockConstructorTestingTNewMisbehaviorReport) *MisbehaviorReport {
	mock := &MisbehaviorReport{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
