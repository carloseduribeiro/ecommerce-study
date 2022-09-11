package repository

import (
	"github.com/ecommerce-study/internal/domain/entity"
)

type ItemRepository interface {
	GetById(id int) (*entity.Item, error)
}
