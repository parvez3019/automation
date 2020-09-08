package main

import (
	. "HotelAutomation/IO"
	. "HotelAutomation/controller"
	. "HotelAutomation/service"
	"fmt"
)

func setupHotelWithPowerAndMotionControllers(createHotelReq CreateHotelRequest) (*MotionController, *HotelService) {
	hotelService := NewHotelService()
	powerControllerService := NewPowerControllerService(hotelService)
	powerAutomationController := NewPowerAutomationController(hotelService, powerControllerService)
	motionController := NewMotionController()
	motionController.AddSubscriber(powerAutomationController)

	hotelService.CreateHotel(createHotelReq)
	powerAutomationController.Init()

	return motionController, hotelService
}

func main() {
	hotelConfig := CreateHotelRequest{NumberOfFloors: 2, MainCorridorPerFloor: 1, SubCorridorPerFloor: 2}
	motionController, hotelService := setupHotelWithPowerAndMotionControllers(hotelConfig)

	fmt.Println("Initial State")
	hotelService.PrintHotelApplianceState()
	input := make(chan string)

	go takeInput(input, *NewConsoleReader())
	go raiseMotionDetectedEvent(hotelService, motionController, input)

	// To run infinitely
	select {}
}

func takeInput(input chan string, reader ConsoleReader) {
	for {
		input <- reader.Read()
	}
}

func raiseMotionDetectedEvent(hotelService *HotelService, controller *MotionController, input chan string) {
	for {
		select {
		case inputString := <-input:
			movementDetectedEvent, err := NewInputParser().Parse(inputString)
			if err != nil {
				fmt.Println(err.Error())
				continue
			}
			controller.RaiseMotionDetectedEvent(movementDetectedEvent)
			hotelService.PrintHotelApplianceState()
		}
	}
}
