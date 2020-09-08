package service

import (
	. "HotelAutomation/model"
	. "HotelAutomation/model/appliances"
	"errors"
)

type PowerControllerServiceI interface {
	RegisterDevices()
	Update(ToggleApplianceRequest) error
	TotalPowerConsumptionAtFloor(int) int
	ToggleApplianceToReverseState(int, CorridorType, ApplianceType, bool)
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

func (c *PowerControllerService) ToggleApplianceToReverseState(floor int,
	corridorType CorridorType, applianceType ApplianceType, toggleTo bool) {
	for key, a := range c.devices {
		if matchLocationAndType(key, floor, corridorType, a, applianceType) && a.IsOn() == !toggleTo {
			a.SetSwitchedOn(toggleTo)
			return
		}
	}
}

func matchLocationAndType(key ApplianceLocationKey, floor int, corridorType CorridorType, a ApplianceStateI, applianceType ApplianceType) bool {
	return key.location.FloorNumber == floor && key.location.CorridorType == corridorType && a.GetType() == applianceType
}

func (c *PowerControllerService) setApplianceToInitialSwitchedState(appliance ApplianceStateI, cType CorridorType) {
	if appliance.GetType() == LIGHT && cType == SUB {
		appliance.SetSwitchedOn(false)
	}
	if appliance.GetType() == LIGHT && cType == MAIN {
		appliance.SetSwitchedOn(true)
	}
	if appliance.GetType() == AC {
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
		aType:    request.AppType,
		location: request.Location,
	}
}

type ApplianceLocationKey struct {
	aType    ApplianceType
	location CorridorLocation
}
