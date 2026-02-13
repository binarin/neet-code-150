package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(minDistance("horse", "ros"))
}

func minDistance(word1 string, word2 string) int {
	memo := map[[2]string]int{}
	var recur func(string, string) int
	recur = func(word1, word2 string) int {
		if val, ok := memo[[2]string{word1, word2}]; ok {
			return val
		}
		if len(word2) == 0 {
			return len(word1)
		}
		if len(word1) == 0 {
			return len(word2)
		}

		best := math.MaxInt
		if word1[0] == word2[0] {
			best = min(best, recur(word1[1:], word2[1:]))
		}
		// replace
		best = min(best, 1+recur(word1[1:], word2[1:]))
		// delete
		best = min(best, 1+recur(word1[1:], word2))
		best = min(best, 1+recur(word1, word2[1:]))
		memo[[2]string{word1, word2}] = best
		return best
	}
	return recur(word1, word2)
}
