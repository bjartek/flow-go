// Code generated by mockery v2.13.0. DO NOT EDIT.

package mock

import (
	mock "github.com/stretchr/testify/mock"

	fvm "github.com/onflow/flow-go/fvm"

	programs "github.com/onflow/flow-go/fvm/programs"

	state "github.com/onflow/flow-go/fvm/state"
)

// VirtualMachine is an autogenerated mock type for the VirtualMachine type
type VirtualMachine struct {
	mock.Mock
}

// Run provides a mock function with given fields: _a0, _a1, _a2, _a3
func (_m *VirtualMachine) Run(_a0 fvm.Context, _a1 fvm.Procedure, _a2 state.View, _a3 *programs.Programs) error {
	ret := _m.Called(_a0, _a1, _a2, _a3)

	var r0 error
	if rf, ok := ret.Get(0).(func(fvm.Context, fvm.Procedure, state.View, *programs.Programs) error); ok {
		r0 = rf(_a0, _a1, _a2, _a3)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type NewVirtualMachineT interface {
	mock.TestingT
	Cleanup(func())
}

// NewVirtualMachine creates a new instance of VirtualMachine. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewVirtualMachine(t NewVirtualMachineT) *VirtualMachine {
	mock := &VirtualMachine{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
