package controller

import (
	mock "HotelAutomation/_mocks"
	. "HotelAutomation/service"
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestShouldAddObservers(t *testing.T) {
	publisher := Publisher{}
	publisher.AddSubscriber(&mock.Subscriber{})
	publisher.AddSubscriber(&mock.Subscriber{})
	publisher.AddSubscriber(&mock.Subscriber{})
	assert.Len(t, publisher.subscribers, 3)
}

func TestShouldReturnErrorIfUpdateFails(t *testing.T) {
	publisher := Publisher{}
	powerController := &mock.Subscriber{}
	powerController.On("Update", MovementDetectedEvent{}).Return(errors.New("SomeErr"))
	publisher.AddSubscriber(powerController)

	err := publisher.NotifyAll(MovementDetectedEvent{})

	assert.EqualError(t, err, "SomeErr")
	assert.Equal(t, 1, len(powerController.Calls))
}


func TestShouldNotifyAllObservers(t *testing.T) {
	publisher := Publisher{}
	powerController1 := &mock.Subscriber{}
	powerController2 := &mock.Subscriber{}
	powerController3 := &mock.Subscriber{}

	powerController1.On("Update", MovementDetectedEvent{}).Return(nil)
	powerController2.On("Update", MovementDetectedEvent{}).Return(nil)
	powerController3.On("Update", MovementDetectedEvent{}).Return(nil)

	publisher.AddSubscriber(powerController1)
	publisher.AddSubscriber(powerController2)
	publisher.AddSubscriber(powerController3)

	err := publisher.NotifyAll(MovementDetectedEvent{})

	assert.Nil(t, err)
	assert.Equal(t, 1, len(powerController1.Calls))
	assert.Equal(t, 1, len(powerController2.Calls))
	assert.Equal(t, 1, len(powerController3.Calls))
}
