// https://leetcode.com/problems/number-of-connected-components-in-an-undirected-graph/
// Medium

package main

import "fmt"

func main() {
	// Example 1: n = 5, edges = [[0,1],[1,2],[3,4]], Output: 2
	n := 5
	edges := [][]int{{0, 1}, {1, 2}, {3, 4}}
	fmt.Println(countComponents(n, edges))
}

func countComponents(n int, edges [][]int) int {
	return 0
}
