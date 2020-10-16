// Programas executáveis iniciam pelo pacote main
package main

/*
	Os códigos em Go são organizados em pacotes
	que podem ser declarados com um ou vários imports
*/

import "fmt"

/*
	ou:
	import (
		"fmt"
	)
*/

// A porta de entrada de um programa Go é a função main()
func main() {
	fmt.Println("Hello, world!")
}
