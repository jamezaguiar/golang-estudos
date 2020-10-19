package main

import "fmt"

func main() {
	i := 1

	// Criamos o ponteiro p, que conterá um endereço de memória
	var p *int = nil // nil == null

	p = &i // Pegamos o endereço de memória da variável i
	*p++
	i++

	fmt.Printf(`Endereço de memória %v, contém o valor %v,
que é igual a variável i, contendo %v
e está no endereço de memória %v
	`, p, *p, i, &i)
}
