package entity

import (
	"time"
)

type OrderOption func(*Order)

type Order struct {
	cpf        CPF
	orderItems []OrderItem
	coupon     *Coupon
	issueDate  time.Time
	freight    Freight
	code       OrderCode
}

func NewOrder(cpf string, issueDate time.Time, sequence int, opts ...OrderOption) (*Order, error) {
	orderCpf, err := NewCPF(cpf)
	if err != nil {
		return nil, err
	}
	order := &Order{
		cpf:       *orderCpf,
		freight:   NewFreight(),
		issueDate: issueDate,
		code:      NewOrderCode(issueDate, sequence),
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

func (o *Order) Code() OrderCode {
	return o.code
}
