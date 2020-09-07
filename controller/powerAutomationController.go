package controller

import (
	. "HotelAutomation/model"
	. "HotelAutomation/model/appliances"
	. "HotelAutomation/service"
	"fmt"
)

type PowerAutomationController struct {
	hotelService     *HotelService
	powerController  *PowerControllerService
	motionController *MotionController
}

func NewPowerAutomationController(hotelService *HotelService) *PowerAutomationController {
	powerController := NewPowerControllerService(hotelService)
	motionController := NewMotionController()
	return &PowerAutomationController{
		hotelService:     hotelService,
		powerController:  powerController,
		motionController: motionController,
	}
}

func (c *PowerAutomationController) Init(request CreateHotelRequest) {
	c.hotelService.CreateHotel(request)
	c.powerController.RegisterDevices()
	c.motionController.addObserver(c.powerController)
}

func (c *PowerAutomationController) MovementDetected(atLocation ApplianceLocation, movement bool) {
	turnOnLightAtLocationRequest := ToggleApplianceRequest{ApplianceType: LIGHT, SwitchOn: movement, Location: atLocation}
	err := c.motionController.NotifyAll(turnOnLightAtLocationRequest)
	if err != nil {
		fmt.Println(err.Error())
	}
	c.verifyTotalPowerConsumptionAtFloor(atLocation)
}

func (c *PowerAutomationController) toggleAppliance(request ToggleApplianceRequest) error {
	return c.powerController.Update(request)
}

func (c *PowerAutomationController) verifyTotalPowerConsumptionAtFloor(atLocation ApplianceLocation) {
	if c.totalPowerConsumptionAtFloorExceeded(atLocation.FloorNumber) {
		c.toggleSubCorridorAC(atLocation, false)
		return
	}
	c.toggleSubCorridorAC(atLocation, true)
}

func (c *PowerAutomationController) totalPowerConsumptionAtFloorExceeded(floorNumber int) bool {
	totalConsumption := c.powerController.TotalPowerConsumptionAtFloor(floorNumber)
	totalMainCorridors := c.hotelService.GetNumberOfCorridors(MAIN)
	totalSubCorridors := c.hotelService.GetNumberOfCorridors(SUB)
	return totalConsumption >= (totalMainCorridors*15)+(totalSubCorridors*10)
}

func (c *PowerAutomationController) toggleSubCorridorAC(atLocation ApplianceLocation, switchOn bool) {
	_ = c.toggleAppliance(ToggleApplianceRequest{
		ApplianceType: AC,
		SwitchOn:      switchOn,
		Location:      atLocation,
	})
}
