// https://leetcode.com/problems/network-delay-time/
// Difficulty: Medium

package main

import "fmt"

func main() {
	// Example 1: times = [[2,1,1],[2,3,1],[3,4,1]], n = 4, k = 2 -> Output: 2
	times := [][]int{{2, 1, 1}, {2, 3, 1}, {3, 4, 1}}
	n := 4
	k := 2
	fmt.Println(networkDelayTime(times, n, k))
}

func networkDelayTime(times [][]int, n int, k int) int {
	return 0
}
