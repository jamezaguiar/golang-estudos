package main

import "fmt"

func main() {
	employees := map[string]float64{
		"José João":      11325.45,
		"Gabriela Silva": 15047.61,
		"Pedro Júnior":   1200.00,
	}
	employees["Rafael Filho"] = 1350.00

	fmt.Println(employees)

	for name, salary := range employees {
		fmt.Printf("%s (Salário: %.2f)\n", name, salary)
	}
}
