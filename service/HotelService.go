package service

import (
	. "HotelAutomation/model"
	"fmt"
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

func (hotelService *HotelService) PrintHotelApplianceState() {
	for _, floor := range hotelService.hotel.GetFloors() {
		fmt.Printf("Floor %d\n", floor.GetLevel())
		printCorridor(floor, MAIN)
		printCorridor(floor, SUB)
	}
}

func printCorridor(floor *Floor, corridorType CorridorType) {
	for _, corridor := range floor.GetCorridors(corridorType) {
		fmt.Printf("%s Corridor %d ", corridor.GetType(), corridor.GetId())
		for _, bulb := range corridor.GetLightBulbs() {
			fmt.Printf("Light %d : %s ", bulb.GetId(), toString(bulb.IsOn()))
		}
		for _, ac := range corridor.GetAirConditioners() {
			fmt.Printf("AC : %s", toString(ac.IsOn()))
		}
		fmt.Println()
	}
}

func toString(on bool) string {
	if on {
		return "ON"
	}
	return "OFF"
}
