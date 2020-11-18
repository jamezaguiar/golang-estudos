package main

import (
	"fmt"
	"time"
)

func routine(ch chan int) {
	time.Sleep(1 * time.Second)
	ch <- 1
	fmt.Println("Depois que foi lido...")
}

func main() {
	ch := make(chan int)

	go routine(ch)

	fmt.Println(<-ch)
	fmt.Println("Dado recebido e lido no canal")
	fmt.Println(<-ch) // Deadlock. Não há mais dados a serem lidos no canal
	fmt.Println("Fim")
}
