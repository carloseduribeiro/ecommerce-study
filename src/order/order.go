package order

import "errors"

type Order struct {
	cpf        CPF
	orderItems []OrderItem
	coupon     *Coupon
}

func NewOrder(cpf string) (*Order, error) {
	orderCpf, err := NewCPF(cpf)
	if err != nil {
		return nil, err
	}
	return &Order{cpf: *orderCpf}, nil
}

func (o *Order) AddItem(item Item, quantity int) {
	o.orderItems = append(o.orderItems, NewOrderItem(item, quantity, item.price))
}

func (o *Order) AddCoupon(coupon Coupon) error {
	if coupon.itsExpired() {
		return errors.New("expired coupon")
	}
	o.coupon = &coupon
	return nil
}

func (o *Order) Total() (total float64) {
	for _, orderItem := range o.orderItems {
		total += orderItem.Total()
	}
	if o.coupon != nil {
		total -= (total * float64(o.coupon.percentage)) / 100
	}
	return
}
