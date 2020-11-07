package main

import "fmt"

/*
	Criamos aqui um struct, uma estrutura que declara um tipo
	Parecido com os objetos de OO.
	Podemos adicionar métodos a esta estrutura, através de uma função
	com receiver (receptor)
*/

type product struct {
	name     string
	price    float64
	discount float64
}

func (p product) discountPrice() float64 {
	return p.price * (1 - p.discount)
}

func main() {
	var product1 product

	product1 = product{
		name:     "Pencil",
		price:    1.99,
		discount: 0.05,
	}

	product2 := product{"Notebook", 4200.00, 0.10}

	fmt.Println(product1, product1.discountPrice())
	fmt.Println(product2, product2.discountPrice())
}
