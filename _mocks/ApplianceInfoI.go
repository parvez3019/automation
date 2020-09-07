// Code generated by mockery v2.2.1. DO NOT EDIT.

package service

import mock "github.com/stretchr/testify/mock"

// ApplianceInfoI is an autogenerated mock type for the ApplianceInfoI type
type ApplianceInfoI struct {
	mock.Mock
}

// GetId provides a mock function with given fields:
func (_m *ApplianceInfoI) GetId() int {
	ret := _m.Called()

	var r0 int
	if rf, ok := ret.Get(0).(func() int); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(int)
	}

	return r0
}

// GetPowerConsumption provides a mock function with given fields:
func (_m *ApplianceInfoI) GetPowerConsumption() int {
	ret := _m.Called()

	var r0 int
	if rf, ok := ret.Get(0).(func() int); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(int)
	}

	return r0
}

// GetType provides a mock function with given fields:
func (_m *ApplianceInfoI) GetType() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// IsOn provides a mock function with given fields:
func (_m *ApplianceInfoI) IsOn() bool {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}
