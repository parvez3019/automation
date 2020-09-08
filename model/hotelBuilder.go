package model

import (
	. "HotelAutomation/model/appliances"
)

type HotelBuilder struct {
	*Hotel
}

func NewHotelBuilder() *HotelBuilder {
	return &HotelBuilder{NewHotel()}
}

func (hb *HotelBuilder) WithOneLightBulbAndOneACInEveryCorridor() *HotelBuilder {
	for _, corridor := range hb.GetAllCorridors() {
		// Not sure about the req of appliance number, we can modify as per requirements
		corridor.AddAirConditioner(NewAirConditioner(corridor.id, AcPowerConsumptionUnit))
		corridor.AddLightBulb(NewLightBulb(corridor.id, LightBulbPowerConsumptionUnit))
	}
	return hb
}

func (hb *HotelBuilder) Build() *Hotel {
	return hb.Hotel
}

func (hb *HotelBuilder) WithFloors(noOfFloor int) *HotelBuilder {
	floors := make([]*Floor, 0)
	for i := 1; i <= noOfFloor; i++ {
		floors = append(floors, NewFloor(i))
	}
	hb.AddFloors(floors)
	return hb
}

func (hb *HotelBuilder) WithCorridors(mainCorridorPerFloor int, subCorridorPerFloor int) *HotelBuilder {
	for _, floor := range hb.GetFloors() {
		floor.
			AddCorridors(createCorridors(MAIN, mainCorridorPerFloor)).
			AddCorridors(createCorridors(SUB, subCorridorPerFloor))
	}
	return hb
}

func createCorridors(cType CorridorType, count int) []*Corridor {
	corridors := make([]*Corridor, 0)
	for i := 1; i <= count; i++ {
		corridors = append(corridors, NewCorridor(cType, i))
	}
	return corridors
}

const AcPowerConsumptionUnit = 10
const LightBulbPowerConsumptionUnit = 5
