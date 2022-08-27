package order

const (
	DefaultDistance     float64 = 1000
	MinimumFreightValue float64 = 10
)

type Freight struct {
	distance float64
	total    float64
}

func NewFreight() Freight {
	return Freight{distance: DefaultDistance}
}

func (f *Freight) AddItem(item Item, quantity int) {
	f.total += (item.Volume() * f.distance * (item.Density() / 100)) * float64(quantity)
}

func (f *Freight) Total() float64 {
	if f.total > 0 && f.total < MinimumFreightValue {
		return 10
	}
	return f.total
}
