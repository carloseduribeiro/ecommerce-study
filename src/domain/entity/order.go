package entity

import (
	"time"
)

type OrderOption func(*Order)

func WithIssueDate(issueDate time.Time) OrderOption {
	return func(o *Order) {
		o.issueDate = issueDate
	}
}

type Order struct {
	cpf        CPF
	orderItems []OrderItem
	coupon     *Coupon
	issueDate  time.Time
	freight    Freight
}

func NewOrder(cpf string, opts ...OrderOption) (*Order, error) {
	orderCpf, err := NewCPF(cpf)
	if err != nil {
		return nil, err
	}
	order := &Order{
		cpf:     *orderCpf,
		freight: NewFreight(),
	}
	for _, opt := range opts {
		opt(order)
	}
	return order, nil
}

func (o *Order) AddItem(item Item, quantity int) {
	o.orderItems = append(o.orderItems, NewOrderItem(item, quantity, item.price))
	o.freight.AddItem(item, quantity)
}

func (o *Order) AddCoupon(coupon Coupon) {
	if !coupon.ItsExpired(o.issueDate) {
		o.coupon = &coupon
	}
}

func (o *Order) Total() (total float64) {
	for _, orderItem := range o.orderItems {
		total += orderItem.Total()
	}
	if o.coupon != nil {
		total -= o.coupon.CalculateDiscount(total)
	}
	total += o.freight.Total()
	return
}
