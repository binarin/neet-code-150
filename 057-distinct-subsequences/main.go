package main

import (
	"fmt"
)

func main() {
	fmt.Println(numDistinct("rabbbit", "rabbit"))
}

func numDistinct(s, t string) int {
	return numDistinctDyn(s, t)
}

func numDistinctDyn(s string, t string) int {
	cache := make([]int, len(t)+1)
	cache[len(t)] = 1
	for sidx := len(s) - 1; sidx >= 0; sidx-- {
		tidxStart := min(len(t)-1, sidx)
		diagCarry := cache[tidxStart+1]
		// fmt.Printf("S: %d - %v\n", sidx, cache[:tidxStart+2])
		for tidx := tidxStart; tidx >= 0; tidx-- {
			thisCount := cache[tidx]
			if s[sidx] == t[tidx] {
				thisCount += diagCarry
			}
			// fmt.Printf("tidx %d, old cache %d, this count %d, diagCarry %d\n", tidx, cache[tidx], thisCount, diagCarry)
			cache[tidx], diagCarry = thisCount, cache[tidx]
		}
	}
	return cache[0]
}

func numDistinctMem(s string, t string) int {
	cache := map[[2]int]int{}

	var recur func(int, int) int
	recur = func(sidx, tidx int) (result int) {
		cacheKey := [2]int{sidx, tidx}
		if val, ok := cache[cacheKey]; ok {
			return val
		}
		defer (func() {
			cache[cacheKey] = result
		})()

		if len(t) == tidx {
			return 1
		}
		if len(s) == sidx {
			return 0
		}
		count := recur(sidx+1, tidx)
		if s[sidx] == t[tidx] {
			count += recur(sidx+1, tidx+1)
		}
		return count
	}

	return recur(0, 0)
}
