package main

import (
	"fmt"
)

func main() {
	fmt.Println(numDistinct("rabbbit", "rabbit"))
}

func numDistinct(s string, t string) int {
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
