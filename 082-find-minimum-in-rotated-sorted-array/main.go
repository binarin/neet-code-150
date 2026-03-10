// https://leetcode.com/problems/find-minimum-in-rotated-sorted-array/
// Difficulty: Medium

package main

import (
	"fmt"
)

func findMin(nums []int) int {
	left, right := 0, len(nums)-1
	for left < right {
		if nums[left] <= nums[right] {
			// sorted sub-array
			return nums[left]
		} else {
			middle := left + (right-left)/2
			if nums[left] <= nums[middle] {
				left = middle + 1 // middle is left biased, we’ve already checked it, and need to downsize
			} else {
				right = middle
			}
		}
	}

	return nums[left]
}

func main() {
	fmt.Println(findMin([]int{3, 4, 5, 1, 2}))
}
