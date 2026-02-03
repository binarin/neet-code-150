package main

import (
	"fmt"
)

func main() {
	fmt.Println(change(5, []int{1, 2, 5}))
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

func change(amount int, coins []int) int {
	// we are using the fact that amounts are unique, so using the
	// first available coin as a cache key is a valid approach.
	type CacheKey struct {
		a int
		f int
	}
	cache := map[CacheKey]int{}

	var recur func(int, []int) int
	recur = func(amount_r int, coins_r []int) int {
		if amount_r < 0 {
			return 0
		}
		if len(coins_r) == 0 {
			return 0
		}

		cache_key := CacheKey{a: amount_r, f: coins_r[0]}
		if value, is_cached := cache[cache_key]; is_cached {
			return value
		}
		var result int
		if amount_r == 0 {
			result = 1
		} else {
			result = recur(amount_r-coins_r[0], coins_r) + recur(amount_r, coins_r[1:])
		}
		cache[cache_key] = result
		return result
	}

	return recur(amount, coins)
}
