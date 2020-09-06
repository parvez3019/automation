package appliances

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewElectricalAppliance(t *testing.T) {
	appliance := NewAppliance(1, 10)
	assert.Equal(t, 1, appliance.GetId())
	assert.Equal(t, 10, appliance.GetPowerConsumption())
	assert.False(t, appliance.IsOn())
}

func TestShouldTurnOnTheElectricalAppliance(t *testing.T) {
	appliance := NewAppliance(1, 10)
	appliance.TurnOn()
	assert.True(t, appliance.IsOn())
}

func TestShouldTurnOFFTheElectricalAppliance(t *testing.T) {
	appliance := NewAppliance(1, 10)
	appliance.TurnOn()
	assert.True(t, appliance.IsOn())
	appliance.TurnOff()
	assert.False(t, appliance.IsOn())
}