package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(maxSubArray([]int{-1}))
}

func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	best_left_drop := make([]int, len(nums)+1)
	total := 0

	for i := range nums {
		total += nums[i]
		best_left_drop[i+1] = min(best_left_drop[i], total)
	}

	best := math.MinInt

	right_idx := len(nums)
	right_drop := 0

	for {
		candidate := total - right_drop - best_left_drop[right_idx-1]
		best = max(candidate, best)

		right_idx--
		if right_idx <= 0 {
			break
		}
		right_drop += nums[right_idx]
	}

	return best
}

func maxSubArrayBf(nums []int) int {
	result := math.MinInt
	for i := 0; i < len(nums); i++ {
		sum := 0
		for j := i; j < len(nums); j++ {
			sum += nums[j]
			result = max(result, sum)
		}
	}
	return result
}
