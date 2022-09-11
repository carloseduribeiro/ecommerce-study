package usecase

import (
	"github.com/ecommerce-study/internal/domain/entity"
	repository2 "github.com/ecommerce-study/internal/domain/repository"
)

type PlaceOrderInputOption func(*PlaceOrderInput)

func WithCoupon(code string) PlaceOrderInputOption {
	return func(o *PlaceOrderInput) {
		o.coupon = &code
	}
}

type ItemInput struct {
	IdItem   int
	Quantity int
}

type PlaceOrderInput struct {
	cpf        string
	orderItems []ItemInput
	coupon     *string
}

func NewPlaceOrderInput(cpf string, orderItems []ItemInput, opts ...PlaceOrderInputOption) PlaceOrderInput {
	result := PlaceOrderInput{cpf: cpf, orderItems: orderItems}
	for _, opt := range opts {
		opt(&result)
	}
	return result
}

type PlaceOrderOutput struct {
	total float64
}

func NewPlaceOrderOutput(total float64) PlaceOrderOutput {
	return PlaceOrderOutput{total: total}
}

func (p PlaceOrderOutput) Total() float64 {
	return p.total
}

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
	order, _ := entity.NewOrder(input.cpf)
	for _, orderItem := range input.orderItems {
		item, err := p.itemRepository.GetById(orderItem.IdItem)
		if err != nil {
			panic(err)
		}
		order.AddItem(*item, orderItem.Quantity)
	}
	if input.coupon != nil {
		coupon, err := p.couponRepository.GetByCode(*input.coupon)
		if err != nil {
			panic(err)
		}
		order.AddCoupon(*coupon)
	}
	err := p.orderRepository.Save(order)
	if err != nil {
		panic(err)
	}
	total := order.Total()
	output := NewPlaceOrderOutput(total)
	return output
}
