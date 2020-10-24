package main

import "fmt"

func gradeForConcept(grade float64) string {
	switch {
	case grade >= 9 && grade <= 10:
		return "A"
	case grade >= 8 && grade <= 9:
		return "B"
	case grade >= 5 && grade <= 8:
		return "C"
	case grade >= 3 && grade <= 5:
		return "D"
	default:
		return "E"
	}
}

func main() {
	fmt.Println(gradeForConcept(8.9))
}
