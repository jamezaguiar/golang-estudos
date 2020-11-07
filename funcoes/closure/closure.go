package main

import "fmt"

func closure() func() {
	x := 10
	var function = func() {
		fmt.Println(x)
	}
	return function
}

func main() {
	x := 20
	fmt.Println(x)

	printX := closure()
	printX()
	/*
		A função retornada pela func closure respeita o escopo
		em que foi declarada, imprimindo assim,
		o valor de x de dentro da função closure.
	*/
}
