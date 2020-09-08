package runner

import (
	. "HotelAutomation/IO"
	. "HotelAutomation/controller"
	. "HotelAutomation/service"
	"fmt"
)

type Application struct {
}

func NewApplication() *Application {
	return &Application{}
}

func (*Application) Run(reader Reader) {
	hotelConfig := getHotelConfig()
	motionController, hotelService := setupHotelWithPowerAndMotionControllers(hotelConfig)

	fmt.Println("Default state")
	hotelService.PrintHotelApplianceState()
	input := make(chan string)

	go takeInput(input, reader)
	go raiseMotionDetectedEvent(hotelService, motionController, input)

	select {
	//Run Infinitely
	}
}

func takeInput(input chan string, reader Reader) {
	for {
		input <- reader.Read()
	}
}

func raiseMotionDetectedEvent(hotelService *HotelService, controller *MotionController, input chan string) {
	for {
		select {
		case inputString := <-input:
			movementDetectedEvent, err := NewInputParser().ParseMovementInput(inputString)
			if err != nil {
				fmt.Println(err.Error())
				continue
			}
			controller.RaiseMotionDetectedEvent(movementDetectedEvent)
			hotelService.PrintHotelApplianceState()
		}
	}
}

func getHotelConfig() CreateHotelRequest {
	var floor, mainCorridor, subCorridor int
	fmt.Print("Number of floors : ")
	fmt.Scanf("%d", &floor)
	fmt.Print("Main corridors per floor: ")
	fmt.Scanf("%d", &mainCorridor)
	fmt.Print("Sub corridors per floor: ")
	fmt.Scanf("%d", &subCorridor)
	return CreateHotelRequest{NumberOfFloors: floor, MainCorridorPerFloor: mainCorridor, SubCorridorPerFloor: subCorridor}
}

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

