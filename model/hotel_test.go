package model

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestShouldReturnHotelWithNoFloors(t *testing.T) {
	assert.Equal(t,0, len(NewHotel().getFloors()))
}

func TestShouldTwoAddedFloorsAddedToHotel(t *testing.T) {
	hotel := NewHotel().addFloors([]*Floor{
		NewFloor(1),
		NewFloor(1),
	})
	assert.Equal(t, 2, len(hotel.getFloors()))
}
