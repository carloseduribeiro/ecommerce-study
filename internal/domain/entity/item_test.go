package entity

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestItem(t *testing.T) {
	t.Run("should create an item with dimension and calculate volume", func(t *testing.T) {
		// given
		dimension := NewDimension(100.0, 30.0, 10.0)
		item, err := NewItem(1, "Instrumentos Musicais", "Guitarra", 1000, WithDimension(dimension))
		require.NoError(t, err)
		// when
		result := item.Volume()
		// then
		assert.Equal(t, 0.03, result)
	})

	t.Run("should create an item with dimensions and calculate density", func(t *testing.T) {
		// given
		item, err := NewItem(1, "Instrumentos Musicais", "Guitarra", 1000, WithDimensions(100, 30, 10), WithWeight(3))
		require.NoError(t, err)
		// when
		result := item.Density()
		// then
		assert.Equal(t, 100.0, result)
	})

	t.Run("should create an item without dimensions and return density and volume equal to zero", func(t *testing.T) {
		// given
		item, err := NewItem(1, "Instrumentos Musicais", "Guitarra", 1000)
		require.NoError(t, err)
		// when
		volume := item.Volume()
		density := item.Density()
		// then
		assert.Zero(t, volume)
		assert.Zero(t, density)
	})
}

func TestWithDimensionsItemOption(t *testing.T) {
	t.Run("should return an error when the given width is less than zero", func(t *testing.T) {
		// given
		width, height, length := -0.1, 0.0, 0.0
		item := &Item{}
		itemOption := WithDimensions(width, height, length)
		// when
		err := itemOption(item)
		// then
		assert.Error(t, err)
	})

	t.Run("should return an error when the given height is less than zero", func(t *testing.T) {
		// given
		width, height, length := 0.0, -1.0, 0.0
		item := &Item{}
		itemOption := WithDimensions(width, height, length)
		// when
		err := itemOption(item)
		// then
		assert.Error(t, err)
	})

	t.Run("should return an error when the given length is less than zero", func(t *testing.T) {
		// given
		width, height, length := 0.0, 0.0, -1.0
		item := &Item{}
		itemOption := WithDimensions(width, height, length)
		// when
		err := itemOption(item)
		// then
		assert.Error(t, err)
	})

	t.Run("should return nil when all parameters are OK", func(t *testing.T) {
		// given
		width, height, length := 0.0, 0.0, 0.0
		item := &Item{}
		itemOption := WithDimensions(width, height, length)
		// when
		err := itemOption(item)
		// then
		assert.NoError(t, err)
	})
}

func TestWithWeightItemOption(t *testing.T) {
	t.Run("should return an error when the given weight is less than zero", func(t *testing.T) {
		// given
		weigth := -0.1
		item := &Item{}
		itemOption := WithWeight(weigth)
		// when
		err := itemOption(item)
		// then
		assert.Error(t, err)
	})

	t.Run("should return nil when all parameters are OK", func(t *testing.T) {
		// given
		weigth := 0.0
		item := &Item{}
		itemOption := WithWeight(weigth)
		// when
		err := itemOption(item)
		// then
		assert.NoError(t, err)
	})
}
