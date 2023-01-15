package entity

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestCoupon(t *testing.T) {
	t.Run("should create an expired Coupon", func(t *testing.T) {
		// given
		expireDate := time.Date(2022, 8, 27, 12, 0, 0, 0, time.UTC)
		coupon := NewCoupon(0, "", WithExpireDate(expireDate))
		timeNow := expireDate.Add(1 * time.Second)
		// when
		result := coupon.ItsExpired(timeNow)
		// when
		assert.True(t, result)
	})

	t.Run("should create a not expired Coupon", func(t *testing.T) {
		// given
		expireDate := time.Date(2022, 8, 27, 12, 0, 0, 0, time.UTC)
		coupon := NewCoupon(0, "", WithExpireDate(expireDate))
		timeNow := expireDate.Add(-1 * time.Second)
		// when
		result := coupon.ItsExpired(timeNow)
		// when
		assert.False(t, result)
	})

	t.Run("should create a not expired Coupon and calculate the discount", func(t *testing.T) {
		// given
		expireDate := time.Date(2022, 8, 27, 12, 0, 0, 0, time.UTC)
		coupon := NewCoupon(20, "", WithExpireDate(expireDate))
		timeNow := expireDate.Add(-1 * time.Second)
		amount := 1000.0
		// when
		itsExpired := coupon.ItsExpired(timeNow)
		discount := coupon.CalculateDiscount(amount)
		// when
		assert.False(t, itsExpired)
		assert.Equal(t, discount, 200.0)
	})

	t.Run("should create a Coupon that never expires", func(t *testing.T) {
		// given
		coupon := NewCoupon(20, "")
		amount := 1000.0
		// when
		itsExpired := coupon.ItsExpired(time.Now())
		discount := coupon.CalculateDiscount(amount)
		// when
		assert.False(t, itsExpired)
		assert.Equal(t, discount, 200.0)
	})
}
