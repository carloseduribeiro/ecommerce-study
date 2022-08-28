package entity

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestItem(t *testing.T) {
	t.Run("should create an item with dimension and calculate volume", func(t *testing.T) {
		// given
		dimension := NewDimension(100.0, 30.0, 10.0)
		item := NewItem(1, "Instrumentos Musicais", "Guitarra", 1000, WithDimension(dimension))
		// when
		result := item.Volume()
		// then
		assert.Equal(t, 0.03, result)
	})

	t.Run("should create an item with dimensions and calculate density", func(t *testing.T) {
		// given
		item := NewItem(1, "Instrumentos Musicais", "Guitarra", 1000, WithDimensions(100, 30, 10), WithWeight(3))
		// when
		result := item.Density()
		// then
		assert.Equal(t, 100.0, result)
	})

	t.Run("should create an item without dimensions and return density and volume equal to zero", func(t *testing.T) {
		// given
		item := NewItem(1, "Instrumentos Musicais", "Guitarra", 1000)
		// when
		volume := item.Volume()
		density := item.Density()
		// then
		assert.Zero(t, volume)
		assert.Zero(t, density)
	})
}
