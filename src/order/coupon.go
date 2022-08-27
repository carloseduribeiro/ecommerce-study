package order

import (
	"github.com/ecommerce-study/src/config/clock"
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

func (coupon *Coupon) itsExpired() bool {
	if coupon.expirationDate.Sub(clock.Time.Now()) > 0 {
		return false
	}
	return true
}
