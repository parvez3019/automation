package service

import (
	. "HotelAutomation/model"
	_ "HotelAutomation/model"
	. "HotelAutomation/model/appliances"
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
		AppType: "Light", TurnOn: true,
		Location: CorridorLocation{FloorNumber: 1, CorridorType: "Main", CorridorNumber: 1},
	}
	err := powerController.Update(toggleApplianceRequest)
	expectedHotelApplianceInfo := []AppliancesInfo{
		{
			Name:             "Light",
			Number:           1,
			IsSwitchedOn:     true,
			PowerConsumption: 5,
			Location:         CorridorLocation{FloorNumber: 1, CorridorType: "Main", CorridorNumber: 1},
		},
		{
			Name:             "AC",
			Number:           1,
			IsSwitchedOn:     true,
			PowerConsumption: 10,
			Location:         CorridorLocation{FloorNumber: 1, CorridorType: "Main", CorridorNumber: 1},
		},
	}
	assert.Nil(t, err)
	assert.ElementsMatch(t, expectedHotelApplianceInfo, hotelService.GetAppliancesInfo())
}

func TestShouldReturnErrorIfApplianceNotFound(t *testing.T) {
	request := CreateHotelRequest{NumberOfFloors: 1, MainCorridorPerFloor: 1, SubCorridorPerFloor: 0}
	_, powerController := setupHotelServiceAndRegisterDevicesToPowerController(request)

	toggleApplianceRequest := ToggleApplianceRequest{
		AppType: "Light", TurnOn: true,
		Location: CorridorLocation{FloorNumber: 2, CorridorType: "Main", CorridorNumber: 1},
	}
	err := powerController.Update(toggleApplianceRequest)
	assert.EqualError(t, err, "ApplianceNotFound")
}

func TestShouldReturnTotalPowerConsumptionByActiveAppliances(t *testing.T)  {
	request := CreateHotelRequest{NumberOfFloors: 1, MainCorridorPerFloor: 1, SubCorridorPerFloor: 1}
	_, pc := setupHotelServiceAndRegisterDevicesToPowerController(request)

	assert.Equal(t, 25, pc.TotalPowerConsumptionAtFloor(1))
}

func TestShouldToggleOffAnOnStateAppliance(t *testing.T) {
	request := CreateHotelRequest{NumberOfFloors: 1, MainCorridorPerFloor: 1, SubCorridorPerFloor: 2}
	_, powerController := setupHotelServiceAndRegisterDevicesToPowerController(request)
	locationKeySub1 := ApplianceLocationKey{
		aType: AC, location: CorridorLocation{FloorNumber: 1, CorridorType: SUB, CorridorNumber: 1},
	}
	locationKeySub2 := ApplianceLocationKey{
		aType: AC, location: CorridorLocation{FloorNumber: 1, CorridorType: SUB, CorridorNumber: 2},
	}
	powerController.devices[locationKeySub2].SetSwitchedOn(false)
	before := powerController.devices[locationKeySub1].IsOn()

	powerController.ToggleApplianceToReverseState(1, SUB, AC, false)

	assert.Equal(t, true, before)
	assert.Equal(t, false, powerController.devices[locationKeySub1].IsOn())
}

func TestShouldToggleOnAnOFFStateAppliance(t *testing.T) {
	request := CreateHotelRequest{NumberOfFloors: 1, MainCorridorPerFloor: 1, SubCorridorPerFloor: 2}
	_, powerController := setupHotelServiceAndRegisterDevicesToPowerController(request)
	locationKeySub1 := ApplianceLocationKey{
		aType:    AC,
		location: CorridorLocation{FloorNumber: 1, CorridorType: SUB, CorridorNumber: 1},
	}
	locationKeySub2 := ApplianceLocationKey{
		aType: AC, location: CorridorLocation{FloorNumber: 1, CorridorType: SUB, CorridorNumber: 2},
	}
	powerController.devices[locationKeySub2].SetSwitchedOn(true)
	appliance := powerController.devices[locationKeySub1]
	appliance.SetSwitchedOn(false)
	before := appliance.IsOn()

	powerController.ToggleApplianceToReverseState(1, SUB, AC, true)
	assert.Equal(t, false, before)
	assert.Equal(t, true, appliance.IsOn())
}

func getInfoAsMainCorridorLightOnAndSubCorridorOffAndBothCorridorACsInOnState() []AppliancesInfo {
	return []AppliancesInfo{
		{
			Name:             "Light",
			Number:           1,
			IsSwitchedOn:     true,
			PowerConsumption: 5,
			Location:         CorridorLocation{FloorNumber: 1, CorridorType: "Main", CorridorNumber: 1},
		},
		{
			Name:             "Light",
			Number:           1,
			IsSwitchedOn:     false,
			PowerConsumption: 5,
			Location:         CorridorLocation{FloorNumber: 1, CorridorType: "Sub", CorridorNumber: 1},
		},
		{
			Name:             "AC",
			Number:           1,
			IsSwitchedOn:     true,
			PowerConsumption: 10,
			Location:         CorridorLocation{FloorNumber: 1, CorridorType: "Main", CorridorNumber: 1},
		},
		{
			Name:             "AC",
			Number:           1,
			IsSwitchedOn:     true,
			PowerConsumption: 10,
			Location:         CorridorLocation{FloorNumber: 1, CorridorType: "Sub", CorridorNumber: 1},
		},
	}
}
