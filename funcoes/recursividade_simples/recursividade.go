package main

import "fmt"

func factorial(number uint) uint {
	switch {
	case number == 0:
		return 1
	default:
		return number * factorial(number-1)
	}
}

func main() {
	result := factorial(5)

	fmt.Println(result)
}
