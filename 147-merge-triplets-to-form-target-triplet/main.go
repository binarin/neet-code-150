// https://leetcode.com/problems/merge-triplets-to-form-target-triplet/description/?envType=problem-list-v2&envId=plakya4j
// Difficulty: Medium

package main

import "fmt"

func mergeTriplets(triplets [][]int, target []int) bool {
	return false
}

func main() {
	triplets := [][]int{{2, 5, 3}, {1, 8, 4}, {1, 7, 5}}
	target := []int{2, 7, 5}
	fmt.Println(mergeTriplets(triplets, target))
}
