package validate_coupon

import (
	"github.com/ecommerce-study/internal/config/clock"
	"github.com/ecommerce-study/internal/domain/repository"
)

type ValidateCoupon struct {
	couponRepository repository.CouponRepository
}

func NewValidateCoupon(couponRepository repository.CouponRepository) ValidateCoupon {
	return ValidateCoupon{
		couponRepository: couponRepository,
	}
}

func (v *ValidateCoupon) Execute(code string) (bool, error) {
	coupon, err := v.couponRepository.GetByCode(code)
	if coupon == nil {
		return false, err
	}
	return !coupon.ItsExpired(clock.Time.Now()), nil
}
