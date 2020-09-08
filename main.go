package main

import (
	. "HotelAutomation/controller"
	. "HotelAutomation/model"
	. "HotelAutomation/service"
	. "bufio"
	"errors"
	"fmt"
	. "os"
	"regexp"
	"strconv"
	"strings"
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

	go takeInput(input)
	go raiseMotionDetectedEvent(hotelService, motionController, input)

	select {}
}

func takeInput(input chan string) {
	for {
		reader := NewReader(Stdin)
		inputStringFromConsole, _ := reader.ReadString('\n')
		input <- inputStringFromConsole
	}

}

func raiseMotionDetectedEvent(hotelService *HotelService, controller *MotionController, input chan string) {
	for {
		select {
		case inputString := <-input:
			movementDetectedEvent, err := parseInputString(inputString)
			if err != nil {
				fmt.Println(err.Error())
				continue
			}
			fmt.Println(movementDetectedEvent)
			controller.RaiseMotionDetectedEvent(movementDetectedEvent)
			hotelService.PrintHotelApplianceState()
		}
	}
}

func parseInputString(inputString string) (MovementDetectedEvent, error) {
	movementDetectedEvent := MovementDetectedEvent{}
	if !validInputString(inputString) {
		return movementDetectedEvent, errors.New("IncorrectInputFormat")
	}
	return mapInputToMotionDetectedEvent(inputString), nil
}

func mapInputToMotionDetectedEvent(inputString string) MovementDetectedEvent {
	var corridorType CorridorType
	isMovementDetected := false
	if strings.HasPrefix(inputString, "No movement") && strings.HasSuffix(inputString, "for a minute") {
		isMovementDetected = false
	}
	if strings.HasPrefix(inputString, "Movement") {
		isMovementDetected = true
	}

	re := regexp.MustCompile("[0-9]+")
	floorNumberAndCorridorNumber := re.FindAllString(inputString, -1)

	if strings.Contains(inputString, "Main") {
		corridorType = MAIN
	}
	if strings.Contains(inputString, "Sub") {
		corridorType = SUB
	}

	floorNumber, _ := strconv.Atoi(floorNumberAndCorridorNumber[0])
	corridorNumber, _ := strconv.Atoi(floorNumberAndCorridorNumber[1])

	return MovementDetectedEvent{
		Movement: isMovementDetected,
		Location: CorridorLocation{
			FloorNumber:    floorNumber,
			CorridorType:   corridorType,
			CorridorNumber: corridorNumber,
		},
	}
}

func validInputString(inputString string) bool {
	movementRegex, _ := regexp.Compile("Movement.*Floor\\s\\d,\\sSub corridor\\s\\d")
	noMovementRegex, _ := regexp.Compile("No\\smovement.*Floor\\s\\d,\\sSub corridor\\s\\d.for\\sa\\sminute")
	return movementRegex.MatchString(inputString) || noMovementRegex.MatchString(inputString)
}

//Movement in Floor 1, Sub corridor 2
//No movement in Floor 1, Sub corridor 2 for a minute
//No movement in Floor 1, Sub corridor 1 for a minute