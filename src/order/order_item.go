package order

type OrderItem struct {
	item     Item
	quantity int
	price    float64
}

func NewOrderItem(item Item, quantity int, price float64) OrderItem {
	return OrderItem{item, quantity, price}
}

func (o OrderItem) Total() (total float64) {
	return o.price * float64(o.quantity)
}
