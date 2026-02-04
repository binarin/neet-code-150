package main

import (
	"fmt"
	"slices"
)

func main() {
	fmt.Println(threeSum([]int{-1, 0, 1, 2, -1, -4}))
}

func threeSum(nums []int) [][]int {
	if len(nums) < 3 {
		return [][]int{}
	}

	lookup := map[int][]int{}
	for i, n := range nums {
		lookup[n] = append(lookup[n], i)
	}

	// 2 elements are enough to identify triplet
	triplets := map[[3]int]bool{}

	pack := func(i int, j int, k int) [3]int {
		as_slice := []int{nums[i], nums[j], nums[k]}
		slices.Sort(as_slice)
		return [3]int{as_slice[0], as_slice[1], as_slice[2]}
	}

	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			third_number := -(nums[i] + nums[j])
			found := false
			var k int
			for _, nums_idx := range lookup[third_number] {
				if nums_idx == i || nums_idx == j {
					continue
				}
				found = true
				k = nums_idx
				break
			}
			if found {
				triplets[pack(i, j, k)] = true
			}
		}
	}

	result := [][]int{}
	for k := range triplets {
		result = append(result, k[:])
	}
	return result
}
