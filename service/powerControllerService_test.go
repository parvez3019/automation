package service

import (
	_ "HotelAutomation/model"
	"github.com/stretchr/testify/assert"
	"testing"
)

func setupHotelServiceAndRegisterDevicesToPowerController(request CreateHotelRequest) (*HotelService, *PowerControllerService) {
	hotelService := NewHotelService()
	hotelService.CreateHotel(request)
	powerController := NewPowerControllerService(hotelService)
	powerController.RegisterDevices()
	return hotelService, powerController
}

func TestShouldRegisterDevicesWithMainCorridorLightOnAndSubCorridorOffAndACInOn(t *testing.T) {
	request := CreateHotelRequest{NumberOfFloors: 1, MainCorridorPerFloor: 1, SubCorridorPerFloor: 1}
	hotelService, _ := setupHotelServiceAndRegisterDevicesToPowerController(request)

	expectedHotelApplianceInfo := getInfoAsMainCorridorLightOnAndSubCorridorOffAndBothCorridorACsInOnState()

	assert.ElementsMatch(t, expectedHotelApplianceInfo, hotelService.GetAppliancesInfo())
}

func TestShouldTurnOnTheLightAtFloorOneMainCorridor(t *testing.T) {
	request := CreateHotelRequest{NumberOfFloors: 1, MainCorridorPerFloor: 1, SubCorridorPerFloor: 0}
	hotelService, powerController := setupHotelServiceAndRegisterDevicesToPowerController(request)

	toggleApplianceRequest := ToggleApplianceRequest{
		ApplianceType: "Light", SwitchOn: true,
		Location: CorridorLocation{FloorNumber: 1, CorridorType: "Main", CorridorNumber: 1},
	}
	err := powerController.Update(toggleApplianceRequest)
	expectedHotelApplianceInfo := []AppliancesInfo{
		{
			Name:             "Light",
			Number:           1,
			IsSwitchedOd:     true,
			PowerConsumption: 5,
			Location:         CorridorLocation{FloorNumber: 1, CorridorType: "Main", CorridorNumber: 1},
		},
		{
			Name:             "AC",
			Number:           1,
			IsSwitchedOd:     true,
			PowerConsumption: 10,
			Location:         CorridorLocation{FloorNumber: 1, CorridorType: "Main", CorridorNumber: 1},
		},
	}
	assert.Nil(t, err)
	assert.ElementsMatch(t, expectedHotelApplianceInfo, hotelService.GetAppliancesInfo())
}

func getInfoAsMainCorridorLightOnAndSubCorridorOffAndBothCorridorACsInOnState() []AppliancesInfo {
	return []AppliancesInfo{
		{
			Name:             "Light",
			Number:           1,
			IsSwitchedOd:     true,
			PowerConsumption: 5,
			Location:         CorridorLocation{FloorNumber: 1, CorridorType: "Main", CorridorNumber: 1},
		},
		{
			Name:             "Light",
			Number:           1,
			IsSwitchedOd:     false,
			PowerConsumption: 5,
			Location:         CorridorLocation{FloorNumber: 1, CorridorType: "Sub", CorridorNumber: 1},
		},
		{
			Name:             "AC",
			Number:           1,
			IsSwitchedOd:     true,
			PowerConsumption: 10,
			Location:         CorridorLocation{FloorNumber: 1, CorridorType: "Main", CorridorNumber: 1},
		},
		{
			Name:             "AC",
			Number:           1,
			IsSwitchedOd:     true,
			PowerConsumption: 10,
			Location:         CorridorLocation{FloorNumber: 1, CorridorType: "Sub", CorridorNumber: 1},
		},
	}
}
