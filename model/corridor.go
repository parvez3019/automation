package model

import (
	. "HotelAutomation/model/appliances"
)

type Corridor struct {
	id         int
	cType      CorridorType
	appliances []ApplianceStateI
}

func NewCorridor(corridorType CorridorType, id int) *Corridor {
	return &Corridor{
		cType:      corridorType,
		id:         id,
		appliances: make([]ApplianceStateI, 0),
	}
}

func (c *Corridor) AddLightBulb(bulb *LightBulb) *Corridor {
	c.appliances = append(c.appliances, bulb)
	return c
}

func (c *Corridor) AddAirConditioner(ac *AirConditioner) *Corridor {
	c.appliances = append(c.appliances, ac)
	return c
}

func (c *Corridor) GetAppliances(applianceType ApplianceType) []ApplianceStateI {
	appliances := make([]ApplianceStateI, 0)
	for _, appliance := range c.appliances {
		if appliance.GetType() == applianceType {
			appliances = append(appliances, appliance)
		}
	}
	return appliances
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
