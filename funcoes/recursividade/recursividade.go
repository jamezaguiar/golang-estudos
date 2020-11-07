package main

import "fmt"

func factorial(number int) (int, error) {
	switch {
	case number < 0:
		return -1, fmt.Errorf("Invalid number: %d", number)
	case number == 0:
		return 1, nil
	default:
		previousFactorial, _ := factorial(number - 1)
		return number * previousFactorial, nil
	}
}

func main() {
	result, err := factorial(-5)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}
}
