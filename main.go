package main

import (
	. "HotelAutomation/controller"
	. "HotelAutomation/service"
)

func main() {
	hotelConfig := CreateHotelRequest{NumberOfFloors: 2, MainCorridorPerFloor: 1, SubCorridorPerFloor: 2}
	motionController := SetupHotelWithPowerAndMotionControllers(hotelConfig)

	motionController.RaiseMotionDetectedEvent(MovementDetectedEvent{
		Movement: true,
		Location: CorridorLocation{FloorNumber: 1, CorridorType: "Main", CorridorNumber: 1},
	})
}

func SetupHotelWithPowerAndMotionControllers(createHotelReq CreateHotelRequest) *MotionController {
	hotelService := NewHotelService()
	powerControllerService := NewPowerControllerService(hotelService)
	powerAutomationController := NewPowerAutomationController(hotelService, powerControllerService)
	motionController := NewMotionController()
	motionController.AddSubscriber(powerAutomationController)

	hotelService.CreateHotel(createHotelReq)
	powerAutomationController.Init()

	return motionController
}
