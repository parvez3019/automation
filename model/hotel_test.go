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
