package main

import (
	"fmt"
)

func main() {
	fmt.Println(uniquePaths(3, 7))
}

func uniquePaths(m int, n int) int {
	if m == 1 || n == 1 {
		return 1
	}
	return uniquePaths(m, n-1) + uniquePaths(m-1, n)
}
