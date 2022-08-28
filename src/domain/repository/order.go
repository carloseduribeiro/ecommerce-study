package repository

import "github.com/ecommerce-study/src/domain/entity"

type OrderRepository interface {
	Save(order *entity.Order) error
}
