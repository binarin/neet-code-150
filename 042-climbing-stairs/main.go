package main

import "fmt"

func main() {
	fmt.Println(climbStairs(2))
}

func climbStairs(n int) int {
	memo := make([]int, n+1)
	var recur func(int) int
	recur = func(n int) int {
		if n < 0 {
			return 0
		}
		if n == 0 {
			return 1
		}
		if memo[n] == 0 {
			memo[n] = recur(n-1) + recur(n-2)
		}
		return memo[n]
	}
	return recur(n)
}
