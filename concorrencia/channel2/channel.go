package main

import (
	"fmt"
	"time"
)

// Channel (canal): Forma de comunicação entre as goroutines
// É um tipo da linguagem Go

func twoThreeFourMult(base int, ch chan int) {
	time.Sleep(time.Second)
	ch <- 2 * base // Enviando dados para o canal

	time.Sleep(time.Second)
	ch <- 3 * base // Enviando dados para o canal

	time.Sleep(3 * time.Second)
	ch <- 4 * base // Enviando dados para o canal
}

// A goroutine causa um "assincronismo" na execução
// Enquanto o canal causa um "sincronismo"...
func main() {
	ch := make(chan int)

	go twoThreeFourMult(2, ch)
	fmt.Println("Go Routine começou")

	a, b := <-ch, <-ch // A execução ficará nesta linha até que os dados sejam recebidos pelo canal
	fmt.Println("Dados recebidos")
	fmt.Println(a, b)

	fmt.Printf("Último dado chegando pelo canal: %d", <-ch)
}
