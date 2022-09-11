package entity

import (
	"time"
)

type CouponOption func(*Coupon)

func WithExpirationDate(expirationDate time.Time) CouponOption {
	return func(coupon *Coupon) {
		coupon.expirationDate = &expirationDate
	}
}

type Coupon struct {
	percentage     uint8
	code           string
	expirationDate *time.Time
}

func NewCoupon(percentage uint8, code string, opts ...CouponOption) Coupon {
	coupon := Coupon{percentage: percentage, code: code}
	for _, opt := range opts {
		opt(&coupon)
	}
	return coupon
}

func (c *Coupon) ItsExpired(dateTimeNow time.Time) bool {
	if c.expirationDate != nil {
		if dateTimeNow.Sub(*c.expirationDate) > 0 {
			return true
		}
	}
	return false
}

func (c *Coupon) CalculateDiscount(amount float64) float64 {
	return (float64(c.percentage) * amount) / 100.0
}

func (c *Coupon) Code() string {
	return c.code
}
