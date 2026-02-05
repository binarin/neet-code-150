package main

import (
	"fmt"
)

func main() {
	fmt.Println(search([]int{4, 5, 6, 7, 8, 9, 1, 2, 3}, 1))
}

func search(nums []int, target int) int {
	// i'm thinking what would be easier/possible - unrotate first(virtually), or search immediately
	// it's clearly a modified binary search
	if len(nums) == 0 {
		return -1
	}

	left := 0
	right := len(nums) - 1
	for {
		// 0 -> 1 -> 2
		center := left + (right-left+1)/2
		if nums[center] == target {
			return center
		}

		var left_slice []int
		if left <= center {
			left_slice = nums[left:center]
		}
		if len(left_slice) > 0 {
			b, e := left_slice[0], left_slice[len(left_slice)-1]
			well_formed := b <= e
			target_in_slice := false
			if well_formed && b <= target && target <= e {
				target_in_slice = true
			}
			if !well_formed && (b <= target || target <= e) {
				target_in_slice = true
			}
			if target_in_slice {
				right = center - 1
				continue
			}
		}

		var right_slice []int
		if center+1 <= right {
			right_slice = nums[center+1 : right+1]
		}
		if len(right_slice) > 0 {
			b, e := right_slice[0], right_slice[len(right_slice)-1]
			well_formed := b <= e
			target_in_slice := false
			if well_formed && b <= target && target <= e {
				target_in_slice = true
			}
			if !well_formed && (b <= target || target <= e) {
				target_in_slice = true
			}
			if target_in_slice {
				left = center + 1
				continue
			}
		}

		return -1
	}
}
