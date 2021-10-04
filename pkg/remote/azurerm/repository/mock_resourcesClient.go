// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package repository

import (
	armresources "github.com/Azure/azure-sdk-for-go/sdk/resources/armresources"
	mock "github.com/stretchr/testify/mock"
)

// mockResourcesClient is an autogenerated mock type for the resourcesClient type
type mockResourcesClient struct {
	mock.Mock
}

// List provides a mock function with given fields: options
func (_m *mockResourcesClient) List(options *armresources.ResourceGroupsListOptions) resourcesListPager {
	ret := _m.Called(options)

	var r0 resourcesListPager
	if rf, ok := ret.Get(0).(func(*armresources.ResourceGroupsListOptions) resourcesListPager); ok {
		r0 = rf(options)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(resourcesListPager)
		}
	}

	return r0
}
