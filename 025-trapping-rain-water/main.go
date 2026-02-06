package main

import (
	"fmt"
	"slices"
)

func main() {
	fmt.Println(trapBf([]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}))
}

func trap(height []int) int {
	return 0
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
