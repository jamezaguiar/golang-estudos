package main

import "fmt"

func getApprovedValue(number int) int {
	/*
		A execução desta linha acontecerá no último momento possível,
		(antes do return) graças ao defer.
	*/
	defer fmt.Println("End!")
	if number > 5000 {
		fmt.Println("High value!")
		return 5000
	}
	fmt.Println("Low value!")
	return number
}

func main() {
	fmt.Println(getApprovedValue(5001))
}
