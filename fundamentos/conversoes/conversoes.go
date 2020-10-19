package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func main() {
	x := 2.4
	y := 2
	fmt.Println(x / float64(y))

	nota := 6.9
	notaFinal := int(nota)
	fmt.Println(notaFinal)

	// Cuidados com a conversão int > string
	fmt.Println("Teste " + string(97))       // Isso retornará um caractere da tabela unicode
	fmt.Println("Teste " + strconv.Itoa(97)) // Isso converterá 97 em string

	// string > int
	num, err := strconv.Atoi("123") // Retorna número e um erro
	fmt.Println(num)
	fmt.Println(err)

	// Booleans
	b, _ := strconv.ParseBool("false")
	fmt.Println(reflect.TypeOf(b))
}
