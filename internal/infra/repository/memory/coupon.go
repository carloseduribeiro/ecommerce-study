package memory

import (
	"github.com/ecommerce-study/internal/domain/entity"
)

type CouponRepositoryMemory struct {
	coupons []entity.Coupon
}

func NewCouponRepository() CouponRepositoryMemory {
	coupons := []entity.Coupon{
		entity.NewCoupon(20, "VALE20"),
	}
	return CouponRepositoryMemory{coupons: coupons}
}

func (c CouponRepositoryMemory) GetByCode(code string) (*entity.Coupon, error) {
	for _, coupon := range c.coupons {
		if coupon.Code() == code {
			return &coupon, nil
		}
	}
	return nil, nil
}
