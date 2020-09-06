package model

import (
	. "HotelAutomation/model/appliances"
)

type Corridor struct {
	id              int
	cType           Type
	lightBulbs      []*LightBulb
	airConditioners []*AirConditioner
}

func NewCorridor(corridorType Type, number int) *Corridor {

	lightBulbs := make([]*LightBulb, 0)
	airConditioners := make([]*AirConditioner, 0)
	return &Corridor{
		cType:           corridorType,
		id:              number,
		lightBulbs:      lightBulbs,
		airConditioners: airConditioners,
	}
}

func (c *Corridor) addLightBulb(bulb *LightBulb) *Corridor {
	c.lightBulbs = append(c.lightBulbs, bulb)
	return c
}

func (c *Corridor) addAirConditioner(ac *AirConditioner) *Corridor {
	c.airConditioners = append(c.airConditioners, ac)
	return c
}

func (c *Corridor) getLightBulbs() []*LightBulb {
	return c.lightBulbs
}

func (c *Corridor) getAirConditioners() []*AirConditioner {
	return c.airConditioners
}

func (c *Corridor) getTypeAsString() string {
	return string(c.cType)
}

type Type string

const (
	MAIN Type = "Main"
	SUB  Type = "Sub"
)
