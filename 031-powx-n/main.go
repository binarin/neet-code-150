package main

import (
	"fmt"
)

func main() {
	fmt.Println(myPow(2.00000, 10))
}

func myPow(x float64, n int) float64 {
	if n == 0 {
		return 1
	}
	if n == 1 {
		return x
	}
	if n < 0 {
		return 1 / myPow(x, -n)
	}
	if n%2 == 0 {
		sub := myPow(x, n/2)
		return sub * sub
	} else {
		return x * myPow(x, n-1)
	}
}
