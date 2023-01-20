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
	stmt := "SELECT id, category, description, price, width, height, length, weight FROM ecommerce.item WHERE id = $1"
	row, _ := i.conn.QueryRow(stmt, id)
	dto := itemDto{}
	err := row.Scan(
		&dto.id, &dto.category, &dto.description, &dto.price, &dto.width, &dto.height, &dto.length, &dto.weight,
	)
	if err != nil {
		return nil, err
	}
	return dto.toEntity(), nil
}

type itemDto struct {
	id          int
	category    string
	description string
	price       float64
	weight      *float64
	width       *float64
	height      *float64
	length      *float64
}

func (i *itemDto) toEntity() *entity.Item {
	itemOptions := make([]entity.ItemOption, 0, 2)
	if i.weight != nil {
		itemOptions = append(itemOptions, entity.WithWeight(*i.weight))
	}
	if i.width != nil && i.height != nil && i.length != nil {
		itemOptions = append(itemOptions, entity.WithDimensions(*i.width, *i.height, *i.length))
	}
	itemEntity, _ := entity.NewItem(i.id, i.category, i.description, i.price, itemOptions...)
	return itemEntity
}
