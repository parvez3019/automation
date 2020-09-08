package controller

import . "HotelAutomation/service"

type Publisher struct {
	subscribers []Subscriber
}

func (o *Publisher) AddSubscriber(obs Subscriber) {
	o.subscribers = append(o.subscribers, obs)
}

func (o *Publisher) NotifyAll(request MovementDetectedEvent) {
	for _, ob := range o.subscribers {
		ob.Update(request)
	}
}
