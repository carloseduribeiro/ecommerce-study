package factory

import "github.com/ecommerce-study/internal/domain/repository"

type Repository interface {
	CreateItemRepository() repository.ItemRepository
	CreateCouponRepository() repository.CouponRepository
	CreateOrderRepository() repository.OrderRepository
}
