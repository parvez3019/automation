// Code generated by mockery v2.2.1. DO NOT EDIT.

package service

import (
	service "HotelAutomation/service"

	mock "github.com/stretchr/testify/mock"
)

// Subscriber is an autogenerated mock type for the Subscriber type
type Subscriber struct {
	mock.Mock
}

// Update provides a mock function with given fields: _a0
func (_m *Subscriber) Update(_a0 service.MovementDetectedEvent) {
	_m.Called(_a0)
}
