// https://leetcode.com/problems/two-sum-ii-input-array-is-sorted/
// Difficulty: Medium

package main

import "fmt"

func twoSum(numbers []int, target int) []int {
	leftIdx, rightIdx := 0, len(numbers)-1
	for {
		for numbers[leftIdx]+numbers[rightIdx] > target {
			fmt.Println(numbers[leftIdx], rightIdx, numbers[rightIdx])
			rightIdx--
		}
		if numbers[leftIdx]+numbers[rightIdx] == target {
			return []int{leftIdx + 1, rightIdx + 1}
		}
		leftIdx++
		for numbers[leftIdx]+numbers[rightIdx] > target {
			rightIdx--
		}
		if numbers[leftIdx]+numbers[rightIdx] == target {
			return []int{leftIdx + 1, rightIdx + 1}
		}
	}
	return nil
}

func main() {
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 9))
}
