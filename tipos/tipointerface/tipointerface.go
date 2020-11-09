package main

import "fmt"

/*
	Apesar de Go ser uma linguagem fortemente tipada, podemos
	usar tipos genéricos através de interfaces.
*/

type course struct {
	name string
}

func main() {
	var anyThing interface{}
	fmt.Println(anyThing)

	anyThing = 3
	fmt.Println(anyThing)

	type dynamic interface{}

	var anyThing2 dynamic = "Hi"
	fmt.Println(anyThing2)

	anyThing2 = true
	fmt.Println(anyThing2)

	anyThing2 = course{"Golang: Explorando a linguagem do Google."}
	fmt.Println(anyThing2)
}
