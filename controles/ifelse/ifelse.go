package main

import "fmt"

func printResult(grade float64) {
	if grade >= 7 {
		fmt.Println("Aprovado com nota", grade)
	} else {
		fmt.Println("Reprovado com nota", grade)
	}
}

func main() {
	printResult(6.3)
}
