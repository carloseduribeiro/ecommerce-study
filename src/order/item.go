package order

type Item struct {
	id          int
	category    string
	description string
	price       float64
}

func NewItem(id int, category, description string, price float64) Item {
	return Item{id, category, description, price}
}
