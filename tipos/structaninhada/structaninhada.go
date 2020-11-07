package main

import "fmt"

type item struct {
	productID int
	quantity  uint
	price     float64
}

type order struct {
	userID int
	items  []item
}

func (o order) totalValue() float64 {
	total := 0.0

	for _, item := range o.items {
		total += item.price * float64(item.quantity)
	}

	return total
}

func main() {
	order1 := order{
		userID: 1,
		items: []item{
			item{
				productID: 1,
				quantity:  2,
				price:     12.99,
			},
			item{
				productID: 2,
				quantity:  1,
				price:     38.99,
			},
			item{
				productID: 3,
				quantity:  4,
				price:     4.99,
			},
		},
	}

	fmt.Printf("Valor total do pedido: R$%.2f", order1.totalValue())
}
