package entity

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDimension(t *testing.T) {
	t.Run("should create the item dimensions", func(t *testing.T) {
		// given
		dimension := NewDimension(100.0, 30.0, 10.0)
		// when
		volume := dimension.Volume()
		// then
		assert.Equal(t, 0.03, volume)
	})
}
