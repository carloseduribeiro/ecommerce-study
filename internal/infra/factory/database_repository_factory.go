package factory

import (
	"github.com/ecommerce-study/internal/domain/repository"
	"github.com/ecommerce-study/internal/infra/database"
	repositoryDB "github.com/ecommerce-study/internal/infra/repository/database"
	repositoryMemory "github.com/ecommerce-study/internal/infra/repository/memory"
)

type DatabaseRepositoryFactory struct {
	connection database.Connection
}

func NewDatabaseRepositoryFactory(conn database.Connection) DatabaseRepositoryFactory {
	return DatabaseRepositoryFactory{connection: conn}
}

func (d DatabaseRepositoryFactory) CreateItemRepository() repository.ItemRepository {
	return repositoryDB.NewItemRepository(d.connection)
}

func (d DatabaseRepositoryFactory) CreateCouponRepository() repository.CouponRepository {
	return repositoryDB.NewCouponRepository(d.connection)
}

func (d DatabaseRepositoryFactory) CreateOrderRepository() repository.OrderRepository {
	// TODO - implement database repository to return here
	memoryOrderRepository := repositoryMemory.NewOrderRepository()
	return &memoryOrderRepository
}
