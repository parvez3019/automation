package appliances

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewLightBulb(t *testing.T) {
	lb := NewLightBulb(1, 10)
	assert.Equal(t, 1, lb.getId())
	assert.Equal(t, 10, lb.getPowerConsumption())
	assert.False(t, lb.isOn())
}