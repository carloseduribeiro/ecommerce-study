package repository

import "github.com/ecommerce-study/src/domain/entity"

type CouponRepository interface {
	GetByCode(code string) (*entity.Coupon, error)
}
