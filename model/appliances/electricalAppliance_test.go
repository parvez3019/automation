package appliances

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewElectricalAppliance(t *testing.T) {
	appliance := NewAppliance(1, 10)
	assert.Equal(t, 1, appliance.getId())
	assert.Equal(t, 10, appliance.getPowerConsumption())
	assert.False(t, appliance.isOn())
}

func TestShouldTurnOnTheElectricalAppliance(t *testing.T) {
	appliance := NewAppliance(1, 10)
	appliance.turnOn()
	assert.True(t, appliance.isOn())
}