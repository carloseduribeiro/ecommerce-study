package place_order

import "time"

type ItemInput struct {
	IdItem   int
	Quantity uint
}

type PlaceOrderInput struct {
	Cpf        string
	OrderItems []ItemInput
	Coupon     *string
	IssueDate  time.Time
}

type PlaceOrderOutput struct {
	total float64
	code  string
}

func NewPlaceOrderOutput(total float64, code string) PlaceOrderOutput {
	return PlaceOrderOutput{total: total, code: code}
}

func (p PlaceOrderOutput) Total() float64 {
	return p.total
}

func (p PlaceOrderOutput) Code() string {
	return p.code
}
