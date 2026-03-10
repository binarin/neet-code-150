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
	return maxProductBf(nums)
}

func maxProductBf(nums []int) int {
	result := math.MinInt
	for i := range nums {
		minProd, maxProd := 1, 1
		for j := i; j < len(nums); j++ {
			minProd *= nums[j]
			maxProd *= nums[j]
			if maxProd < minProd {
				minProd, maxProd = maxProd, minProd
			}
			result = max(result, minProd, maxProd)
		}
	}
	return result
}
