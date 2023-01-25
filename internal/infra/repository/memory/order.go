package memory

import (
	"fmt"
	"github.com/ecommerce-study/internal/domain/entity"
)

type OrderRepositoryMemory struct {
	orders []entity.Order
}

func NewOrderRepository() OrderRepositoryMemory {
	orders := make([]entity.Order, 0, 0)
	return OrderRepositoryMemory{orders: orders}
}

func (o *OrderRepositoryMemory) Save(order *entity.Order) error {
	o.orders = append(o.orders, *order)
	return nil
}

func (o *OrderRepositoryMemory) Count() (*int, error) {
	tmp := len(o.orders)
	return &tmp, nil
}

func (o *OrderRepositoryMemory) GetByCode(code string) (*entity.Order, error) {
	for _, order := range o.orders {
		if order.Code() == code {
			return &order, nil
		}
	}
	return nil, fmt.Errorf("order not found with the given code")
}
