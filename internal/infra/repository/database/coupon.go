package database

import (
	"github.com/ecommerce-study/internal/domain/entity"
	"github.com/ecommerce-study/internal/infra/database"
	"time"
)

type CouponRepository struct {
	conn database.Connection
}

func NewCouponRepository(connection database.Connection) *CouponRepository {
	return &CouponRepository{connection}
}

func (c *CouponRepository) GetByCode(code string) (*entity.Coupon, error) {
	stmt := "SELECT code, percentage, expire_date FROM ccca.coupon WHERE code = $1"
	row, _ := c.conn.QueryRow(stmt, code)
	dto := couponDto{}
	err := row.Scan(&dto.code, &dto.percentage, &dto.expireDate)
	if err != nil {
		return nil, err
	}
	coupon := dto.toEntity()
	return &coupon, nil
}

type couponDto struct {
	code       string
	percentage uint8
	expireDate *time.Time
}

func (c couponDto) toEntity() entity.Coupon {
	couponOptions := make([]entity.CouponOption, 0, 1)
	if c.expireDate != nil {
		couponOptions = append(couponOptions, entity.WithExpireDate(*c.expireDate))
	}
	return entity.NewCoupon(c.percentage, c.code, couponOptions...)
}
