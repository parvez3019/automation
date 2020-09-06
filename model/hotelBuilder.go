package model

import . "HotelAutomation/model/appliances"

type HotelBuilder struct {
	*Hotel
}

func NewHotelBuilder() *HotelBuilder {
	return &HotelBuilder{NewHotel()}
}

func (hb *HotelBuilder) WithOneLightBulbAndOneACInEveryCorridor() *HotelBuilder {
	for _, corridor := range hb.getCorridors() {
		corridor.addAirConditioner(NewAirConditioner(1, 10))
		corridor.addLightBulb(NewLightBulb(1, 5))
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
	hb.addFloors(floors)
	return hb
}

func (hb *HotelBuilder) WithCorridors(mainCorridorPerFloor int, subCorridorPerFloor int) *HotelBuilder {
	for _, floor := range hb.getFloors() {
		floor.
			addCorridors(createCorridors(MAIN, mainCorridorPerFloor), MAIN).
			addCorridors(createCorridors(SUB, subCorridorPerFloor), SUB)
	}
	return hb
}

func createCorridors(cType Type, count int) []*Corridor {
	corridors := make([]*Corridor, 0)
	for i := 1; i <= count; i++ {
		corridors = append(corridors, NewCorridor(cType, i))
	}
	return corridors
}
