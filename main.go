package main

import (
	. "HotelAutomation/controller"
	. "HotelAutomation/model"
	. "HotelAutomation/service"
	"fmt"
)

func main() {
	hotelConfig := CreateHotelRequest{NumberOfFloors: 2, MainCorridorPerFloor: 1, SubCorridorPerFloor: 2}
	motionController, hotelService := SetupHotelWithPowerAndMotionControllers(hotelConfig)

	hotelService.PrintHotelApplianceState()
	motionController.RaiseMotionDetectedEvent(MovementDetectedEvent{
		Movement: true,
		Location: CorridorLocation{FloorNumber: 1, CorridorType: SUB, CorridorNumber: 2},
	})
	fmt.Println("-------------------------")
	fmt.Println("After Input 1")
	hotelService.PrintHotelApplianceState()
	motionController.RaiseMotionDetectedEvent(MovementDetectedEvent{
		Movement: true,
		Location: CorridorLocation{FloorNumber: 1, CorridorType: SUB, CorridorNumber: 1},
	})
	fmt.Println("-------------------------")
	fmt.Println("After Input 2")
	hotelService.PrintHotelApplianceState()
}

func SetupHotelWithPowerAndMotionControllers(createHotelReq CreateHotelRequest) (*MotionController, *HotelService) {
	hotelService := NewHotelService()
	powerControllerService := NewPowerControllerService(hotelService)
	powerAutomationController := NewPowerAutomationController(hotelService, powerControllerService)
	motionController := NewMotionController()
	motionController.AddSubscriber(powerAutomationController)

	hotelService.CreateHotel(createHotelReq)
	powerAutomationController.Init()

	return motionController, hotelService
}

type LocationApplianceFormat struct {
	number    string
	corridors Corridor
}
type CorridorFormat struct {
	cType  string
	number string
	lights []LightFormat
	ac     ACFormat
}
type ACFormat struct {
	name  string
	state string
}
type LightFormat struct {
	name   string
	number string
	state  string
}
