package main

import (
	"fmt"
	"math"
	"reflect"
)

func main() {
	// Inteiros
	fmt.Println(reflect.TypeOf(1))

	/*
		Números sem sinal (apenas positivos) podem ser:
		uint8
		uint16
		uint32
		uint64
	*/

	// u vem de unsigned (sem sinal)
	var num1 uint8 = 5
	fmt.Println("O valor da variável num1 é do tipo", reflect.TypeOf(num1))

	/*
		Números com sinal (podendo ser positivos ou negativos) podem ser:
		int8
		int16
		int32
		int64
	*/

	var num2 int64 = math.MaxInt64

	fmt.Printf("O valor máximo de um %s é %d\n", reflect.TypeOf(num2), num2)

	var unicodeTable rune = 'a' // Representa um mapeamento da tabela unicode
	fmt.Println("O rune é", reflect.TypeOf(unicodeTable))
	fmt.Println(unicodeTable)

	// Números reais
	var num3 float32 = 49.99
	fmt.Println("O tipo de num3 é", reflect.TypeOf(num3))

	// Booleans
	bo := true
	fmt.Println("O tipo de bo é", reflect.TypeOf(bo))

	// String
	s1 := "Olá, meu nome é Jamerson Aguiar"
	fmt.Println("O tipo da variável s1 é", reflect.TypeOf(s1))
	fmt.Println("O tamanho da string s1 é", len(s1))
	// Múltiplas linhas
	s2 := `Olá,
	meu
	nome
	é
	Jamerson
	Aguiar`
	fmt.Println("O tamanho da string s2 é", len(s2))

	// Char? Não existe o tipo char no Go
	var a = 'a'
	fmt.Println("O tipo da variável a é", reflect.TypeOf(a))
	/*
		Semelhante ao tipo rune, teremos uma referência na tabela unicode
		ao declarar um caractere dentro de aspas simples.
	*/
}
