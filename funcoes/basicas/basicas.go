package main

import "fmt"

func f1() {
	fmt.Println("F1")
}

func f2(p1 string, p2 string) {
	fmt.Printf("F2: %s %s\n", p1, p2)
}

func f3() string {
	return "F3"
}

func f4(p1, p2 string) string {
	return fmt.Sprintf("F4: %s %s", p1, p2)
}

func f5() (string, string) {
	return "First return", "Second return"
}

func main() {
	f1()
	f2("O", "lá")
	f3Res := f3()
	fmt.Println(f3Res)
	f4Res := f4("O", "lá")
	fmt.Println(f4Res)
	f5Res1, f5Res2 := f5()
	fmt.Println(f5Res1, f5Res2)
}
