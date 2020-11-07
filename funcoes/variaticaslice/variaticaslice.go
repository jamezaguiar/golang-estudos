package main

import "fmt"

func printApproved(approved ...string) {
	fmt.Println("Approved list")
	for i, student := range approved {
		fmt.Printf("%d) %s\n", i+1, student)
	}
}

func main() {
	approved := []string{"Mary", "Peter", "John", "Scott"}

	printApproved(approved...)
}
