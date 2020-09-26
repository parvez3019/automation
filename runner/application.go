package runner

import (
	. "HotelAutomation/parser"
	. "HotelAutomation/controller"
	. "HotelAutomation/service"
	"fmt"
	"time"
)

type Application struct {
	reader Reader
}

func NewApplication(reader Reader) *Application {
	return &Application{reader: reader}
}

func (app *Application) Run() {
	hotelConfig := getHotelConfig()
	motionController, hotelService, changeInApplianceState := setupHotelWithPowerAndMotionControllers(hotelConfig)
	f := NewFormatter()

	fmt.Println("Default state")
	fmt.Println(f.ApplianceInfoToString(hotelService.GetAppliancesInfo()))
	input := make(chan string)

	go takeInput(input, app.reader)
	go raiseMotionDetectedEvent(hotelService, motionController, input, f)
	go printCurrentStateOfHotel(hotelService, changeInApplianceState, f)

	select {
	//Run Infinitely
	}
}

func printCurrentStateOfHotel(service *HotelService, change chan bool, f *Formatter) {
	for {
		select {
		case <-change:
			fmt.Println(f.ApplianceInfoToString(service.GetAppliancesInfo()))
		}
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
			err = controller.RaiseMotionDetectedEvent(movementDetectedEvent)
			if err != nil {
				fmt.Println(err.Error())
			}
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

func setupHotelWithPowerAndMotionControllers(createHotelReq CreateHotelRequest) (*MotionController, *HotelService, chan bool) {
	hotelService := NewHotelService()
	powerControllerService := NewPowerControllerService(hotelService)
	powerAutomationController := NewPowerAutomationController(hotelService, powerControllerService)
	motionController := NewMotionController()
	motionController.AddSubscriber(powerAutomationController)

	hotelService.CreateHotel(createHotelReq)
	changeInApplianceState := make(chan bool)
	powerAutomationController.Init(10*time.Second, changeInApplianceState)

	return motionController, hotelService, changeInApplianceState
}
