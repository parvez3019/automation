package controller

import (
	mock "HotelAutomation/_mocks"
	. "HotelAutomation/service"
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestShouldReturnError(t *testing.T) {
	motionController := NewMotionController()
	powerController := &mock.Subscriber{}
	motionController.AddSubscriber(powerController)
	powerController.On("Update", MovementDetectedEvent{}).Return(errors.New("err"))

	err := motionController.RaiseMotionDetectedEvent(MovementDetectedEvent{})

	assert.EqualError(t, err, "err")
	assert.Equal(t, 1, len(powerController.Calls))
}

func TestShouldNotifyAll(t *testing.T) {
	motionController := NewMotionController()
	powerController1 := &mock.Subscriber{}
	powerController2 := &mock.Subscriber{}
	powerController3 := &mock.Subscriber{}

	motionController.AddSubscriber(powerController1)
	motionController.AddSubscriber(powerController2)
	motionController.AddSubscriber(powerController3)

	powerController1.On("Update", MovementDetectedEvent{}).Return(nil)
	powerController2.On("Update", MovementDetectedEvent{}).Return(nil)
	powerController3.On("Update", MovementDetectedEvent{}).Return(nil)

	err := motionController.RaiseMotionDetectedEvent(MovementDetectedEvent{})

	assert.Nil(t, err)
	assert.Equal(t, 1, len(powerController1.Calls))
	assert.Equal(t, 1, len(powerController2.Calls))
	assert.Equal(t, 1, len(powerController3.Calls))
}
