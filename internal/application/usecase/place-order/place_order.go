package place_order

import (
	"github.com/ecommerce-study/internal/domain/entity"
	repository2 "github.com/ecommerce-study/internal/domain/repository"
)

type PlaceOrder struct {
	itemRepository   repository2.ItemRepository
	orderRepository  repository2.OrderRepository
	couponRepository repository2.CouponRepository
}

func NewPlaceOrder(itemRepository repository2.ItemRepository, orderRepository repository2.OrderRepository, couponRepository repository2.CouponRepository) PlaceOrder {
	return PlaceOrder{
		itemRepository:   itemRepository,
		orderRepository:  orderRepository,
		couponRepository: couponRepository,
	}
}

func (p PlaceOrder) Execute(input PlaceOrderInput) PlaceOrderOutput {
	count, err := p.orderRepository.Count()
	if err != nil {
		panic(err)
	}
	sequence := *count + 1
	order, _ := entity.NewOrder(input.Cpf, input.IssueDate, sequence)
	for _, orderItem := range input.OrderItems {
		item, err := p.itemRepository.GetById(orderItem.IdItem)
		if err != nil {
			panic(err)
		}
		order.AddItem(*item, orderItem.Quantity)
	}
	if input.Coupon != nil {
		coupon, err := p.couponRepository.GetByCode(*input.Coupon)
		if err != nil {
			panic(err)
		}
		order.AddCoupon(*coupon)
	}
	err = p.orderRepository.Save(order)
	if err != nil {
		panic(err)
	}
	total := order.Total()
	output := NewPlaceOrderOutput(total, order.Code().Value())
	return output
}
