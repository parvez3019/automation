package model

import (
	. "HotelAutomation/model/appliances"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestShouldCreateACorridorWithTypeMainAndZeroLightBulbsAndACs(t *testing.T) {
	corridor := NewCorridor(MAIN, 1)
	assert.Equal(t, MAIN, corridor.GetType())
	assert.Equal(t, 0, len(corridor.GetAppliances(LIGHT)))
	assert.Equal(t, 0, len(corridor.GetAppliances(AC)))
}

func TestShouldCreateACorridorWithTypeSubAndZeroLightBulbsAndACs(t *testing.T) {
	corridor := NewCorridor(SUB, 1)
	assert.Equal(t, SUB, corridor.GetType())
	assert.Equal(t, 0, len(corridor.GetAppliances(LIGHT)))
	assert.Equal(t, 0, len(corridor.GetAppliances(AC)))
}

func TestShouldAddLightAndACToCorridor(t *testing.T) {
	corridor := NewCorridor(SUB, 1).
		AddLightBulb(NewLightBulb(1, 5)).
		AddLightBulb(NewLightBulb(2, 5)).
		AddAirConditioner(NewAirConditioner(1, 10))

	assert.Equal(t, 2, len(corridor.GetAppliances(LIGHT)))
	assert.Equal(t, 1, len(corridor.GetAppliances(AC)))
}
