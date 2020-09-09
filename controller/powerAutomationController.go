package controller

import (
	. "HotelAutomation/model"
	. "HotelAutomation/model/appliances"
	. "HotelAutomation/service"
	"fmt"
)

type Subscriber interface {
	Update(MovementDetectedEvent) error
}

type PowerAutomationController struct {
	hotelService    HotelServiceI
	powerController PowerControllerServiceI
}

func NewPowerAutomationController(
	hService HotelServiceI,
	pController PowerControllerServiceI,
) *PowerAutomationController {
	return &PowerAutomationController{
		hotelService:    hService,
		powerController: pController,
	}
}

func (c *PowerAutomationController) Init() {
	c.powerController.RegisterDevices()
}

func (c *PowerAutomationController) Update(request MovementDetectedEvent) error{
	toggleLightBulbRequest := ToggleApplianceRequest{AppType: LIGHT, TurnOn: request.Movement, Location: request.Location}
	err := c.powerController.Update(toggleLightBulbRequest)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	c.verifyAndToggleACBasedOnTotalPowerConsumption(request.Location)
	return nil
}

func (c *PowerAutomationController) verifyAndToggleACBasedOnTotalPowerConsumption(atLocation CorridorLocation) {
	c.toggleSubCorridorAC(atLocation.FloorNumber, true)
	if c.totalPowerConsumptionAtFloorExceeded(atLocation.FloorNumber) {
		c.toggleSubCorridorAC(atLocation.FloorNumber, false)
		return
	}
}

func (c *PowerAutomationController) totalPowerConsumptionAtFloorExceeded(floorNumber int) bool {
	totalConsumption := c.powerController.TotalPowerConsumptionAtFloor(floorNumber)
	totalMainCorridors := c.hotelService.GetNumberOfCorridors(floorNumber, MAIN)
	totalSubCorridors := c.hotelService.GetNumberOfCorridors(floorNumber, SUB)

	powerThreshold := (totalMainCorridors * MainCorridorPowerConsumptionThresholdMultiplier) +
		(totalSubCorridors * SubCorridorPowerConsumptionThresholdMultiplier)
	return totalConsumption > powerThreshold
}

func (c *PowerAutomationController) toggleSubCorridorAC(floor int, switchOn bool) {
	c.powerController.ToggleApplianceToReverseState(floor, SUB, AC, switchOn)
}

const (
	MainCorridorPowerConsumptionThresholdMultiplier = 15
	SubCorridorPowerConsumptionThresholdMultiplier  = 10
)
