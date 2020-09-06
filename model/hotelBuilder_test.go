package model

import (
	"github.com/magiconair/properties/assert"
	"testing"
)

func TestShouldBuildHotelNoFloors(t *testing.T) {
	createHotelRequest := CreateHotelRequest{
		NumberOfFloors: 0,
	}
	hotel := NewHotelBuilder().
		WithFloorsAndCorridors(createHotelRequest).
		Build()
	floors := hotel.getFloors()
	assert.Equal(t, 0, len(floors))
}

func TestShouldBuildHotelWith2FloorsAnd1Main2SubCorridorAtEachFloor(t *testing.T) {
	createHotelRequest := CreateHotelRequest{
		NumberOfFloors:       2,
		MainCorridorPerFloor: 1,
		SubCorridorPerFloor:  2,
	}
	hotel := NewHotelBuilder().
		WithFloorsAndCorridors(createHotelRequest).
		Build()

	assert.Equal(t, 2, len(hotel.getFloors()))
	assert.Equal(t, 6, len(hotel.getCorridors()))
	for i, f := range hotel.getFloors() {
		assert.Equal(t, i+1, f.GetLevel())
		assert.Equal(t, 1, len(f.getCorridors(MAIN)))
		assert.Equal(t, 2, len(f.getCorridors(SUB)))
	}
}


func TestShouldBuildHotelWithoutLightBulbAndACInAnyCorridor(t *testing.T) {
	createHotelRequest := CreateHotelRequest{NumberOfFloors: 2, MainCorridorPerFloor: 1, SubCorridorPerFloor: 2}
	hotel := NewHotelBuilder().
		WithFloorsAndCorridors(createHotelRequest).
		WithOneLightBulbAndOneACInEveryCorridor().
		Build()

	assert.Equal(t, 6, len(hotel.getCorridors()))
	for _, c := range hotel.getCorridors() {
		assert.Equal(t, 1, len(c.getLightBulbs()))
		assert.Equal(t, 1, len(c.getAirConditioners()))
	}

}

func TestShouldBuildHotelWithLightBulbAndACInEachCorridor(t *testing.T) {
	createHotelRequest := CreateHotelRequest{NumberOfFloors: 2, MainCorridorPerFloor: 1, SubCorridorPerFloor: 2}
	hotel := NewHotelBuilder().
		WithFloorsAndCorridors(createHotelRequest).
		WithOneLightBulbAndOneACInEveryCorridor().
		Build()

	assert.Equal(t, 6, len(hotel.getCorridors()))
	for _, c := range hotel.getCorridors() {
		assert.Equal(t, 1, len(c.getLightBulbs()))
		assert.Equal(t, 1, len(c.getAirConditioners()))
	}

}
