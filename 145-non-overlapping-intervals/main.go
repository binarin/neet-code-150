// https://leetcode.com/problems/non-overlapping-intervals/
// Difficulty: Medium

package main

import "fmt"

func eraseOverlapIntervals(intervals [][]int) int {
	return 0
}

func main() {
	// Example 1: intervals = [[1,2],[2,3],[3,4],[1,3]] -> Output: 1
	intervals := [][]int{{1, 2}, {2, 3}, {3, 4}, {1, 3}}
	fmt.Println(eraseOverlapIntervals(intervals))
}
