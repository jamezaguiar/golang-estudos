package main

import (
	"fmt"
	"time"
)

func main() {
	// Semelhante ao while (que não existe em Go)
	i := 0
	for i < 10 {
		i++
		fmt.Println(i)
	}

	// For padrão
	for i := 0; i <= 20; i += 2 {
		fmt.Printf("%d \n", i)
	}

	// Estruturas aninhadas
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			fmt.Print("Par ")
		} else {
			fmt.Print("Impar ")
		}
	}

	// Loop infinito
	for {
		fmt.Println("Loop infinito")
		time.Sleep(time.Second)
	}
}
