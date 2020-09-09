package parser

import (
	. "HotelAutomation/model"
	. "HotelAutomation/service"
	"errors"
	"regexp"
	"strconv"
	"strings"
)

type InputParser struct {
}

func NewInputParser() *InputParser {
	return &InputParser{}
}

func(*InputParser) ParseMovementInput(inputString string) (MovementDetectedEvent, error) {
	movementDetectedEvent := MovementDetectedEvent{}
	if !validInputString(inputString) {
		return movementDetectedEvent, errors.New("IncorrectInputFormat")
	}
	return mapInputToMotionDetectedEvent(inputString), nil
}

func mapInputToMotionDetectedEvent(inputString string) MovementDetectedEvent {
	floorNumber, corridorNumber := findFloorAndCorridorNumber(inputString)
	return MovementDetectedEvent{
		Movement: findMovement(inputString),
		Location: CorridorLocation{
			FloorNumber:    floorNumber,
			CorridorType:   findCorridorType(inputString),
			CorridorNumber: corridorNumber,
		},
	}
}

func findCorridorType(inputString string) CorridorType {
	var corridorType CorridorType
	if strings.Contains(inputString, "Main") {
		corridorType = MAIN
	}
	if strings.Contains(inputString, "Sub") {
		corridorType = SUB
	}
	return corridorType
}

func findFloorAndCorridorNumber(inputString string) (int, int) {
	re := regexp.MustCompile("[0-9]+")
	floorNumberAndCorridorNumber := re.FindAllString(inputString, -1)
	floorNumber, _ := strconv.Atoi(floorNumberAndCorridorNumber[0])
	corridorNumber, _ := strconv.Atoi(floorNumberAndCorridorNumber[1])
	return floorNumber, corridorNumber
}

func findMovement(inputString string) bool {
	if strings.HasPrefix(inputString, "Movement") {
		return true
	}
	return false
}

func validInputString(inputString string) bool {
	movementRegex, _ := regexp.Compile("Movement.*Floor\\s\\d,\\sSub corridor\\s\\d")
	noMovementRegex, _ := regexp.Compile("No\\smovement.*Floor\\s\\d,\\sSub corridor\\s\\d.for\\sa\\sminute")
	return movementRegex.MatchString(inputString) || noMovementRegex.MatchString(inputString)
}

//Movement in Floor 1, Sub corridor 2
//No movement in Floor 1, Sub corridor 2 for a minute
//No movement in Floor 1, Sub corridor 1 for a minute
