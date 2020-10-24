package main

import "fmt"

func main() {
	/*
		Arrays são estruturas homogêneas (mesmo tipo)
		e estática (tamanho fixo)
	*/

	var grades [3]float64
	grades[0], grades[1], grades[2] = 7.8, 4.3, 9.1
	fmt.Println(grades)

	total := 0.0
	for i := 0; i < len(grades); i++ {
		total += grades[i]
	}
	average := total / float64(len(grades))
	fmt.Printf("%.2f\n", average)
}
