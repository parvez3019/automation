package controller

import (
	. "HotelAutomation/model"
	. "HotelAutomation/model/appliances"
	. "HotelAutomation/service"
	"fmt"
)

type Subscriber interface {
	Update(MovementDetectedEvent)
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

func (c *PowerAutomationController) Update(request MovementDetectedEvent) {
	toggleLightBulbRequest := ToggleApplianceRequest{AppType: LIGHT, TurnOn: request.Movement, Location: request.Location}
	err := c.powerController.Update(toggleLightBulbRequest)
	if err != nil {
		fmt.Println(err.Error())
	}
	c.verifyAndToggleACBasedOnTotalPowerConsumption(request.Location)
}

func (c *PowerAutomationController) GetCurrentStateInfoOfApplianceWithLocation() []AppliancesInfo {
	return c.hotelService.GetAppliancesInfo()
}

func (c *PowerAutomationController) toggleAppliance(request ToggleApplianceRequest) error {
	return c.powerController.Update(request)
}

func (c *PowerAutomationController) verifyAndToggleACBasedOnTotalPowerConsumption(atLocation CorridorLocation) {
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

func (c *PowerAutomationController) toggleSubCorridorAC(atLocation CorridorLocation, switchOn bool) {
	_ = c.toggleAppliance(ToggleApplianceRequest{
		AppType:  AC,
		TurnOn:   switchOn,
		Location: atLocation,
	})
}

const (
	MainCorridorPowerConsumptionThresholdMultiplier = 15
	SubCorridorPowerConsumptionThresholdMultiplier  = 10
)
