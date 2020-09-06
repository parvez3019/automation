package model

type HotelBuilder struct{}

func NewHotelBuilder() *HotelBuilder {
	return &HotelBuilder{}
}

func (hb *HotelBuilder) BuildHotel(request CreateHotelRequest) *Hotel {
	return NewHotel().
		addFloors(buildFloorsWithCorridors(request))
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
