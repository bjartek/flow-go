// Code generated by mockery v2.13.1. DO NOT EDIT.

package mock

import (
	context "context"

	delta "github.com/onflow/flow-go/engine/execution/state/delta"
	entity "github.com/onflow/flow-go/module/mempool/entity"

	execution "github.com/onflow/flow-go/engine/execution"

	flow "github.com/onflow/flow-go/model/flow"

	mock "github.com/stretchr/testify/mock"
)

// ComputationManager is an autogenerated mock type for the ComputationManager type
type ComputationManager struct {
	mock.Mock
}

// ComputeBlock provides a mock function with given fields: ctx, parentBlockExecutionResultID, block, snapshot
func (_m *ComputationManager) ComputeBlock(ctx context.Context, parentBlockExecutionResultID flow.Identifier, block *entity.ExecutableBlock, snapshot delta.StorageSnapshot) (*execution.ComputationResult, error) {
	ret := _m.Called(ctx, parentBlockExecutionResultID, block, snapshot)

	var r0 *execution.ComputationResult
	if rf, ok := ret.Get(0).(func(context.Context, flow.Identifier, *entity.ExecutableBlock, delta.StorageSnapshot) *execution.ComputationResult); ok {
		r0 = rf(ctx, parentBlockExecutionResultID, block, snapshot)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*execution.ComputationResult)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, flow.Identifier, *entity.ExecutableBlock, delta.StorageSnapshot) error); ok {
		r1 = rf(ctx, parentBlockExecutionResultID, block, snapshot)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ExecuteScript provides a mock function with given fields: ctx, script, arguments, blockHeader, snapshot
func (_m *ComputationManager) ExecuteScript(ctx context.Context, script []byte, arguments [][]byte, blockHeader *flow.Header, snapshot delta.StorageSnapshot) ([]byte, error) {
	ret := _m.Called(ctx, script, arguments, blockHeader, snapshot)

	var r0 []byte
	if rf, ok := ret.Get(0).(func(context.Context, []byte, [][]byte, *flow.Header, delta.StorageSnapshot) []byte); ok {
		r0 = rf(ctx, script, arguments, blockHeader, snapshot)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, []byte, [][]byte, *flow.Header, delta.StorageSnapshot) error); ok {
		r1 = rf(ctx, script, arguments, blockHeader, snapshot)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAccount provides a mock function with given fields: addr, header, snapshot
func (_m *ComputationManager) GetAccount(addr flow.Address, header *flow.Header, snapshot delta.StorageSnapshot) (*flow.Account, error) {
	ret := _m.Called(addr, header, snapshot)

	var r0 *flow.Account
	if rf, ok := ret.Get(0).(func(flow.Address, *flow.Header, delta.StorageSnapshot) *flow.Account); ok {
		r0 = rf(addr, header, snapshot)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*flow.Account)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(flow.Address, *flow.Header, delta.StorageSnapshot) error); ok {
		r1 = rf(addr, header, snapshot)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewComputationManager interface {
	mock.TestingT
	Cleanup(func())
}

// NewComputationManager creates a new instance of ComputationManager. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewComputationManager(t mockConstructorTestingTNewComputationManager) *ComputationManager {
	mock := &ComputationManager{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
