package main

import (
	"fmt"
	"slices"
)

func main() {
	fmt.Println(threeSum([]int{-1, 0, 1, 2, -1, -4}))
}

func threeSum(nums []int) [][]int {
	triplets := map[[3]int]bool{}
	if len(nums) < 3 {
		return [][]int{}
	}
	for i := 0; i < len(nums)-2; i++ {
		for j := i + 1; j < len(nums)-1; j++ {
			for k := j + 1; k < len(nums); k++ {
				if nums[i]+nums[j]+nums[k] == 0 {
					triplet_slice := []int{nums[i], nums[j], nums[k]}
					slices.Sort(triplet_slice)
					triplet_array := [3]int{triplet_slice[0], triplet_slice[1], triplet_slice[2]}
					triplets[triplet_array] = true
				}
			}
		}
	}
	result := [][]int{}
	for k := range triplets {
		result = append(result, k[:])
	}
	return result
}
