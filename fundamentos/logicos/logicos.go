package main

import "fmt"

func buy(work1 bool, work2 bool) (bool, bool, bool) {
	buyTV50 := work1 && work2
	buyTV32 := work1 != work2
	buyIceCream := work1 || work2

	return buyTV50, buyTV32, buyIceCream
}

func main() {
	tv50, tv32, iceCream := buy(true, true)
	fmt.Printf("TV50: %t, TV32: %t, Ice Cream: %t, Healthy: %t",
		tv50, tv32, iceCream, !iceCream)
}
