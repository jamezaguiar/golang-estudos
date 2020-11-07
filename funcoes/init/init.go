package main

import "fmt"

/*
	Por convenção, dentro de programas Go, podemos ter a função init.
	Ela será executada antes mesmo da função main, mesmo sem ser chamada.
*/

func init() {
	fmt.Println("Initializing")
}

func main() {
	fmt.Println("Main")
}
