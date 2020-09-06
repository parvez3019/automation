package model

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestShouldReturnNewFloorWithLevel1(t *testing.T) {
	assert.Equal(t, 1, NewFloor(1).GetLevel())
}

func TestShouldAddOneMainCorridorAndTwoSubCorridor(t *testing.T) {
	floor := NewFloor(1).
		addCorridors([]*Corridor{NewCorridor(MAIN, 1)}, MAIN).
		addCorridors([]*Corridor{NewCorridor(SUB, 1), NewCorridor(SUB, 2)}, SUB)
	assert.Equal(t, 1, len(floor.getCorridors(MAIN)))
	assert.Equal(t, 1, floor.getCorridors(MAIN)[0].getId())
	assert.Equal(t, 2, len(floor.getCorridors(SUB)))
	assert.Equal(t, 1, floor.getCorridors(SUB)[0].getId())
	assert.Equal(t, 2, floor.getCorridors(SUB)[1].getId())
}
