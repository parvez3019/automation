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
		corridor.AddAirConditioner(NewAirConditioner(1, AcPowerConsumptionUnit))
		corridor.AddLightBulb(NewLightBulb(1, LightBulbPowerConsumptionUnit))
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
			AddCorridors(createCorridors(MAIN, mainCorridorPerFloor), MAIN).
			AddCorridors(createCorridors(SUB, subCorridorPerFloor), SUB)
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
