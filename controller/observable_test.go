package controller

import (
	mock "HotelAutomation/_mocks"
	. "HotelAutomation/service"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestShouldAddObservers(t *testing.T) {
	observable := Observable{}
	observable.addObserver(NewPowerControllerService(nil))
	observable.addObserver(NewPowerControllerService(nil))
	observable.addObserver(NewPowerControllerService(nil))
	assert.Len(t, observable.Observers, 3)
}

func TestShouldNotifyAllObservers(t *testing.T) {
	observable := Observable{}
	powerController1 := &mock.ObserverI{}
	powerController2 := &mock.ObserverI{}
	powerController3 := &mock.ObserverI{}

	powerController1.On("Update", ToggleApplianceRequest{}).Return(nil)
	powerController2.On("Update", ToggleApplianceRequest{}).Return(nil)
	powerController3.On("Update", ToggleApplianceRequest{}).Return(nil)

	observable.addObserver(powerController1)
	observable.addObserver(powerController2)
	observable.addObserver(powerController3)

	err := observable.NotifyAll(ToggleApplianceRequest{})

	assert.Nil(t, err)
	assert.Equal(t, 1, len(powerController1.Calls))
	assert.Equal(t, 1, len(powerController2.Calls))
	assert.Equal(t, 1, len(powerController3.Calls))
}