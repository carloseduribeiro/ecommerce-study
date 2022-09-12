package simulate_freight

type ItemInput struct {
	IdItem   int
	Quantity int
}
type SimulateFreightInput struct {
	OrderItems []ItemInput
}

type SimulateFreightOutput struct {
	total float64
}

func (o SimulateFreightOutput) Total() float64 {
	return o.total
}
