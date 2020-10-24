package main

import "fmt"

func main() {
	numbers := [...]int{1, 2, 3, 4, 5} // O compilador contará o tamanho do array

	for i, number := range numbers {
		// Usando o range, teremos como retorno o índice e o elemento
		fmt.Printf("%d) %d\n", i+1, number)
	}
}
