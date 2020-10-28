package main

import (
	"fmt"
)

func main() {
	arr := [5]int{1, 2, 3, 4, 5} // Array
	// Slice define um pedaço de um array.
	sli := arr[1:3] // Este slice irá do índice 1 ao índice 3, sem incluir o índice 3
	sli2 := arr[:2] // Este slice irá do início do array até o índice 2, sem incluir o índice 2
	fmt.Println(arr, sli, sli2)

}
