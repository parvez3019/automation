package service

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestShouldReturnApplianceInfo(t *testing.T) {
	request := CreateHotelRequest{NumberOfFloors: 2, MainCorridorPerFloor: 1, SubCorridorPerFloor: 2}
	hotelService := NewHotelService()
	hotelService.CreateHotel(request)

	expectedAppliances := []AppliancesInfo{
		{Name: "Light", Number: 1, IsSwitchedOd: false, PowerConsumption: 5, Location: CorridorLocation{FloorNumber: 1, CorridorType: "Main", CorridorNumber: 1}},
		{Name: "Light", Number: 1, IsSwitchedOd: false, PowerConsumption: 5, Location: CorridorLocation{FloorNumber: 1, CorridorType: "Sub", CorridorNumber: 1}},
		{Name: "Light", Number: 1, IsSwitchedOd: false, PowerConsumption: 5, Location: CorridorLocation{FloorNumber: 1, CorridorType: "Sub", CorridorNumber: 2}},
		{Name: "Light", Number: 1, IsSwitchedOd: false, PowerConsumption: 5, Location: CorridorLocation{FloorNumber: 2, CorridorType: "Main", CorridorNumber: 1}},
		{Name: "Light", Number: 1, IsSwitchedOd: false, PowerConsumption: 5, Location: CorridorLocation{FloorNumber: 2, CorridorType: "Sub", CorridorNumber: 1}},
		{Name: "Light", Number: 1, IsSwitchedOd: false, PowerConsumption: 5, Location: CorridorLocation{FloorNumber: 2, CorridorType: "Sub", CorridorNumber: 2}},
		{Name: "AC", Number: 1, IsSwitchedOd: false, PowerConsumption: 10, Location: CorridorLocation{FloorNumber: 1, CorridorType: "Main", CorridorNumber: 1}},
		{Name: "AC", Number: 1, IsSwitchedOd: false, PowerConsumption: 10, Location: CorridorLocation{FloorNumber: 1, CorridorType: "Sub", CorridorNumber: 1}},
		{Name: "AC", Number: 1, IsSwitchedOd: false, PowerConsumption: 10, Location: CorridorLocation{FloorNumber: 1, CorridorType: "Sub", CorridorNumber: 2}},
		{Name: "AC", Number: 1, IsSwitchedOd: false, PowerConsumption: 10, Location: CorridorLocation{FloorNumber: 2, CorridorType: "Main", CorridorNumber: 1}},
		{Name: "AC", Number: 1, IsSwitchedOd: false, PowerConsumption: 10, Location: CorridorLocation{FloorNumber: 2, CorridorType: "Sub", CorridorNumber: 1}},
		{Name: "AC", Number: 1, IsSwitchedOd: false, PowerConsumption: 10, Location: CorridorLocation{FloorNumber: 2, CorridorType: "Sub", CorridorNumber: 2}},
	}

	assert.ElementsMatch(t, expectedAppliances, hotelService.GetAppliancesInfo())
}

func TestShouldGetNumberOfCorridors(t *testing.T) {
	request := CreateHotelRequest{NumberOfFloors: 2, MainCorridorPerFloor: 1, SubCorridorPerFloor: 2}
	hotelService := NewHotelService()
	hotelService.CreateHotel(request)

	assert.Equal(t, 1, hotelService.GetNumberOfCorridors(1, "Main"))
	assert.Equal(t, 2, hotelService.GetNumberOfCorridors(1, "Sub"))
}
