package main

import "fmt"

func inc1(n int) {
	n++
}

func inc2(n *int) {
	*n++
}

func main() {
	n := 1

	inc1(n) // A função inc1 recebe uma cópia de n
	fmt.Println(n)

	inc2(&n) // A função inc2 recebe o endereço de memória de n, e trabalha diretamente no valor
	fmt.Println(n)
}
