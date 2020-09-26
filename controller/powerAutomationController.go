package controller

import (
	. "HotelAutomation/model"
	. "HotelAutomation/model/appliances"
	. "HotelAutomation/service"
	"context"
	"fmt"
	"time"
)

type Subscriber interface {
	Update(MovementDetectedEvent) error
}

type PowerAutomationController struct {
	hotelService                 HotelServiceI
	powerController              PowerControllerServiceI
	noMotionTimeout              time.Duration
	ctx                          context.Context
	addMotionDetectedChannel     chan CorridorLocation
	motionDetectedTimeoutChannel chan CorridorLocation
	motionEnabledAppliance       map[CorridorLocation]int
	changeInApplianceState       chan bool
}

func NewPowerAutomationController(
	hService HotelServiceI,
	pController PowerControllerServiceI,

) *PowerAutomationController {
	return &PowerAutomationController{
		hotelService:                 hService,
		powerController:              pController,
		motionEnabledAppliance:       make(map[CorridorLocation]int),
		ctx:                          context.Background(),
		addMotionDetectedChannel:     make(chan CorridorLocation),
		motionDetectedTimeoutChannel: make(chan CorridorLocation),
		changeInApplianceState:       make(chan bool),
	}
}

func (c *PowerAutomationController) Init(timeout time.Duration, changeInApplianceState chan bool) {
	c.powerController.RegisterDevices()
	c.noMotionTimeout = timeout
	c.changeInApplianceState = changeInApplianceState
}

func (c *PowerAutomationController) Update(request MovementDetectedEvent) error {
	err := c.toggleLight(request.Movement, request.Location)
	if err != nil {
		return err
	}
	go c.lookup()
	go c.addToMotionEnabledAppliance(request.Location)
	return nil
}

func (c *PowerAutomationController) toggleLight(toggle bool, location CorridorLocation) error {
	toggleLightBulbRequest := ToggleApplianceRequest{AppType: LIGHT, TurnOn: toggle, Location: location}
	err := c.powerController.Update(toggleLightBulbRequest)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	c.verifyAndToggleACBasedOnTotalPowerConsumption(location)
	return nil
}

func (c *PowerAutomationController) verifyAndToggleACBasedOnTotalPowerConsumption(atLocation CorridorLocation) {
	c.toggleSubCorridorAC(atLocation.FloorNumber, true)
	if c.totalPowerConsumptionAtFloorExceeded(atLocation.FloorNumber) {
		c.toggleSubCorridorAC(atLocation.FloorNumber, false)
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

func (c *PowerAutomationController) addToMotionEnabledAppliance(location CorridorLocation) {
	c.addMotionDetectedChannel <- location
	cxt, _ := context.WithTimeout(c.ctx, c.noMotionTimeout)
	select {
	case <-cxt.Done():
		c.motionDetectedTimeoutChannel <- location
	}
}

func (c *PowerAutomationController) lookup() {
	for {
		select {
		case location := <-c.addMotionDetectedChannel:
			c.motionEnabledAppliance[location] = c.motionEnabledAppliance[location] + 1
		case location := <-c.motionDetectedTimeoutChannel:
			c.motionEnabledAppliance[location] = c.motionEnabledAppliance[location] - 1
			c.turnOffLight(location)
		}
	}
}

func (c *PowerAutomationController) turnOffLight(location CorridorLocation) {
	if c.motionEnabledAppliance[location] == 0 {
		_ = c.toggleLight(false, location)
		c.changeInApplianceState <- true
	}
}

const (
	MainCorridorPowerConsumptionThresholdMultiplier = 15
	SubCorridorPowerConsumptionThresholdMultiplier  = 10
)
