package model

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestShouldReturnHotelWithNoFloors(t *testing.T) {
	assert.Equal(t, 0, len(NewHotel().GetFloors()))
}

func TestShouldTwoAddedFloorsAddedToHotel(t *testing.T) {
	hotel := NewHotel().AddFloors([]*Floor{
		NewFloor(1),
		NewFloor(1),
	})
	assert.Equal(t, 2, len(hotel.GetFloors()))
}

func TestShouldGetNumberOfCorridorsAtFloor(t *testing.T) {
	floor1 := NewFloor(1).AddCorridors([]*Corridor{
		NewCorridor(MAIN, 1), NewCorridor(SUB, 1), NewCorridor(SUB, 2),
	})
	floor2 := NewFloor(2).AddCorridors([]*Corridor{
		NewCorridor(MAIN, 1), NewCorridor(SUB, 1), NewCorridor(SUB, 2),
	})
	hotel := NewHotel().AddFloors([]*Floor{floor1, floor2})

	assert.Len(t, hotel.GetCorridorsAtFloor(1, MAIN), 1)
	assert.Len(t, hotel.GetCorridorsAtFloor(2, MAIN), 1)
	assert.Len(t, hotel.GetCorridorsAtFloor(1, SUB), 2)
	assert.Len(t, hotel.GetCorridorsAtFloor(2, SUB), 2)
	assert.Len(t, hotel.GetAllCorridors(), 6)
}
