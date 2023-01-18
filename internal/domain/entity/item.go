package entity

import "fmt"

const invalidParamErrStr = "invalid parameter value: %s is less then zero"

type ItemOption func(item *Item) error

func WithDimensions(width, height, length float64) ItemOption {
	return func(item *Item) error {
		if width < 0 {
			return fmt.Errorf(invalidParamErrStr, "width")
		}
		if height < 0 {
			return fmt.Errorf(invalidParamErrStr, "height")
		}
		if length < 0 {
			return fmt.Errorf(invalidParamErrStr, "length")
		}
		item.dimension = NewDimension(width, height, length)
		return nil
	}
}

func WithDimension(dimension Dimension) ItemOption {
	return func(item *Item) error {
		item.dimension = dimension
		return nil
	}
}

func WithWeight(weight float64) ItemOption {
	return func(item *Item) error {
		if weight < 0 {
			return fmt.Errorf(invalidParamErrStr, "weight")
		}
		item.weight = weight
		return nil
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

func (i Item) Category() string {
	return i.category
}

func (i Item) Description() string {
	return i.description
}

func (i Item) Price() float64 {
	return i.price
}
