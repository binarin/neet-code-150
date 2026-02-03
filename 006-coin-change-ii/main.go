package main

import (
	"fmt"
)

func main() {
	fmt.Println("hello")
}

func changeBf(amount int, coins []int) int {
	// let’s start with recursive solution, easier to make it correct
	// can rewrite later if stack explodes
	if amount < 0 {
		return 0
	}
	if amount == 0 {
		return 1
	}
	if len(coins) == 0 {
		return 0
	}

	return changeBf(amount-coins[0], coins) + changeBf(amount, coins[1:])
}
