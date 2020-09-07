package controller

import (
	. "HotelAutomation/service"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
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
	powerController1 := new(MockPowerController)
	powerController2 := new(MockPowerController)
	powerController3 := new(MockPowerController)

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

type MockPowerController struct {
	mock.Mock
}

func (_m *MockPowerController) Update(request ToggleApplianceRequest) error {
	ret := _m.Called(request)

	var r0 error
	if rf, ok := ret.Get(0).(func(ToggleApplianceRequest) error); ok {
		r0 = rf(request)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
