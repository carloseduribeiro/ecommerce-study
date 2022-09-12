package simulate_freight

import (
	"github.com/ecommerce-study/internal/infra/repository/memory"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSimulateFreight(t *testing.T) {
	t.Run("should simulate the freight of an order", func(t *testing.T) {
		// given
		itemRepository := memory.NewItemRepository()
		usecase := NewSimulateFreight(&itemRepository)
		input := SimulateFreightInput{
			OrderItems: []ItemInput{
				{IdItem: 1, Quantity: 1},
				{IdItem: 2, Quantity: 1},
				{IdItem: 3, Quantity: 3},
			},
		}
		// when
		output, _ := usecase.execute(input)
		expected := 260.0
		// when
		assert.Equal(t, expected, output.Total())
	})
}
