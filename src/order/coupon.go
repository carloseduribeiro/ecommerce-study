package order

type Coupon struct {
	percentage uint8
	code       string
}

func NewCoupon(percentage uint8, code string) Coupon {
	return Coupon{percentage: percentage, code: code}
}
