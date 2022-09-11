package repository

import (
	"github.com/ecommerce-study/internal/domain/entity"
)

type CouponRepository interface {
	GetByCode(code string) (*entity.Coupon, error)
}
