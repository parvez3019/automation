package model

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestShouldCreateACorridorWithTypeMainAndZeroLightBulbsAndACs(t *testing.T) {
	corridor := NewCorridor(MAIN, 1)
	assert.Equal(t, "Main", corridor.GetType())
	assert.Equal(t, 0, len(corridor.GetLightBulbs()))
	assert.Equal(t, 0, len(corridor.GetAirConditioners()))
}

func TestShouldCreateACorridorWithTypeSubAndZeroLightBulbsAndACs(t *testing.T) {
	corridor := NewCorridor(SUB, 1)
	assert.Equal(t, "Sub", corridor.GetType())
	assert.Equal(t, 0, len(corridor.GetLightBulbs()))
	assert.Equal(t, 0, len(corridor.GetAirConditioners()))
}
