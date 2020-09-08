package service

import (
	. "HotelAutomation/model"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestShouldReturnApplianceInfo(t *testing.T) {
	request := CreateHotelRequest{NumberOfFloors: 2, MainCorridorPerFloor: 1, SubCorridorPerFloor: 2}
	hotelService := NewHotelService()
	hotelService.CreateHotel(request)

	expectedAppliances := []AppliancesInfo{
		{Name: "Light", Number: 1, IsSwitchedOd: false, PowerConsumption: 5, Location: CorridorLocation{FloorNumber: 1, CorridorType: MAIN, CorridorNumber: 1}},
		{Name: "Light", Number: 1, IsSwitchedOd: false, PowerConsumption: 5, Location: CorridorLocation{FloorNumber: 2, CorridorType: MAIN, CorridorNumber: 1}},
		{Name: "AC", Number: 1, IsSwitchedOd: false, PowerConsumption: 10, Location: CorridorLocation{FloorNumber: 1, CorridorType: MAIN, CorridorNumber: 1}},
		{Name: "AC", Number: 1, IsSwitchedOd: false, PowerConsumption: 10, Location: CorridorLocation{FloorNumber: 2, CorridorType: MAIN, CorridorNumber: 1}},
		{Name: "Light", Number: 1, IsSwitchedOd: false, PowerConsumption: 5, Location: CorridorLocation{FloorNumber: 1, CorridorType: SUB, CorridorNumber: 1}},
		{Name: "Light", Number: 2, IsSwitchedOd: false, PowerConsumption: 5, Location: CorridorLocation{FloorNumber: 1, CorridorType: SUB, CorridorNumber: 2}},
		{Name: "Light", Number: 1, IsSwitchedOd: false, PowerConsumption: 5, Location: CorridorLocation{FloorNumber: 2, CorridorType: SUB, CorridorNumber: 1}},
		{Name: "Light", Number: 2, IsSwitchedOd: false, PowerConsumption: 5, Location: CorridorLocation{FloorNumber: 2, CorridorType: SUB, CorridorNumber: 2}},
		{Name: "AC", Number: 1, IsSwitchedOd: false, PowerConsumption: 10, Location: CorridorLocation{FloorNumber: 1, CorridorType: SUB, CorridorNumber: 1}},
		{Name: "AC", Number: 2, IsSwitchedOd: false, PowerConsumption: 10, Location: CorridorLocation{FloorNumber: 1, CorridorType: SUB, CorridorNumber: 2}},
		{Name: "AC", Number: 1, IsSwitchedOd: false, PowerConsumption: 10, Location: CorridorLocation{FloorNumber: 2, CorridorType: SUB, CorridorNumber: 1}},
		{Name: "AC", Number: 2, IsSwitchedOd: false, PowerConsumption: 10, Location: CorridorLocation{FloorNumber: 2, CorridorType: SUB, CorridorNumber: 2}},
	}

	appliances := hotelService.GetAppliancesInfo()
	assert.Len(t, appliances, 12)
	assert.ElementsMatch(t, expectedAppliances, appliances)
}

func TestShouldGetNumberOfCorridors(t *testing.T) {
	request := CreateHotelRequest{NumberOfFloors: 2, MainCorridorPerFloor: 1, SubCorridorPerFloor: 2}
	hotelService := NewHotelService()
	hotelService.CreateHotel(request)

	assert.Equal(t, 1, hotelService.GetNumberOfCorridors(1, MAIN))
	assert.Equal(t, 2, hotelService.GetNumberOfCorridors(1, SUB))
}
