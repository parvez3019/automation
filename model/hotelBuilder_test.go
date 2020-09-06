package model

import (
	"github.com/magiconair/properties/assert"
	"testing"
)

func TestShouldBuildHotelNoFloors(t *testing.T) {
	createHotelRequest := CreateHotelRequest{
		NumberOfFloors: 0,
	}
	hotel := NewHotelBuilder().BuildHotel(createHotelRequest)
	floors := hotel.getFloors()
	assert.Equal(t, 0, len(floors))
}

func TestShouldBuildHotelWith2FloorsAnd1Main2SubCorridorAtEachFloor(t *testing.T) {
	createHotelRequest := CreateHotelRequest{
		NumberOfFloors:       2,
		MainCorridorPerFloor: 1,
		SubCorridorPerFloor:  2,
	}
	hotel := NewHotelBuilder().BuildHotel(createHotelRequest)
	floors := hotel.getFloors()
	assert.Equal(t, 2, len(floors))
	firstFloor := floors[0]
	secondFloor := floors[1]
	assert.Equal(t, 1, firstFloor.GetLevel())
	assert.Equal(t, 2, secondFloor.GetLevel())
	assert.Equal(t, 1, len(firstFloor.getCorridors(MAIN)))
	assert.Equal(t, 1, len(secondFloor.getCorridors(MAIN)))
	assert.Equal(t, 2, len(firstFloor.getCorridors(SUB)))
	assert.Equal(t, 2, len(secondFloor.getCorridors(SUB)))
}
