package entity

import (
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
		order, err := NewOrder(invalidCpf, time.Now(), 0)
		// then
		assert.Error(t, err)
		assert.Nil(t, order)
	})

	t.Run("should create an order with three items", func(t *testing.T) {
		// given
		order, err := NewOrder("17185070031", time.Now(), 0)
		require.NoError(t, err)
		item1, err := NewItem(1, "Instrumentos Musicais", "Guitarra", 1000.0)
		require.NoError(t, err)
		item2, err := NewItem(2, "Instrumentos Musicais", "Amplificador", 5000.0)
		require.NoError(t, err)
		item3, err := NewItem(3, "Instrumentos Musicais", "Cabo", 30.0)
		require.NoError(t, err)
		order.AddItem(item1, 1)
		order.AddItem(item2, 1)
		order.AddItem(item3, 3)
		// when
		total := order.Total()
		// then
		assert.Equal(t, 6090.0, total)
	})

	t.Run("should create an order with three items and a discount coupon", func(t *testing.T) {
		// given
		expirationDate := time.Date(2022, 8, 24, 0, 0, 0, 0, time.UTC)
		coupon := NewCoupon(20, "20OFF", WithExpireDate(expirationDate))
		issueDate := expirationDate.Add(-1 * time.Hour)
		order, err := NewOrder("17185070031", issueDate, 1)
		require.NoError(t, err)
		item1, err := NewItem(1, "Instrumentos Musicais", "Guitarra", 1000.0)
		require.NoError(t, err)
		item2, err := NewItem(2, "Instrumentos Musicais", "Amplificador", 5000.0)
		require.NoError(t, err)
		item3, err := NewItem(3, "Instrumentos Musicais", "Cabo", 30.0)
		require.NoError(t, err)
		order.AddItem(item1, 1)
		order.AddItem(item2, 1)
		order.AddItem(item3, 3)
		// when
		order.AddCoupon(coupon)
		total := order.Total()
		// when
		assert.Equal(t, 4872.0, total)
	})

	t.Run("should create an order with three items and an expired discount coupon", func(t *testing.T) {
		// given
		expirationDate := time.Date(2022, 8, 24, 0, 0, 0, 0, time.UTC)
		coupon := NewCoupon(20, "20OFF", WithExpireDate(expirationDate))
		issueDate := expirationDate.Add(1 * time.Hour)
		order, err := NewOrder("17185070031", issueDate, 1)
		require.NoError(t, err)
		item1, err := NewItem(1, "Instrumentos Musicais", "Guitarra", 1000.0)
		require.NoError(t, err)
		item2, err := NewItem(2, "Instrumentos Musicais", "Amplificador", 5000.0)
		require.NoError(t, err)
		item3, err := NewItem(3, "Instrumentos Musicais", "Cabo", 30.0)
		require.NoError(t, err)
		order.AddItem(item1, 1)
		order.AddItem(item2, 1)
		order.AddItem(item3, 3)
		// when
		order.AddCoupon(coupon)
		total := order.Total()
		// when
		assert.Equal(t, 6090.0, total)
	})

	t.Run("should create an order with three items and calculate shipping", func(t *testing.T) {
		// given
		order, err := NewOrder("17185070031", time.Now(), 0)
		require.NoError(t, err)
		item1, err := NewItem(1, "Instrumentos Musicais", "Guitarra", 1000.0, WithDimensions(100, 30, 10), WithWeight(3))
		require.NoError(t, err)
		item2, err := NewItem(2, "Instrumentos Musicais", "Amplificador", 5000.0, WithDimensions(100, 50, 50), WithWeight(20))
		require.NoError(t, err)
		item3, err := NewItem(3, "Instrumentos Musicais", "Cabo", 30.0, WithDimensions(10, 10, 10), WithWeight(1))
		require.NoError(t, err)
		// when
		order.AddItem(item1, 1)
		order.AddItem(item2, 1)
		order.AddItem(item3, 3)
		total := order.Total()
		// then
		// shipping formula: volume * 1000 * (density/100)
		expectedTotal := 6090.0 + 30 + 200 + 10 + 10 + 10
		assert.Equal(t, expectedTotal, total)
	})

	t.Run("should create an order with three items and calculate minimum shipping", func(t *testing.T) {
		// given
		order, err := NewOrder("17185070031", time.Now(), 0)
		require.NoError(t, err)
		item, err := NewItem(3, "Instrumentos Musicais", "Cabo", 30.0, WithDimensions(10, 10, 10), WithWeight(0.9))
		require.NoError(t, err)
		// when
		order.AddItem(item, 1)
		total := order.Total()
		// then
		// shipping formula: volume * 1000 * (density/100)
		expectedTotal := 30.0 + MinimumFreightValue
		assert.Equal(t, expectedTotal, total)
	})

	t.Run("should create an order and calculate order code", func(t *testing.T) {
		// given
		issueDate := time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC)
		order, err := NewOrder("17185070031", issueDate, 1)
		require.NoError(t, err)
		// when
		obtained := order.Code()
		expected := "202100000001"
		// then
		assert.Equal(t, expected, obtained)
	})

	t.Run("should execute an order option function when its given to constructor", func(t *testing.T) {
		// given
		orderExecuted := false
		fakeOrderOption := func(order *Order) { orderExecuted = true }
		// when
		_, err := NewOrder("17185070031", time.Now(), 1, fakeOrderOption)
		require.NoError(t, err)
		// then
		assert.True(t, orderExecuted)
	})
}
