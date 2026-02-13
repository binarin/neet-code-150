package main

import (
	"fmt"
)

func main() {
	fmt.Println(uniquePaths(3, 7))
}

func uniquePaths(m int, n int) int {
	memo := [][]int{}
	for range m {
		memo = append(memo, make([]int, n))
	}
	var recur func(int, int) int
	recur = func(m, n int) int {
		if m == 1 || n == 1 {
			return 1
		}
		if already := memo[m-1][n-1]; already == 0 {
			memo[m-1][n-1] = recur(m, n-1) + recur(m-1, n)
		}
		return memo[m-1][n-1]
	}

	return recur(m, n)
}
