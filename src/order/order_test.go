package order

import (
	"github.com/ecommerce-study/src/config/clock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestNewOrder(t *testing.T) {
	t.Run("should not create an order with invalid CPF", func(t *testing.T) {
		// given
		invalidCpf := "93847575438"
		// when
		order, err := NewOrder(invalidCpf)
		// then
		assert.Error(t, err)
		assert.Nil(t, order)
	})

	t.Run("should create an order with three items", func(t *testing.T) {
		// given
		order, err := NewOrder("17185070031")
		require.NoError(t, err)
		// when
		order.AddItem(NewItem(1, "Instrumentos Musicais", "Guitarra", 1000.0), 1)
		order.AddItem(NewItem(2, "Instrumentos Musicais", "Amplificador", 5000.0), 1)
		order.AddItem(NewItem(3, "Instrumentos Musicais", "Cabo", 30.0), 3)
		total := order.Total()
		// then
		assert.Equal(t, 6090.0, total)
	})

	t.Run("should create an order with three items and a discount coupon", func(t *testing.T) {
		mockClock := clock.GetClockMock()
		// given
		order, err := NewOrder("17185070031")
		require.NoError(t, err)
		order.AddItem(NewItem(1, "Instrumentos Musicais", "Guitarra", 1000.0), 1)
		order.AddItem(NewItem(2, "Instrumentos Musicais", "Amplificador", 5000.0), 1)
		order.AddItem(NewItem(3, "Instrumentos Musicais", "Cabo", 30.0), 3)
		expirationDate := time.Date(2022, 8, 24, 0, 0, 0, 0, time.UTC)
		coupon := NewCoupon(20, "20OFF", expirationDate)
		mockClock.Set(expirationDate.Add(-1 * time.Hour))
		// when
		err = order.AddCoupon(coupon)
		total := order.Total()
		// when
		assert.NoError(t, err)
		assert.Equal(t, 4872.0, total)
	})

	t.Run("should not create an order with an expired coupon", func(t *testing.T) {
		mockClock := clock.GetClockMock()
		// given
		expirationDate := time.Date(2022, 8, 24, 0, 0, 0, 0, time.UTC)
		coupon := NewCoupon(20, "20OFF", expirationDate)
		order, err := NewOrder("17185070031")
		require.NoError(t, err)
		mockClock.Set(expirationDate.Add(1 * time.Hour))
		// when
		err = order.AddCoupon(coupon)
		// then
		assert.Error(t, err)
	})
}
