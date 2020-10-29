package main

import "fmt"

func main() {
	employeesPerChar := map[string]map[string]float64{
		"G": {
			"Gabriela Silva": 15468.45,
			"Guga Pereira":   14476.54,
		},
		"J": {
			"José João": 7642.15,
		},
		"P": {
			"Pedro Júnior": 78761.85,
		},
	}

	fmt.Println(employeesPerChar)

	for char, employee := range employeesPerChar {
		fmt.Println(char, employee)
	}
}
