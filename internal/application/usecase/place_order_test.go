package usecase

import (
	memory2 "github.com/ecommerce-study/internal/infra/repository/memory"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPlaceOrder(t *testing.T) {
	t.Run("should place order", func(t *testing.T) {
		// given
		itemRepository := memory2.NewItemRepository()
		orderRepository := memory2.NewOrderRepository()
		couponRepository := memory2.NewCouponRepository()
		coupon := "VALE20"
		input := PlaceOrderInput{
			cpf: "17185070031",
			orderItems: []ItemInput{
				{IdItem: 1, Quantity: 1},
				{IdItem: 2, Quantity: 1},
				{IdItem: 3, Quantity: 3},
			},
			coupon: &coupon,
		}
		placeOrder := NewPlaceOrder(itemRepository, orderRepository, couponRepository)
		// when
		output := placeOrder.Execute(input)
		// then
		assert.Equal(t, 4872.0, output.Total())
	})
}
