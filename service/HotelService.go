package service

import . "HotelAutomation/model"

type HotelService struct {
	hotel *Hotel
}

func NewHotelService() *HotelService {
	return &HotelService{}
}

func (hotelService *HotelService) CreateHotel(request CreateHotelRequest) {
	hotelService.hotel = NewHotelBuilder().
		WithFloors(request.NumberOfFloors).
		WithCorridors(request.MainCorridorPerFloor, request.SubCorridorPerFloor).
		WithOneLightBulbAndOneACInEveryCorridor().
		Build()
}

func (hotelService *HotelService) PrintHotelAppliancesState() string {
	return "Hello"
}

type CreateHotelRequest struct {
	NumberOfFloors       int
	MainCorridorPerFloor int
	SubCorridorPerFloor  int
}
