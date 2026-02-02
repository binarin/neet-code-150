package main

import "fmt"

func main() {
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 9))
	fmt.Println(twoSum([]int{3, 2, 4}, 6))
	fmt.Println(twoSum([]int{3, 3}, 6))
}

func twoSumBruteforce(nums []int, target int) []int {
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return nil
}

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i, n := range nums {
		m[n] = i
	}
	for i, n := range nums {
		need := target - n
		if val, ok := m[need]; ok && m[need] != i {
			return []int{i, val}
		}
	}
	return nil
}
