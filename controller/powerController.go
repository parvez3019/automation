package controller

import (
	. "HotelAutomation/model/appliances"
	. "HotelAutomation/service"
	"errors"
)

type PowerController struct {
	hotelService *HotelService
	devices      map[ApplianceLocationKey]ApplianceStateI
}

func NewPowerController(hotelService *HotelService, ) *PowerController {
	return &PowerController{hotelService: hotelService, devices: make(map[ApplianceLocationKey]ApplianceStateI, 0)}
}

func (c *PowerController) RegisterDevices() {
	for _, appliance := range c.hotelService.GetAppliances() {
		applianceLocationKey := createKeyFromApplianceStateI(appliance)
		c.setApplianceToInitialSwitchedState(appliance.Appliance)
		c.devices[applianceLocationKey] = appliance.Appliance
	}
}

func (c *PowerController) ToggleApplianceState(request ToggleApplianceRequest) error {
	applianceLocationKey := createApplianceLocationKeyFromToggleRequest(request)
	if appliance, ok := c.devices[applianceLocationKey]; ok {
		appliance.SetSwitchedOn(request.SwitchOn)
		return nil
	}
	return errors.New("ApplianceNotFound")
}

func (c *PowerController) setApplianceToInitialSwitchedState(appliance ApplianceStateI) {
	if appliance.GetType() == string(LIGHT) {
		appliance.SetSwitchedOn(false)
	}
	if appliance.GetType() == string(AC) {
		appliance.SetSwitchedOn(true)
	}
}

func createKeyFromApplianceStateI(a Appliances) ApplianceLocationKey {
	return ApplianceLocationKey{
		aType:    a.Appliance.GetType(),
		number:   a.Appliance.GetId(),
		location: a.Location,
	}
}

func createApplianceLocationKeyFromToggleRequest(request ToggleApplianceRequest) ApplianceLocationKey {
	return ApplianceLocationKey{
		aType:    string(request.ApplianceType),
		number:   request.ApplianceNumber,
		location: request.Location,
	}
}

type ApplianceLocationKey struct {
	aType    string
	number   int
	location ApplianceLocation
}
