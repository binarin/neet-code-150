package main

import (
	"fmt"
	"slices"
)

func main() {
	fmt.Println(trapBf([]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}))
}

func trap(height []int) int {
	if len(height) < 3 {
		return 0
	}
	max_from_left := make([]int, len(height))
	max_from_right := make([]int, len(height))
	for i, so_far := 0, 0; i < len(max_from_left); i++ {
		max_from_left[i] = so_far
		so_far = max(so_far, height[i])
	}
	for i, so_far := len(max_from_right)-1, 0; i >= 0; i-- {
		max_from_right[i] = so_far
		so_far = max(so_far, height[i])
	}
	result := 0
	for i := 1; i < len(height)-1; i++ {
		fill_level := min(max_from_left[i], max_from_right[i])
		if height[i] < fill_level {
			result += fill_level - height[i]
		}
	}
	return result
}

func trapBf(height []int) int {
	heightDomain := slices.Max(height)
	result := 0
	for h := 1; h <= heightDomain; h++ {
		max_left := height[0]
		for i := 1; i < len(height)-1; i++ {
			cur_max_left := max_left
			max_left = max(height[i], max_left)
			if height[i] >= h {
				continue
			}
			if cur_max_left < h {
				continue
			}
			for j := i + 1; j < len(height); j++ {
				if height[j] >= h {
					result++
					break
				}
			}
		}
	}
	return result
}
