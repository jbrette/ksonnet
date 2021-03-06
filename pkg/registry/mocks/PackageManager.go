// Copyright 2018 The ksonnet authors
//
//
//    Licensed under the Apache License, Version 2.0 (the "License");
//    you may not use this file except in compliance with the License.
//    You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
//    Unless required by applicable law or agreed to in writing, software
//    distributed under the License is distributed on an "AS IS" BASIS,
//    WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//    See the License for the specific language governing permissions and
//    limitations under the License.

// Code generated by mockery v1.0.0. DO NOT EDIT.
package mocks

import mock "github.com/stretchr/testify/mock"
import pkg "github.com/ksonnet/ksonnet/pkg/pkg"
import prototype "github.com/ksonnet/ksonnet/pkg/prototype"

// PackageManager is an autogenerated mock type for the PackageManager type
type PackageManager struct {
	mock.Mock
}

// Find provides a mock function with given fields: _a0
func (_m *PackageManager) Find(_a0 string) (pkg.Package, error) {
	ret := _m.Called(_a0)

	var r0 pkg.Package
	if rf, ok := ret.Get(0).(func(string) pkg.Package); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(pkg.Package)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Packages provides a mock function with given fields:
func (_m *PackageManager) Packages() ([]pkg.Package, error) {
	ret := _m.Called()

	var r0 []pkg.Package
	if rf, ok := ret.Get(0).(func() []pkg.Package); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]pkg.Package)
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

// Prototypes provides a mock function with given fields:
func (_m *PackageManager) Prototypes() (prototype.Prototypes, error) {
	ret := _m.Called()

	var r0 prototype.Prototypes
	if rf, ok := ret.Get(0).(func() prototype.Prototypes); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(prototype.Prototypes)
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
