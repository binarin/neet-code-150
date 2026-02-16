package main

import (
	"fmt"
)

func main() {
	// fmt.Println(searchMatrix([][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}}, 13))
	fmt.Println(searchMatrix([][]int{{1, 3}}, 3))
}

func searchMatrix(matrix [][]int, target int) bool {
	height := len(matrix)
	width := len(matrix[0])
	left := 0
	right := height*width - 1
	get := func(pos int) int {
		return matrix[pos/width][pos%width]
	}

	if get(left) == target || get(right) == target {
		return true
	}

	for {
		midpoint := (left + right) / 2
		mv := get(midpoint)
		if target < mv && right != midpoint {
			right = midpoint
		} else if target == mv {
			return true
		} else if target > mv && left != midpoint {
			left = midpoint
		} else {
			return false
		}
	}
}
