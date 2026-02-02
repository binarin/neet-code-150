package main

import "fmt"

func main() {
	fmt.Println(twoSums([]int{2, 7, 11, 15}, 9))
}

func twoSums(nums []int, target int) []int {
	for i := 0; i < len(nums) - 1; i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] + nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return nil
}
