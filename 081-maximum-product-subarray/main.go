// https://leetcode.com/problems/maximum-product-subarray/
// Difficulty: Medium

package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(maxProduct([]int{2, 3, -2, 4}))
}

func maxProduct(nums []int) int {
	result := math.MinInt
	idx := len(nums) - 1
	maxProduct, minProduct := 1, 1
	for idx >= 0 {
		maxProduct, minProduct = max(nums[idx], nums[idx]*minProduct, nums[idx]*maxProduct), min(nums[idx], nums[idx]*minProduct, nums[idx]*maxProduct)
		result = max(maxProduct, result)
		idx--
	}
	return result
}

func maxProductBf(nums []int) int {
	result := math.MinInt
	for i := range nums {
		prod := 1
		for j := i; j < len(nums); j++ {
			prod *= nums[j]
			result = max(result, prod)
		}
	}
	return result
}
