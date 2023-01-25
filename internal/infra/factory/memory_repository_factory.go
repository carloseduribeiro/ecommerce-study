package factory

import (
	"github.com/ecommerce-study/internal/domain/repository"
	"github.com/ecommerce-study/internal/infra/repository/memory"
)

type MemoryRepositoryFactory struct {
}

func NewMemoryRepositoryFactory() MemoryRepositoryFactory {
	return MemoryRepositoryFactory{}
}

func (m *MemoryRepositoryFactory) CreateItemRepository() repository.ItemRepository {
	memoryItemRepository := memory.NewItemRepository()
	return memoryItemRepository
}
func (m *MemoryRepositoryFactory) CreateCouponRepository() repository.CouponRepository {
	memoryCouponRepository := memory.NewCouponRepository()
	return memoryCouponRepository
}
func (m *MemoryRepositoryFactory) CreateOrderRepository() repository.OrderRepository {
	memoryOrderRepository := memory.NewOrderRepository()
	return &memoryOrderRepository
}
