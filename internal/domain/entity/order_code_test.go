package entity

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestOrderCode(t *testing.T) {
	t.Run("should create an order code", func(t *testing.T) {
		// given
		issueDate := time.Date(2022, time.September, 1, 1, 0, 0, 0, time.UTC)
		sequence := 1
		// when
		orderCode := NewOrderCode(issueDate, sequence)
		code := orderCode.Value()
		expected := "202200000001"
		// then
		assert.Equal(t, expected, code)
	})
}
