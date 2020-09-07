package model

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestShouldBuildHotelNoFloors(t *testing.T) {
	hotel := NewHotelBuilder().
		Build()
	floors := hotel.GetFloors()
	assert.Equal(t, 0, len(floors))
}

func TestShouldBuildHotelWith2FloorsAnd1Main2SubCorridorAtEachFloor(t *testing.T) {
	hotel := NewHotelBuilder().
		WithFloors(2).
		WithCorridors(1, 2).
		Build()

	assert.Equal(t, 2, len(hotel.GetFloors()))
	assert.Equal(t, 6, len(hotel.GetAllCorridors()))
	for i, f := range hotel.GetFloors() {
		assert.Equal(t, i+1, f.GetLevel())
		assert.Equal(t, 1, len(f.GetCorridors(MAIN)))
		assert.Equal(t, 2, len(f.GetCorridors(SUB)))
	}
}

func TestShouldBuildHotelWithoutLightBulbAndACInAnyCorridor(t *testing.T) {
	hotel := NewHotelBuilder().
		WithFloors(2).
		WithCorridors(1, 2).
		WithOneLightBulbAndOneACInEveryCorridor().
		Build()

	assert.Equal(t, 6, len(hotel.GetAllCorridors()))
	for _, c := range hotel.GetAllCorridors() {
		assert.Equal(t, 1, len(c.GetLightBulbs()))
		assert.Equal(t, 1, len(c.GetAirConditioners()))
	}

}

func TestShouldBuildHotelWithLightBulbAndACInEachCorridor(t *testing.T) {
	hotel := NewHotelBuilder().
		WithFloors(2).
		WithCorridors(1, 2).
		WithOneLightBulbAndOneACInEveryCorridor().
		Build()

	assert.Equal(t, 6, len(hotel.GetAllCorridors()))
	for _, c := range hotel.GetAllCorridors() {
		assert.Equal(t, 1, len(c.GetLightBulbs()))
		assert.Equal(t, 1, len(c.GetAirConditioners()))
	}

}
