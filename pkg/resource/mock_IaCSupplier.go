// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package resource

import (
	"github.com/snyk/driftctl/enumeration/resource"
	mock "github.com/stretchr/testify/mock"
)

// MockIaCSupplier is an autogenerated mock type for the IaCSupplier type
type MockIaCSupplier struct {
	mock.Mock
}

// Resources provides a mock function with given fields:
func (_m *MockIaCSupplier) Resources() ([]*resource.Resource, error) {
	ret := _m.Called()

	var r0 []*resource.Resource
	if rf, ok := ret.Get(0).(func() []*resource.Resource); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*resource.Resource)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SourceCount provides a mock function with given fields:
func (_m *MockIaCSupplier) SourceCount() uint {
	ret := _m.Called()

	var r0 uint
	if rf, ok := ret.Get(0).(func() uint); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(uint)
	}

	return r0
}
