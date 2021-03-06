// Code generated by mockery v1.0.0. DO NOT EDIT.
package mocks

import discovery "k8s.io/client-go/discovery"
import mock "github.com/stretchr/testify/mock"
import runtime "k8s.io/apimachinery/pkg/runtime"

// ObjectInfo is an autogenerated mock type for the ObjectInfo type
type ObjectInfo struct {
	mock.Mock
}

// ResourceName provides a mock function with given fields: d, o
func (_m *ObjectInfo) ResourceName(d discovery.ServerResourcesInterface, o runtime.Object) string {
	ret := _m.Called(d, o)

	var r0 string
	if rf, ok := ret.Get(0).(func(discovery.ServerResourcesInterface, runtime.Object) string); ok {
		r0 = rf(d, o)
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}
