package entity

type ItemOption func(item *Item)

func WithDimensions(width, height, length float64) ItemOption {
	return func(item *Item) {
		item.dimension = NewDimension(width, height, length)
	}
}

func WithDimension(dimension Dimension) ItemOption {
	return func(item *Item) {
		item.dimension = dimension
	}
}

func WithWeight(weight float64) ItemOption {
	return func(item *Item) {
		item.weight = weight
	}
}

type Item struct {
	id          int
	category    string
	description string
	price       float64
	weight      float64
	dimension   Dimension
}

func NewItem(id int, category, description string, price float64, opts ...ItemOption) Item {
	item := &Item{
		id:          id,
		category:    category,
		description: description,
		price:       price,
		dimension:   Dimension{},
		weight:      0.0,
	}
	for _, opt := range opts {
		opt(item)
	}
	return *item
}

func (i Item) Volume() float64 {
	return i.dimension.Volume()
}

func (i Item) Density() float64 {
	if i.weight > 0 {
		return i.weight / i.Volume()
	}
	return 0
}

func (i Item) Id() int {
	return i.id
}
