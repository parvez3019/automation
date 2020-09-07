package appliances

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewLightBulb(t *testing.T) {
	lb := NewLightBulb(1, 10)
	assert.Equal(t, 1, lb.GetId())
	assert.Equal(t, 10, lb.GetPowerConsumption())
	assert.Equal(t, "Light", lb.GetType())
	assert.False(t, lb.IsOn())
}
