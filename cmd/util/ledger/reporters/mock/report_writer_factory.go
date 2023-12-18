// Code generated by mockery v2.21.4. DO NOT EDIT.

package mock

import (
	mock "github.com/stretchr/testify/mock"

	reporters "github.com/onflow/flow-go/cmd/util/ledger/reporters"
)

// ReportWriterFactory is an autogenerated mock type for the ReportWriterFactory type
type ReportWriterFactory struct {
	mock.Mock
}

// ReportWriter provides a mock function with given fields: dataNamespace
func (_m *ReportWriterFactory) ReportWriter(dataNamespace string) reporters.ReportWriter {
	ret := _m.Called(dataNamespace)

	var r0 reporters.ReportWriter
	if rf, ok := ret.Get(0).(func(string) reporters.ReportWriter); ok {
		r0 = rf(dataNamespace)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(reporters.ReportWriter)
		}
	}

	return r0
}

type mockConstructorTestingTNewReportWriterFactory interface {
	mock.TestingT
	Cleanup(func())
}

// NewReportWriterFactory creates a new instance of ReportWriterFactory. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewReportWriterFactory(t mockConstructorTestingTNewReportWriterFactory) *ReportWriterFactory {
	mock := &ReportWriterFactory{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
