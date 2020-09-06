package service

import (
	. "HotelAutomation/model"
	. "HotelAutomation/model/appliances"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestShouldReturnEmptyListOfApplianceInCaseOfNoFloors(t *testing.T) {
	assert.Equal(t, []ApplianceInfo{}, mapToApplianceStateDto([]*Floor{}))
}

func TestShouldReturnListOfApplianceFromMainCorridorOnlyInCaseOfNoSubCorridors(t *testing.T) {
	mainCorridor := NewCorridor(MAIN, 1).
		AddLightBulb(NewLightBulb(1, 5)).
		AddAirConditioner(NewAirConditioner(1, 10))
	floor := NewFloor(1).
		AddCorridors([]*Corridor{mainCorridor}, MAIN)

	expectedApplianceInfo := []ApplianceInfo{
		{
			Name: "Light", Number: 1, IsSwitchedOd: false, PowerConsumption: 5,
			Location: ApplianceLocation{floorNumber: 1, corridorType: "Main", corridorNumber: 1},
		},
		{
			Name: "AC", Number: 1, IsSwitchedOd: false, PowerConsumption: 10,
			Location: ApplianceLocation{floorNumber: 1, corridorType: "Main", corridorNumber: 1},
		},
	}

	applianceInfos := mapToApplianceStateDto([]*Floor{floor})
	assert.ElementsMatch(t, expectedApplianceInfo, applianceInfos)
}

func TestShouldReturnListOfApplianceFromSubCorridorOnlyInCaseOfNoMainCorridors(t *testing.T) {
	subCorridor := NewCorridor(SUB, 1).
		AddLightBulb(NewLightBulb(1, 5)).
		AddAirConditioner(NewAirConditioner(1, 10))
	floor := NewFloor(1).
		AddCorridors([]*Corridor{subCorridor}, SUB)

	expectedApplianceInfo := []ApplianceInfo{
		{
			Name: "Light", Number: 1, IsSwitchedOd: false, PowerConsumption: 5,
			Location: ApplianceLocation{floorNumber: 1, corridorType: "Sub", corridorNumber: 1},
		},
		{
			Name: "AC", Number: 1, IsSwitchedOd: false, PowerConsumption: 10,
			Location: ApplianceLocation{floorNumber: 1, corridorType: "Sub", corridorNumber: 1},
		},
	}

	applianceInfos := mapToApplianceStateDto([]*Floor{floor})
	assert.ElementsMatch(t, expectedApplianceInfo, applianceInfos)
}

func TestShouldReturnListOfApplianceFromAllCorridor(t *testing.T) {
	mainCorridor := NewCorridor(MAIN, 1).
		AddLightBulb(NewLightBulb(1, 5)).
		AddAirConditioner(NewAirConditioner(1, 10))
	subCorridor := NewCorridor(SUB, 1).
		AddLightBulb(NewLightBulb(1, 5)).
		AddAirConditioner(NewAirConditioner(1, 10))
	floor := NewFloor(1).
		AddCorridors([]*Corridor{mainCorridor}, SUB).
		AddCorridors([]*Corridor{subCorridor}, MAIN)

	expectedApplianceInfo := []ApplianceInfo{
		{
			Name: "Light", Number: 1, IsSwitchedOd: false, PowerConsumption: 5,
			Location: ApplianceLocation{floorNumber: 1, corridorType: "Main", corridorNumber: 1},
		},
		{
			Name: "AC", Number: 1, IsSwitchedOd: false, PowerConsumption: 10,
			Location: ApplianceLocation{floorNumber: 1, corridorType: "Main", corridorNumber: 1},
		},
		{
			Name: "Light", Number: 1, IsSwitchedOd: false, PowerConsumption: 5,
			Location: ApplianceLocation{floorNumber: 1, corridorType: "Sub", corridorNumber: 1},
		},
		{
			Name: "AC", Number: 1, IsSwitchedOd: false, PowerConsumption: 10,
			Location: ApplianceLocation{floorNumber: 1, corridorType: "Sub", corridorNumber: 1},
		},
	}

	applianceInfos := mapToApplianceStateDto([]*Floor{floor})
	assert.ElementsMatch(t, expectedApplianceInfo, applianceInfos)
}
