package controller

import . "HotelAutomation/service"

type Observable struct {
	Observers []ObserverI
}

func (o *Observable) addObserver(obs ObserverI) {
	o.Observers = append(o.Observers, obs)
}

func (o *Observable) NotifyAll(request ToggleApplianceRequest) error {
	for _, ob := range o.Observers {
		err := ob.Update(request)
		if err != nil {
			return err
		}
	}
	return nil
}
