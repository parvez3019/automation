package controller

import (
	. "HotelAutomation/model"
	. "HotelAutomation/model/appliances"
	. "HotelAutomation/service"
	"fmt"
)

type PowerAutomationController struct {
	hotelService     HotelServiceI
	powerController  PowerControllerServiceI
	motionController *MotionController
}

func NewPowerAutomationController(
	hService HotelServiceI,
	pController PowerControllerServiceI,
	mController *MotionController,
) *PowerAutomationController {
	return &PowerAutomationController{
		hotelService:     hService,
		powerController:  pController,
		motionController: mController,
	}
}

func (c *PowerAutomationController) Init(request CreateHotelRequest) {
	c.hotelService.CreateHotel(request)
	c.powerController.RegisterDevices()
	c.motionController.addObserver(c.powerController)
	//fmt.Println(c.hotelService.GetAppliancesInfo())
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
	return totalConsumption >=
		(totalMainCorridors*MainCorridorPowerConsumptionThresholdMultiplier)+
			(totalSubCorridors*SubCorridorPowerConsumptionThresholdMultiplier)
}

func (c *PowerAutomationController) toggleSubCorridorAC(atLocation ApplianceLocation, switchOn bool) {
	_ = c.toggleAppliance(ToggleApplianceRequest{
		ApplianceType: AC,
		SwitchOn:      switchOn,
		Location:      atLocation,
	})
}

const (
	MainCorridorPowerConsumptionThresholdMultiplier = 15
	SubCorridorPowerConsumptionThresholdMultiplier  = 10
)
