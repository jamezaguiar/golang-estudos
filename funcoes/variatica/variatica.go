package main

import "fmt"

func average(numbers ...float64) float64 {
	total := 0.0

	for _, number := range numbers {
		total += number
	}

	return total / float64(len(numbers))
}

func main() {
	fmt.Printf("Média: %.2f\n", average(8.4, 9.2, 10, 5.3))
}
