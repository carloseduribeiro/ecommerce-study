package database

import (
	"github.com/ecommerce-study/internal/domain/entity"
	"github.com/ecommerce-study/internal/infra/database"
)

type ItemRepository struct {
	conn database.Connection
}

func NewItemRepository(connection database.Connection) *ItemRepository {
	return &ItemRepository{conn: connection}
}

func (i *ItemRepository) GetById(id int) (*entity.Item, error) {
	stmt := "select id, category, description, price, width, height, length, weight from ccca.item where id = $1"
	row, err := i.conn.QueryRow(stmt, id)
	if err != nil {
		return nil, err
	}
	dto := itemDto{}
	err = row.Scan(
		&dto.id, &dto.category, &dto.description, &dto.price, &dto.width, &dto.height, &dto.length, &dto.weight,
	)
	if err != nil {
		return nil, err
	}
	item := dto.toEntity()
	return &item, nil
}

type itemDto struct {
	id          int
	category    string
	description string
	price       float64
	weight      float64
	width       float64
	height      float64
	length      float64
}

func (i *itemDto) toEntity() entity.Item {
	return entity.NewItem(
		i.id,
		i.category,
		i.description,
		i.price,
		entity.WithWeight(i.weight),
		entity.WithDimensions(i.width, i.height, i.length),
	)
}
