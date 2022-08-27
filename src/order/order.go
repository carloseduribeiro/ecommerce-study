package order

import (
	"time"
)

type Order struct {
	cpf        CPF
	orderItems []OrderItem
	coupon     *Coupon
	issueDate  time.Time
}

func NewOrder(cpf string, issueDate time.Time) (*Order, error) {
	orderCpf, err := NewCPF(cpf)
	if err != nil {
		return nil, err
	}
	return &Order{cpf: *orderCpf, issueDate: issueDate}, nil
}

func (o *Order) AddItem(item Item, quantity int) {
	o.orderItems = append(o.orderItems, NewOrderItem(item, quantity, item.price))
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
	return
}
