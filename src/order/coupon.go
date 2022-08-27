package order

import (
	"time"
)

type Coupon struct {
	percentage     uint8
	code           string
	expirationDate time.Time
}

func NewCoupon(percentage uint8, code string, expirationDate time.Time) Coupon {
	return Coupon{percentage: percentage, code: code, expirationDate: expirationDate}
}

func (coupon *Coupon) ItsExpired(dateTimeNow time.Time) bool {
	if coupon.expirationDate.Sub(dateTimeNow) > 0 {
		return false
	}
	return true
}

func (coupon *Coupon) CalculateDiscount(amount float64) float64 {
	return (float64(coupon.percentage) * amount) / 100.0
}
