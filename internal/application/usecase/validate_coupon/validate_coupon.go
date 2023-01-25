package validate_coupon

import (
	"github.com/ecommerce-study/internal/config/clock"
	"github.com/ecommerce-study/internal/domain/factory"
	"github.com/ecommerce-study/internal/domain/repository"
)

type ValidateCoupon struct {
	couponRepository repository.CouponRepository
}

func NewValidateCoupon(repositoryFactory factory.Repository) ValidateCoupon {
	return ValidateCoupon{
		couponRepository: repositoryFactory.CreateCouponRepository(),
	}
}

func (v *ValidateCoupon) Execute(code string) (bool, error) {
	coupon, err := v.couponRepository.GetByCode(code)
	if coupon == nil {
		return false, err
	}
	return !coupon.ItsExpired(clock.Time.Now()), nil
}
