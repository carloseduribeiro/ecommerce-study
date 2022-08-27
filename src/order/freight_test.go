package order

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFreight(t *testing.T) {
	t.Run("should calculate item freight", func(t *testing.T) {
		// given
		item := NewItem(1, "Instrumentos Musicais", "Guitarra", 1000, WithDimensions(100, 30, 10), WithWeight(3))
		freight := NewFreight()
		// when
		freight.AddItem(item, 2)
		total := freight.Total()
		// then
		assert.Equal(t, 60.0, total)
	})

	t.Run("should calculate minimum item freight", func(t *testing.T) {
		// given
		item := NewItem(3, "Instrumentos Musicais", "Cabo", 30.0, WithDimensions(10, 10, 10), WithWeight(0.9))
		freight := NewFreight()
		// when
		freight.AddItem(item, 1)
		total := freight.Total()
		// then
		assert.Equal(t, MinimumFreightValue, total)
	})
}
