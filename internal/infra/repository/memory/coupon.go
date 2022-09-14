package memory

import (
	"errors"
	"github.com/ecommerce-study/internal/domain/entity"
	"time"
)

type CouponRepositoryMemory struct {
	coupons []entity.Coupon
}

func NewCouponRepository() CouponRepositoryMemory {
	coupons := []entity.Coupon{
		entity.NewCoupon(20, "VALE20", entity.WithExpirationDate(time.Date(2022, 9, 14, 0, 0, 0, 0, time.UTC))),
	}
	return CouponRepositoryMemory{coupons: coupons}
}

func (c CouponRepositoryMemory) GetByCode(code string) (*entity.Coupon, error) {
	for _, coupon := range c.coupons {
		if coupon.Code() == code {
			return &coupon, nil
		}
	}
	return nil, errors.New("coupon not found")
}
