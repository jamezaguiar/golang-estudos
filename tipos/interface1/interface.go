package main

import "fmt"

type printable interface {
	toString() string
}

type user struct {
	firstName string
	lastName  string
}

type product struct {
	name  string
	price float64
}

// Em Go, interfaces são implementadas implicitamente.
func (u user) toString() string {
	return u.firstName + " " + u.lastName
}

func (p product) toString() string {
	return fmt.Sprintf("%s - R$%.2f", p.name, p.price)
}

func print(p printable) {
	fmt.Println(p.toString())
}

func main() {
	var anyThing printable = user{"Jamerson", "Aguiar"}
	fmt.Println(anyThing.toString())
	print(anyThing)

	anyThing = product{"Jeans", 179.90}
	fmt.Println(anyThing.toString())
	print(anyThing)

	p2 := product{"Shirt", 79.90}
	fmt.Println(p2.toString())
	print(p2)
}
