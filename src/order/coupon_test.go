package order

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestCoupon(t *testing.T) {
	t.Run("should create an expired Coupon", func(t *testing.T) {
		// given
		expirationDate := time.Date(2022, 8, 27, 12, 0, 0, 0, time.UTC)
		coupon := NewCoupon(0, "", expirationDate)
		timeNow := expirationDate.Add(1 * time.Second)
		// when
		result := coupon.ItsExpired(timeNow)
		// when
		assert.True(t, result)
	})

	t.Run("should create a not expired Coupon", func(t *testing.T) {
		// given
		expirationDate := time.Date(2022, 8, 27, 12, 0, 0, 0, time.UTC)
		coupon := NewCoupon(0, "", expirationDate)
		timeNow := expirationDate.Add(-1 * time.Second)
		// when
		result := coupon.ItsExpired(timeNow)
		// when
		assert.False(t, result)
	})

	t.Run("should create a not expired Coupon and calculate the discount", func(t *testing.T) {
		// given
		expirationDate := time.Date(2022, 8, 27, 12, 0, 0, 0, time.UTC)
		coupon := NewCoupon(20, "", expirationDate)
		timeNow := expirationDate.Add(-1 * time.Second)
		amount := 1000.0
		// when
		itsExpirted := coupon.ItsExpired(timeNow)
		discount := coupon.CalculateDiscount(amount)
		// when
		assert.False(t, itsExpirted)
		assert.Equal(t, discount, 200.0)
	})
}
