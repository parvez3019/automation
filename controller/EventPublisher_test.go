package controller

import (
	mock "HotelAutomation/_mocks"
	. "HotelAutomation/service"
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

func TestShouldNotifyAllObservers(t *testing.T) {
	publisher := Publisher{}
	powerController1 := &mock.Subscriber{}
	powerController2 := &mock.Subscriber{}
	powerController3 := &mock.Subscriber{}

	powerController1.On("Update", MovementDetectedEvent{}).Return()
	powerController2.On("Update", MovementDetectedEvent{}).Return()
	powerController3.On("Update", MovementDetectedEvent{}).Return()

	publisher.AddSubscriber(powerController1)
	publisher.AddSubscriber(powerController2)
	publisher.AddSubscriber(powerController3)

	publisher.NotifyAll(MovementDetectedEvent{})

	assert.Equal(t, 1, len(powerController1.Calls))
	assert.Equal(t, 1, len(powerController2.Calls))
	assert.Equal(t, 1, len(powerController3.Calls))
}
