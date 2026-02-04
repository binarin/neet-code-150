package main

import (
	"fmt"
)

func main() {
	fmt.Println(maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}))
}

func maxArea(height []int) int {
	if len(height) < 2 {
		return 0
	}
	left := 0
	right := len(height) - 1
	max_area := 0
	for left < right {
		area := min(height[left], height[right]) * (right - left)
		max_area = max(max_area, area)
		if height[left] < height[right] { // the smaller can never improve situation
			left++
		} else {
			right--
		}
	}
	return max_area
}

func maxAreaBf(height []int) int {
	if len(height) < 2 {
		return 0
	}
	max_area := 0
	for i := 0; i < len(height)-1; i++ {
		for j := i + 1; j < len(height); j++ {
			area := (j - i) * min(height[i], height[j])
			max_area = max(area, max_area)
		}
	}

	return max_area
}
