// https://leetcode.com/problems/number-of-islands/
// Difficulty: Medium

package main

import "fmt"

func numIslands(grid [][]byte) int {
	return 0
}

func main() {
	grid := [][]byte{
		{'1', '1', '1', '1', '0'},
		{'1', '1', '0', '1', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '0', '0', '0'},
	}
	fmt.Println(numIslands(grid))
}
