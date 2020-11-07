package main

import "fmt"

type grade float64

func (g grade) between(start, end float64) bool {
	return float64(g) >= start && float64(g) <= end
}

func gradeForConcept(g grade) string {
	if g.between(9.0, 10.0) {
		return "A"
	} else if g.between(7.0, 8.99) {
		return "B"
	} else if g.between(5.0, 6.99) {
		return "C"
	} else if g.between(3.0, 4.99) {
		return "D"
	} else {
		return "E"
	}
}

func main() {
	fmt.Println(gradeForConcept(9.8))
	fmt.Println(gradeForConcept(6.9))
	fmt.Println(gradeForConcept(2.3))
	fmt.Println(gradeForConcept(4.7))
}
