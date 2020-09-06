package model

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestShouldReturnNewFloorWithLevel1(t *testing.T) {
	assert.Equal(t, 1,  NewFloor(1).getLevel())
}
