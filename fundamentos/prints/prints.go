package main

import "fmt"

func main() {
	fmt.Print("Mesma")
	fmt.Print(" Linha.")

	fmt.Println("\nOutra")
	fmt.Println("linha.")

	x := 3.141516
	// fmt.Println("O valor de x é: " + x)
	/*
		Não é possível concatenar uma variável como x numa string
		Para fazer isso, temos a solução abaixo:
	*/
	xString := fmt.Sprint(x)
	fmt.Println("O valor de x é: " + xString)

	// Mas é possível também fazer apenas o seguinte:
	fmt.Println("O valor de x é:", x)

	// Ou usar o Printf
	fmt.Printf("O valor de x é: %.2f", x)

	a := 1
	b := 1.999
	c := false
	d := "Hi"

	fmt.Printf("\n%d %f %1.f %t %s", a, b, b, c, d)
}
