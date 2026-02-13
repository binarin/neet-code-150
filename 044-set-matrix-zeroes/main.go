package main

import (
	"fmt"
)

func main() {
	matrix := [][]int{{1, 1, 1}, {1, 0, 1}, {1, 1, 1}}
	setZeroes(matrix)
	fmt.Println(matrix)
}

func setZeroes(matrix [][]int) {
	width := len(matrix[0])
	height := len(matrix)

	flush_rows := make([]bool, height)
	flush_cols := make([]bool, width)

	for y := range matrix {
		for x := range matrix[0] {
			if matrix[y][x] == 0 {
				flush_rows[y] = true
				flush_cols[x] = true
			}
		}
	}

	for y := range matrix {
		for x := range matrix[0] {
			if flush_rows[y] || flush_cols[x] {
				matrix[y][x] = 0
			}
		}
	}
}
