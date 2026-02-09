package main

import (
	"fmt"
)

func main() {
	matrix := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	rotate(matrix)
	fmt.Println(matrix)
}

func rotate(matrix [][]int) {
	// diagonally
	n := len(matrix[0])
	for offset := 0; offset < n/2; offset++ {
		// 6: 5, 3, 1
		// 5: 4, 2
		top := offset
		bottom := n - offset - 1
		left := offset
		right := n - offset - 1
		num_cycles := n - 1 - offset*2
		for cycle_start := range num_cycles {
			tmp := matrix[top][left+cycle_start]
			matrix[top][left+cycle_start] = matrix[bottom-cycle_start][left]
			matrix[bottom-cycle_start][left] = matrix[bottom][right-cycle_start]
			matrix[bottom][right-cycle_start] = matrix[top+cycle_start][right]
			matrix[top+cycle_start][right] = tmp
		}
	}
}
