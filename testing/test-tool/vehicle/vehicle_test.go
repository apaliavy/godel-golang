package vehicle

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewVehicle(t *testing.T) {
	engine := &Engine{
		name: "BMW",
	}
	wheels := []Wheel{
		{color: "black"}, {color: "black"}, {color: "black"}, {color: "black"},
	}

	v, err := NewVehicle(engine, wheels)

	require.NoError(t, err)
	require.NotNil(t, v)

	assert.Equal(t, "BMW", v.engine.name)
	assert.Len(t, v.wheels, 4)

	for _, w := range v.wheels {
		assert.Equal(t, "black", w.color)
	}
}
