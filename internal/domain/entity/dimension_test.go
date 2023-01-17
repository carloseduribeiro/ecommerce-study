package entity

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDimension(t *testing.T) {
	t.Run("should create the item dimensions", func(t *testing.T) {
		// given
		width, height, length := 100.0, 30.0, 10.0
		// when
		dimension := NewDimension(width, height, length)
		// then
		assert.Equal(t, width, dimension.width)
		assert.Equal(t, height, dimension.height)
		assert.Equal(t, length, dimension.length)
	})

	t.Run("should calculate volume", func(t *testing.T) {
		// given
		width, height, length := 100.0, 30.0, 10.0
		expected := 0.03
		// when
		dimension := NewDimension(width, height, length)
		obtained := dimension.Volume()
		// then
		assert.Equal(t, expected, obtained)
	})

	t.Run("should return volume equal to zero when length is zero", func(t *testing.T) {
		// given
		width, height, length := 100.0, 30.0, 0.0
		expected := 0.0
		// when
		dimension := NewDimension(width, height, length)
		obtained := dimension.Volume()
		// then
		assert.Equal(t, expected, obtained)
	})

	t.Run("should return volume equal to zero when height is zero", func(t *testing.T) {
		// given
		width, height, length := 100.0, 0.0, 10.0
		expected := 0.0
		// when
		dimension := NewDimension(width, height, length)
		obtained := dimension.Volume()
		// then
		assert.Equal(t, expected, obtained)
	})

	t.Run("should return volume equal to zero when width is zero", func(t *testing.T) {
		// given
		width, height, length := 0.0, 30.0, 10.0
		expected := 0.0
		// when
		dimension := NewDimension(width, height, length)
		obtained := dimension.Volume()
		// then
		assert.Equal(t, expected, obtained)
	})
}
