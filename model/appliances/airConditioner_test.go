package appliances

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewAirConditioner(t *testing.T) {
	ac := NewAirConditioner(1, 10)
	assert.Equal(t, 1, ac.getId())
	assert.Equal(t, 10, ac.getPowerConsumption())
	assert.False(t, ac.isOn())
}