package model

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestShouldReturnNewFloorWithLevel1(t *testing.T) {
	assert.Equal(t, 1, NewFloor(1).GetLevel())
}

func TestShouldReturnEmptyListInCaseOfInvalidCorridor(t *testing.T)  {
	floor := NewFloor(1).
		AddCorridors([]*Corridor{NewCorridor(MAIN, 1)})
	assert.Len(t, floor.GetCorridors("ABC"), 0)
}

func TestShouldAddOneMainCorridorAndTwoSubCorridor(t *testing.T) {
	floor := NewFloor(1).
		AddCorridors([]*Corridor{NewCorridor(MAIN, 1)}).
		AddCorridors([]*Corridor{NewCorridor(SUB, 1), NewCorridor(SUB, 2)})
	assert.Equal(t, 1, len(floor.GetCorridors(MAIN)))
	assert.Equal(t, 1, floor.GetCorridors(MAIN)[0].GetId())

	assert.Equal(t, 2, len(floor.GetCorridors(SUB)))
	assert.Equal(t, 1, floor.GetCorridors(SUB)[0].GetId())
	assert.Equal(t, 2, floor.GetCorridors(SUB)[1].GetId())
}
