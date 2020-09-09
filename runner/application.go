package runner

import (
	. "HotelAutomation/parser"
	. "HotelAutomation/controller"
	. "HotelAutomation/service"
	"fmt"
)

type Application struct {
	reader Reader
}

func NewApplication(reader Reader) *Application {
	return &Application{reader: reader}
}

func (app *Application) Run() {
	hotelConfig := getHotelConfig()
	motionController, hotelService := setupHotelWithPowerAndMotionControllers(hotelConfig)
	f := NewFormatter()

	fmt.Println("Default state")
	fmt.Println(f.ApplianceInfoToString(hotelService.GetAppliancesInfo()))
	input := make(chan string)

	go takeInput(input, app.reader)
	go raiseMotionDetectedEvent(hotelService, motionController, input, f)

	select {
	//Run Infinitely
	}
}

func takeInput(input chan string, reader Reader) {
	for {
		input <- reader.ReadLine()
	}
}

func raiseMotionDetectedEvent(hotelService *HotelService, controller *MotionController, input chan string, f *Formatter) {
	for {
		select {
		case inputString := <-input:
			movementDetectedEvent, err := NewInputParser().ParseMovementInput(inputString)
			if err != nil {
				fmt.Println(err.Error())
				continue
			}
			controller.RaiseMotionDetectedEvent(movementDetectedEvent)
			printHotelState(f, hotelService)
		}
	}
}

func printHotelState(f *Formatter, hotelService *HotelService) {
	fmt.Println(f.ApplianceInfoToString(hotelService.GetAppliancesInfo()))
	fmt.Println("--------------------")
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

