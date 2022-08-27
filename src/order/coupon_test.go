package order

import (
	"github.com/ecommerce-study/src/config/clock"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestCoupon(t *testing.T) {
	mockClock := clock.GetClockMock()

	t.Run("should return true if it is expired", func(t *testing.T) {
		// given
		expirationDate := time.Date(2022, 8, 27, 12, 0, 0, 0, time.UTC)
		coupon := Coupon{expirationDate: expirationDate}
		mockClock.Set(expirationDate.Add(1 * time.Second))
		// when
		result := coupon.itsExpired()
		// when
		assert.True(t, result)
	})

	t.Run("should return false if it is not expired", func(t *testing.T) {
		// given
		expirationDate := time.Date(2022, 8, 27, 12, 0, 0, 0, time.UTC)
		coupon := Coupon{expirationDate: expirationDate}
		mockClock.Set(expirationDate.Add(-1 * time.Second))
		// when
		result := coupon.itsExpired()
		// when
		assert.False(t, result)
	})
}
