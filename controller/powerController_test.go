package controller

import (
	_ "HotelAutomation/model"
	. "HotelAutomation/service"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestShouldRegisterDevicesWithLightInOffStateAndACInOn(t *testing.T) {
	request := CreateHotelRequest{NumberOfFloors: 1, MainCorridorPerFloor: 1, SubCorridorPerFloor: 0}
	hotelService := NewHotelService()
	hotelService.CreateHotel(request)
	powerController := NewPowerController(hotelService)
	powerController.RegisterDevices()

	expectedHotelApplianceInfo := []AppliancesInfo{
		{
			Name:             "Light",
			Number:           1,
			IsSwitchedOd:     false,
			PowerConsumption: 5,
			Location:         ApplianceLocation{FloorNumber: 1, CorridorType: "Main", CorridorNumber: 1},
		},
		{
			Name:             "AC",
			Number:           1,
			IsSwitchedOd:     true,
			PowerConsumption: 10,
			Location:         ApplianceLocation{FloorNumber: 1, CorridorType: "Main", CorridorNumber: 1},
		},
	}

	assert.ElementsMatch(t, expectedHotelApplianceInfo, hotelService.GetAppliancesInfo())
}

func TestShouldTurnOnTheLightAtFloorOneMainCorridor(t *testing.T) {
	request := CreateHotelRequest{NumberOfFloors: 1, MainCorridorPerFloor: 1, SubCorridorPerFloor: 0}
	hotelService := NewHotelService()
	hotelService.CreateHotel(request)
	powerController := NewPowerController(hotelService)
	powerController.RegisterDevices()
	toggleApplianceRequest := ToggleApplianceRequest{
		ApplianceType: "Light", ApplianceNumber: 1, SwitchOn: true,
		Location: ApplianceLocation{FloorNumber: 1, CorridorType: "Main", CorridorNumber: 1},
	}
	err := powerController.ToggleApplianceState(toggleApplianceRequest)
	expectedHotelApplianceInfo := []AppliancesInfo{
		{
			Name:             "Light",
			Number:           1,
			IsSwitchedOd:     true,
			PowerConsumption: 5,
			Location:         ApplianceLocation{FloorNumber: 1, CorridorType: "Main", CorridorNumber: 1},
		},
		{
			Name:             "AC",
			Number:           1,
			IsSwitchedOd:     true,
			PowerConsumption: 10,
			Location:         ApplianceLocation{FloorNumber: 1, CorridorType: "Main", CorridorNumber: 1},
		},
	}
	assert.Nil(t, err)
	assert.ElementsMatch(t, expectedHotelApplianceInfo, hotelService.GetAppliancesInfo())
}
