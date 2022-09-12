package simulate_freight

import (
	"errors"
	"github.com/ecommerce-study/internal/domain/entity"
	"github.com/ecommerce-study/internal/domain/repository"
)

type SimulateFreight struct {
	itemRepository repository.ItemRepository
}

func NewSimulateFreight(itemRepository repository.ItemRepository) SimulateFreight {
	return SimulateFreight{
		itemRepository: itemRepository,
	}
}

func (s *SimulateFreight) execute(input SimulateFreightInput) (*SimulateFreightOutput, error) {
	freight := entity.NewFreight()
	for _, orderItemInput := range input.OrderItems {
		item, err := s.itemRepository.GetById(orderItemInput.IdItem)
		if err != nil {
			return nil, errors.New("item not found")
		}
		freight.AddItem(*item, orderItemInput.Quantity)
	}
	output := &SimulateFreightOutput{total: freight.Total()}
	return output, nil
}
