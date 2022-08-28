package repository

import "github.com/ecommerce-study/src/domain/entity"

type ItemRepository interface {
	GetById(id int) (*entity.Item, error)
}
