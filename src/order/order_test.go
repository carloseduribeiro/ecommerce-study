package order

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestNewOrder(t *testing.T) {
	t.Run("should not create an order with invalid CPF", func(t *testing.T) {
		order, err := NewOrder("93847575438")
		assert.Error(t, err)
		assert.Nil(t, order)
	})

	t.Run("should create an order with three items", func(t *testing.T) {
		order, err := NewOrder("17185070031")
		require.NoError(t, err)
		order.AddItem(NewItem(1, "Instrumentos Musicais", "Guitarra", 1000.0), 1)
		order.AddItem(NewItem(2, "Instrumentos Musicais", "Amplificador", 5000.0), 1)
		order.AddItem(NewItem(3, "Instrumentos Musicais", "Cabo", 30.0), 3)
		total := order.Total()
		assert.Equal(t, 6090.0, total)
	})

	t.Run("should create an order with three items with a discount coupon", func(t *testing.T) {
		order, err := NewOrder("17185070031")
		require.NoError(t, err)
		order.AddItem(NewItem(1, "Instrumentos Musicais", "Guitarra", 1000.0), 1)
		order.AddItem(NewItem(2, "Instrumentos Musicais", "Amplificador", 5000.0), 1)
		order.AddItem(NewItem(3, "Instrumentos Musicais", "Cabo", 30.0), 3)
		coupon := NewCoupon(20, "20OFF")
		order.AddCoupon(coupon)
		total := order.Total()
		assert.Equal(t, 4872.0, total)
	})
}
