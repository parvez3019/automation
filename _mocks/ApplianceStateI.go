// Code generated by mockery v2.2.1. DO NOT EDIT.

package service

import mock "github.com/stretchr/testify/mock"

// ApplianceStateI is an autogenerated mock type for the ApplianceStateI type
type ApplianceStateI struct {
	mock.Mock
}

// GetId provides a mock function with given fields:
func (_m *ApplianceStateI) GetId() int {
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
func (_m *ApplianceStateI) GetPowerConsumption() int {
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
func (_m *ApplianceStateI) GetType() string {
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
func (_m *ApplianceStateI) IsOn() bool {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// SetSwitchedOn provides a mock function with given fields: _a0
func (_m *ApplianceStateI) SetSwitchedOn(_a0 bool) {
	_m.Called(_a0)
}
