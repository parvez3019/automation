package model

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestShouldCreateACorridorWithTypeMainAndZeroLightBulbsAndACs(t *testing.T) {
	corridor := NewCorridor(MAIN, 1)
	assert.Equal(t, "Main", corridor.getTypeAsString())
	assert.Equal(t, 0, len(corridor.getLightBulbs()))
	assert.Equal(t, 0, len(corridor.getAirConditioners()))
}

func TestShouldCreateACorridorWithTypeSubAndZeroLightBulbsAndACs(t *testing.T) {
	corridor := NewCorridor(SUB, 1)
	assert.Equal(t, "Sub", corridor.getTypeAsString())
	assert.Equal(t, 0, len(corridor.getLightBulbs()))
	assert.Equal(t, 0, len(corridor.getAirConditioners()))
}

