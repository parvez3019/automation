package model

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestShouldReturnHotelWithNoFloors(t *testing.T) {
	assert.Equal(t,0, len(NewHotel().getFloors()))
}

func TestShouldTwoAddedFloorsAddedToHotel(t *testing.T) {
	hotel := NewHotel().addFloor(NewFloor(1)).
		addFloor(NewFloor(1))
	assert.Equal(t, 2, len(hotel.getFloors()))
}
