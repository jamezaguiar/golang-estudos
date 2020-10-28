package main

import "fmt"

func main() {
	s := make([]int, 10)
	s[9] = 12
	fmt.Println(s)

	s = make([]int, 10, 20)
	/*
		No caso acima, o Go criará um slice de 10 posições,
		mas internamente, ele criará um array de 20 posições,
		e o slice referenciará este array.
	*/
	fmt.Println(s, len(s), cap(s)) // Slice, tamanho do slice, capacidade (tamanho do array interno)

	s = append(s, 1, 2, 3, 4, 5, 6, 7, 8, 9, 0)
	fmt.Println(s, len(s), cap(s))

	s = append(s, 1) // O array interno crescerá dinâmicamente ao ultrapassar sua capacidade
	fmt.Println(s, len(s), cap(s))
}
