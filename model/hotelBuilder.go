package model

import . "HotelAutomation/model/appliances"

type HotelBuilder struct {
	*Hotel
}

func NewHotelBuilder() *HotelBuilder {
	return &HotelBuilder{NewHotel()}
}

func (hb *HotelBuilder) WithFloorsAndCorridors(request CreateHotelRequest) *HotelBuilder {
	hb.addFloors(buildFloorsWithCorridors(request))
	return hb
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

func buildFloorsWithCorridors(request CreateHotelRequest) []*Floor {
	floors := make([]*Floor, 0)
	for i := 1; i <= request.NumberOfFloors; i++ {
		currentFloor := NewFloor(i).
			addCorridors(createCorridors(MAIN, request.MainCorridorPerFloor), MAIN).
			addCorridors(createCorridors(SUB, request.SubCorridorPerFloor), SUB)
		floors = append(floors, currentFloor)
	}
	return floors
}

func createCorridors(cType Type, count int) []*Corridor {
	corridors := make([]*Corridor, 0)
	for i := 1; i <= count; i++ {
		corridors = append(corridors, NewCorridor(cType, i))
	}
	return corridors
}

type CreateHotelRequest struct {
	NumberOfFloors       int
	MainCorridorPerFloor int
	SubCorridorPerFloor  int
}
