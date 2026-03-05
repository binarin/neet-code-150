// https://leetcode.com/problems/k-closest-points-to-origin/description/?envType=problem-list-v2&envId=plakya4j
// Difficulty: Medium

package main

import "fmt"

func main() {
	// Example 1: points = [[1,3],[-2,2]], k = 1
	// Expected output: [[-2,2]]
	points := [][]int{{1, 3}, {-2, 2}}
	k := 1
	result := kClosest(points, k)
	fmt.Println(result)
}

func kClosest(points [][]int, k int) [][]int {
	return nil
}
