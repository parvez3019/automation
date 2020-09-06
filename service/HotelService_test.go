package service

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestShouldReturnApplianceInfo(t *testing.T) {
	request := CreateHotelRequest{NumberOfFloors: 2, MainCorridorPerFloor: 1, SubCorridorPerFloor: 2}
	hotelService := NewHotelService()
	hotelService.CreateHotel(request)

	expectedAppliances := []ApplianceInfo{
		{Name: "Light", Number: 1, IsSwitchedOd: false, PowerConsumption: 5, Location: ApplianceLocation{floorNumber: 1, corridorType: "Main", corridorNumber: 1}},
		{Name: "Light", Number: 1, IsSwitchedOd: false, PowerConsumption: 5, Location: ApplianceLocation{floorNumber: 1, corridorType: "Sub", corridorNumber: 1}},
		{Name: "Light", Number: 1, IsSwitchedOd: false, PowerConsumption: 5, Location: ApplianceLocation{floorNumber: 1, corridorType: "Sub", corridorNumber: 2}},
		{Name: "Light", Number: 1, IsSwitchedOd: false, PowerConsumption: 5, Location: ApplianceLocation{floorNumber: 2, corridorType: "Main", corridorNumber: 1}},
		{Name: "Light", Number: 1, IsSwitchedOd: false, PowerConsumption: 5, Location: ApplianceLocation{floorNumber: 2, corridorType: "Sub", corridorNumber: 1}},
		{Name: "Light", Number: 1, IsSwitchedOd: false, PowerConsumption: 5, Location: ApplianceLocation{floorNumber: 2, corridorType: "Sub", corridorNumber: 2}},
		{Name: "AC", Number: 1, IsSwitchedOd: false, PowerConsumption: 10, Location: ApplianceLocation{floorNumber: 1, corridorType: "Main", corridorNumber: 1}},
		{Name: "AC", Number: 1, IsSwitchedOd: false, PowerConsumption: 10, Location: ApplianceLocation{floorNumber: 1, corridorType: "Sub", corridorNumber: 1}},
		{Name: "AC", Number: 1, IsSwitchedOd: false, PowerConsumption: 10, Location: ApplianceLocation{floorNumber: 1, corridorType: "Sub", corridorNumber: 2}},
		{Name: "AC", Number: 1, IsSwitchedOd: false, PowerConsumption: 10, Location: ApplianceLocation{floorNumber: 2, corridorType: "Main", corridorNumber: 1}},
		{Name: "AC", Number: 1, IsSwitchedOd: false, PowerConsumption: 10, Location: ApplianceLocation{floorNumber: 2, corridorType: "Sub", corridorNumber: 1}},
		{Name: "AC", Number: 1, IsSwitchedOd: false, PowerConsumption: 10, Location: ApplianceLocation{floorNumber: 2, corridorType: "Sub", corridorNumber: 2}},
	}

	assert.ElementsMatch(t, expectedAppliances, hotelService.GetAppliances())
}
