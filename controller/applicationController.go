package controller

import (
	. "HotelAutomation/service"
)

type PowerAutomationController struct {
	hotelService    *HotelService
	powerController *PowerController
}

func NewApplicationController(hotelService *HotelService) *PowerAutomationController {
	powerController := NewPowerController(hotelService)
	return &PowerAutomationController{hotelService: hotelService, powerController: powerController}
}

func (c *PowerAutomationController) Init(request CreateHotelRequest) {
	c.hotelService.CreateHotel(request)
	c.powerController.RegisterDevices()
}

func (c *PowerAutomationController) ToggleAppliance(request ToggleApplianceRequest) error {
	return c.powerController.ToggleApplianceState(request)
}
