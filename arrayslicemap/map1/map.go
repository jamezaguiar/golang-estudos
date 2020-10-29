package main

import "fmt"

func main() {
	// var approved map[int]string // Não funciona
	approved := make(map[int]string) // Mapas devem ser inicializados

	approved[121310451] = "Mary"
	approved[274243354] = "John"
	approved[413248476] = "Anne"

	fmt.Println(approved)

	for cpf, name := range approved {
		fmt.Printf("%s (CPF: %d)\n", name, cpf)
	}

	delete(approved, 413248476) // Deleta Anne do mapa
}
