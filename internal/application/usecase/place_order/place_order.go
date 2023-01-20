package place_order

import (
	"github.com/ecommerce-study/internal/domain/entity"
	"github.com/ecommerce-study/internal/domain/repository"
)

type PlaceOrder struct {
	itemRepository   repository.ItemRepository
	orderRepository  repository.OrderRepository
	couponRepository repository.CouponRepository
}

func NewPlaceOrder(itemRepository repository.ItemRepository, orderRepository repository.OrderRepository, couponRepository repository.CouponRepository) PlaceOrder {
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
		var item *entity.Item
		item, err = p.itemRepository.GetById(orderItem.IdItem)
		if err != nil {
			panic(err)
		}
		order.AddItem(item, orderItem.Quantity)
	}
	if input.Coupon != nil {
		var coupon *entity.Coupon
		coupon, err = p.couponRepository.GetByCode(*input.Coupon)
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
