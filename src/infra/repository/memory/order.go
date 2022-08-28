package memory

import "github.com/ecommerce-study/src/domain/entity"

type OrderRepositoryMemory struct {
	orders []entity.Order
}

func NewOrderRepository() OrderRepositoryMemory {
	orders := make([]entity.Order, 0)
	return OrderRepositoryMemory{orders: orders}
}

func (o OrderRepositoryMemory) Save(order *entity.Order) error {
	o.orders = append(o.orders, *order)
	return nil
}
