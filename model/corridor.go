package model

import (
	. "HotelAutomation/model/appliances"
)

type Corridor struct {
	id              int
	cType           CorridorType
	lightBulbs      []*LightBulb
	airConditioners []*AirConditioner
}

func NewCorridor(corridorType CorridorType, id int) *Corridor {
	lightBulbs := make([]*LightBulb, 0)
	airConditioners := make([]*AirConditioner, 0)
	return &Corridor{
		cType:           corridorType,
		id:              id,
		lightBulbs:      lightBulbs,
		airConditioners: airConditioners,
	}
}

func (c *Corridor) AddLightBulb(bulb *LightBulb) *Corridor {
	c.lightBulbs = append(c.lightBulbs, bulb)
	return c
}

func (c *Corridor) AddAirConditioner(ac *AirConditioner) *Corridor {
	c.airConditioners = append(c.airConditioners, ac)
	return c
}

func (c *Corridor) GetLightBulbs() []*LightBulb {
	return c.lightBulbs
}

func (c *Corridor) GetAirConditioners() []*AirConditioner {
	return c.airConditioners
}

func (c *Corridor) GetType() CorridorType {
	return c.cType
}

func (c *Corridor) GetId() int {
	return c.id
}

type CorridorType string

const (
	MAIN CorridorType = "Main"
	SUB  CorridorType = "Sub"
)
