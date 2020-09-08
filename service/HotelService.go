package service

import (
	. "HotelAutomation/model"
)

type HotelService struct {
	hotel *Hotel
}

type HotelServiceI interface {
	CreateHotel(request CreateHotelRequest)
	GetAppliancesInfo() []AppliancesInfo
	GetAppliances() []Appliances
	GetNumberOfCorridors(int, CorridorType) int
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

func (hotelService *HotelService) GetAppliancesInfo() []AppliancesInfo {
	return mapApplianceToApplianceInfo(hotelService.GetAppliances())
}

func (hotelService *HotelService) GetAppliances() []Appliances {
	return mapToAppliances(hotelService.hotel.GetFloors())
}

func (hotelService *HotelService) GetNumberOfCorridors(floorNumber int, corridorType CorridorType) int {
	return len(hotelService.hotel.GetCorridorsAtFloor(floorNumber, corridorType))
}