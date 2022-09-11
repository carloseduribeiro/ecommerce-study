package memory

import (
	"github.com/ecommerce-study/internal/domain/entity"
)

type ItemRepositoryMemory struct {
	items []entity.Item
}

func NewItemRepository() ItemRepositoryMemory {
	items := []entity.Item{
		entity.NewItem(1, "Instrumentos Musicais", "Guitarra", 1000.0),
		entity.NewItem(2, "Instrumentos Musicais", "Amplificador", 5000.0),
		entity.NewItem(3, "Instrumentos Musicais", "Cabo", 30.0),
	}
	return ItemRepositoryMemory{items: items}
}

func (i ItemRepositoryMemory) GetById(id int) (*entity.Item, error) {
	for _, item := range i.items {
		if item.Id() == id {
			return &item, nil
		}
	}
	return nil, nil
}
