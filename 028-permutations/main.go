package main

import (
	"fmt"
)

func main() {
	fmt.Println(permute([]int{1, 2, 3}))
}

func permute(nums []int) [][]int {
	if len(nums) == 0 {
		return [][]int{[]int{}}
	}
	result := [][]int{}
	for _, sub_perm := range permute(nums[1:]) {
		for insert_point := range nums {
			perm := make([]int, len(nums))
			if insert_point <= len(sub_perm) {
				copy(perm, sub_perm[0:insert_point])
			}
			perm[insert_point] = nums[0]
			if insert_point < len(sub_perm) {
				copy(perm[insert_point+1:], sub_perm[insert_point:])
			}
			result = append(result, perm)
		}
	}
	return result
}
