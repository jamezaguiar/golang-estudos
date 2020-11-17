package main

import "fmt"

func main() {
	ch := make(chan int, 1)

	ch <- 2           // Envio de dados para o canal (escrita)
	fmt.Println(<-ch) // Recebendo dados do canal (leitura)
}
