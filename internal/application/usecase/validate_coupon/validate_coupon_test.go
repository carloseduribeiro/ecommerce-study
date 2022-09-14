package validate_coupon

import (
	"github.com/ecommerce-study/internal/config/clock"
	"github.com/ecommerce-study/internal/infra/repository/memory"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestValidateCoupon(t *testing.T) {
	t.Run("should validate a coupon", func(t *testing.T) {
		// given
		couponRepository := memory.NewCouponRepository()
		usecase := NewValidateCoupon(&couponRepository)
		clockMock := clock.GetClockMock()
		clockMock.Set(time.Date(2022, 1, 1, 0, 0, 0, 0, time.UTC))
		// when
		isValid, _ := usecase.Execute("VALE20")
		// when
		assert.True(t, isValid)
	})

	t.Run("should return false when the repository returns an error", func(t *testing.T) {
		// given
		couponRepository := memory.NewCouponRepository()
		usecase := NewValidateCoupon(&couponRepository)
		// when
		isValid, err := usecase.Execute("VALE100")
		// when
		assert.False(t, isValid)
		assert.Error(t, err)
	})

	t.Run("should validate a non-existing coupon", func(t *testing.T) {
		// given
		couponRepository := memory.NewCouponRepository()
		usecase := NewValidateCoupon(&couponRepository)
		// when
		isValid, err := usecase.Execute("VALE100")
		// when
		assert.Error(t, err)
		assert.False(t, isValid)
	})
}
