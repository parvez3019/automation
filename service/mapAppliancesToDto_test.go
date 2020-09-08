package service

import (
	. "HotelAutomation/model"
	. "HotelAutomation/model/appliances"
	"github.com/stretchr/testify/assert"
	"testing"
)

var lightBulb *LightBulb
var airConditioner *AirConditioner
var mainCorridor *Corridor
var subCorridor *Corridor

func init() {
	lightBulb = NewLightBulb(1, 5)
	airConditioner = NewAirConditioner(1, 10)
	mainCorridor = NewCorridor(MAIN, 1).
		AddLightBulb(lightBulb).
		AddAirConditioner(airConditioner)
	subCorridor = NewCorridor(SUB, 1).
		AddLightBulb(lightBulb).
		AddAirConditioner(airConditioner)
}

func TestShouldReturnEmptyListOfApplianceInCaseOfNoFloors(t *testing.T) {
	assert.Equal(t, []Appliances{}, mapToAppliances([]*Floor{}))
}

func TestShouldReturnListOfApplianceFromMainCorridorOnlyInCaseOfNoSubCorridors(t *testing.T) {
	floor := NewFloor(1).
		AddCorridors([]*Corridor{mainCorridor})

	expectedApplianceInfo := []AppliancesInfo{
		{
			Name: "Light", Number: 1, IsSwitchedOn: false, PowerConsumption: 5,
			Location: CorridorLocation{FloorNumber: 1, CorridorType: "Main", CorridorNumber: 1},
		},
		{
			Name: "AC", Number: 1, IsSwitchedOn: false, PowerConsumption: 10,
			Location: CorridorLocation{FloorNumber: 1, CorridorType: "Main", CorridorNumber: 1},
		},
	}

	appliances := mapToAppliances([]*Floor{floor})
	assert.ElementsMatch(t, expectedApplianceInfo, mapApplianceToApplianceInfo(appliances))
}

func TestShouldReturnListOfApplianceFromSubCorridorOnlyInCaseOfNoMainCorridors(t *testing.T) {
	floor := NewFloor(1).
		AddCorridors([]*Corridor{subCorridor})

	expectedApplianceInfo := []AppliancesInfo{
		{
			Name: "Light", Number: 1, IsSwitchedOn: false, PowerConsumption: 5,
			Location: CorridorLocation{FloorNumber: 1, CorridorType: "Sub", CorridorNumber: 1},
		},
		{
			Name: "AC", Number: 1, IsSwitchedOn: false, PowerConsumption: 10,
			Location: CorridorLocation{FloorNumber: 1, CorridorType: "Sub", CorridorNumber: 1},
		},
	}

	appliances := mapToAppliances([]*Floor{floor})
	assert.ElementsMatch(t, expectedApplianceInfo, mapApplianceToApplianceInfo(appliances))
}

func TestShouldReturnListOfApplianceFromAllCorridor(t *testing.T) {
	floor := NewFloor(1).
		AddCorridors([]*Corridor{mainCorridor}).
		AddCorridors([]*Corridor{subCorridor})

	expectedApplianceInfo := []AppliancesInfo{
		{
			Name: "Light", Number: 1, IsSwitchedOn: false, PowerConsumption: 5,
			Location: CorridorLocation{FloorNumber: 1, CorridorType: "Main", CorridorNumber: 1},
		},
		{
			Name: "AC", Number: 1, IsSwitchedOn: false, PowerConsumption: 10,
			Location: CorridorLocation{FloorNumber: 1, CorridorType: "Main", CorridorNumber: 1},
		},
		{
			Name: "Light", Number: 1, IsSwitchedOn: false, PowerConsumption: 5,
			Location: CorridorLocation{FloorNumber: 1, CorridorType: "Sub", CorridorNumber: 1},
		},
		{
			Name: "AC", Number: 1, IsSwitchedOn: false, PowerConsumption: 10,
			Location: CorridorLocation{FloorNumber: 1, CorridorType: "Sub", CorridorNumber: 1},
		},
	}

	appliances := mapToAppliances([]*Floor{floor})
	assert.ElementsMatch(t, expectedApplianceInfo, mapApplianceToApplianceInfo(appliances))
}
