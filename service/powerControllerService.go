package service

import (
	. "HotelAutomation/model"
	. "HotelAutomation/model/appliances"
	"errors"
)

type PowerControllerServiceI interface {
	RegisterDevices()
	Update(request ToggleApplianceRequest) error
	TotalPowerConsumptionAtFloor(floorNumber int) int
}

type PowerControllerService struct {
	hotelService HotelServiceI
	devices      map[ApplianceLocationKey]ApplianceStateI
}

func NewPowerControllerService(hotelService HotelServiceI) *PowerControllerService {
	devices := make(map[ApplianceLocationKey]ApplianceStateI, 0)
	return &PowerControllerService{hotelService: hotelService, devices: devices}
}

func (c *PowerControllerService) RegisterDevices() {
	for _, appliance := range c.hotelService.GetAppliances() {
		applianceLocationKey := createKeyFromApplianceStateI(appliance)
		c.setApplianceToInitialSwitchedState(appliance.Appliance, appliance.Location.CorridorType)
		c.devices[applianceLocationKey] = appliance.Appliance
	}
}

func (c *PowerControllerService) Update(request ToggleApplianceRequest) error {
	applianceLocationKey := createApplianceLocationKeyFromToggleRequest(request)
	if appliance, ok := c.devices[applianceLocationKey]; ok {
		appliance.SetSwitchedOn(request.TurnOn)
		return nil
	}
	return errors.New("ApplianceNotFound")
}

func (c *PowerControllerService) TotalPowerConsumptionAtFloor(floorNumber int) int {
	totalPower := 0
	for location, appliance := range c.devices {
		if location.location.FloorNumber == floorNumber && appliance.IsOn() {
			totalPower += appliance.GetPowerConsumption()
		}
	}
	return totalPower
}

func (c *PowerControllerService) setApplianceToInitialSwitchedState(appliance ApplianceStateI, cType CorridorType) {
	if appliance.GetType() == string(LIGHT) && cType == SUB {
		appliance.SetSwitchedOn(false)
	}
	if appliance.GetType() == string(LIGHT) && cType == MAIN {
		appliance.SetSwitchedOn(true)
	}
	if appliance.GetType() == string(AC) {
		appliance.SetSwitchedOn(true)
	}
}

func createKeyFromApplianceStateI(a Appliances) ApplianceLocationKey {
	return ApplianceLocationKey{
		aType:    a.Appliance.GetType(),
		location: a.Location,
	}
}

func createApplianceLocationKeyFromToggleRequest(request ToggleApplianceRequest) ApplianceLocationKey {
	return ApplianceLocationKey{
		aType:    string(request.AppType),
		location: request.Location,
	}
}

type ApplianceLocationKey struct {
	aType    string
	location CorridorLocation
}
