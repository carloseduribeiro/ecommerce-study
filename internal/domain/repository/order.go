package repository

import (
	"github.com/ecommerce-study/internal/domain/entity"
)

type OrderRepository interface {
	Save(order *entity.Order) error
}
