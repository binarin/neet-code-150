// https://leetcode.com/problems/graph-valid-tree/
// Difficulty: Medium

package main

import "fmt"

func main() {
	// Example 1: n = 5, edges = [[0,1],[0,2],[0,3],[1,4]] -> true
	n := 5
	edges := [][]int{{0, 1}, {0, 2}, {0, 3}, {1, 4}}
	fmt.Println(validTree(n, edges))
}

func validTree(n int, edges [][]int) bool {
	return false
}
