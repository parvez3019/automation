package appliances

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewElectricalAppliance(t *testing.T) {
	appliance := NewAppliance(1, 10, LIGHT)
	assert.Equal(t, 1, appliance.GetId())
	assert.Equal(t, 10, appliance.GetPowerConsumption())
	assert.False(t, appliance.IsOn())
	assert.Equal(t, "Light", appliance.GetType())
}

func TestShouldTurnOnTheElectricalAppliance(t *testing.T) {
	appliance := NewAppliance(1, 10, LIGHT)
	appliance.SetSwitchedOn(true)
	assert.True(t, appliance.IsOn())
}

func TestShouldTurnOFFTheElectricalAppliance(t *testing.T) {
	appliance := NewAppliance(1, 10, LIGHT)
	appliance.SetSwitchedOn(true)
	assert.True(t, appliance.IsOn())
	appliance.SetSwitchedOn(false)
	assert.False(t, appliance.IsOn())
}
