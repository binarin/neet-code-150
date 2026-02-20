package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(subsets([]int{1, 2, 3}))
}

func pow(a, b int) int {
	return int(math.Pow(float64(a), float64(b)))
}

func subsets(nums []int) [][]int {
	total := pow(2, len(nums))
	result := make([][]int, 0, total)
	for num := range total {
		idx := 0
		subset := []int{}
		for num > 0 {
			if num%2 == 1 {
				subset = append(subset, nums[idx])
			}
			idx++
			num /= 2
		}
		result = append(result, subset)
	}
	return result
}
