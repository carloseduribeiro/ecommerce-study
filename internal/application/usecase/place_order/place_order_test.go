package place_order

import (
	memory2 "github.com/ecommerce-study/internal/infra/repository/memory"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestPlaceOrder(t *testing.T) {
	t.Run("should place order", func(t *testing.T) {
		// given
		itemRepository := memory2.NewItemRepository()
		orderRepository := memory2.NewOrderRepository()
		couponRepository := memory2.NewCouponRepository()
		coupon := "VALE20"
		input := PlaceOrderInput{
			Cpf: "17185070031",
			OrderItems: []ItemInput{
				{IdItem: 1, Quantity: 1},
				{IdItem: 2, Quantity: 1},
				{IdItem: 3, Quantity: 3},
			},
			Coupon: &coupon,
		}
		placeOrder := NewPlaceOrder(&itemRepository, &orderRepository, &couponRepository)
		// when
		output := placeOrder.Execute(input)
		// then
		assert.Equal(t, 4872.0, output.Total())
	})

	t.Run("should place order and generate order code", func(t *testing.T) {
		// given
		itemRepository := memory2.NewItemRepository()
		orderRepository := memory2.NewOrderRepository()
		couponRepository := memory2.NewCouponRepository()
		coupon := "VALE20"
		issueDate := time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC)
		input := PlaceOrderInput{
			Cpf: "17185070031",
			OrderItems: []ItemInput{
				{IdItem: 1, Quantity: 1},
				{IdItem: 2, Quantity: 1},
				{IdItem: 3, Quantity: 3},
			},
			Coupon:    &coupon,
			IssueDate: issueDate,
		}
		placeOrder := NewPlaceOrder(&itemRepository, &orderRepository, &couponRepository)
		// when
		placeOrder.Execute(input)
		output := placeOrder.Execute(input)
		// then
		assert.Equal(t, "202100000002", output.Code())
	})
}
