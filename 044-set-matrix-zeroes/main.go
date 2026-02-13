package main

import (
	"fmt"
)

func main() {
	matrix := [][]int{{0}, {0}}
	setZeroes(matrix)
	fmt.Println(matrix)
}

func setZeroes(matrix [][]int) {
	width := len(matrix[0])
	height := len(matrix)

	var clear_top, clear_left bool
	for y := range height {
		if matrix[y][0] == 0 {
			clear_left = true
			break
		}
	}
	for x := range width {
		if matrix[0][x] == 0 {
			clear_top = true
			break
		}
	}

	for y := 1; y < height; y++ {
		for x := 1; x < width; x++ {
			if matrix[y][x] == 0 {
				matrix[0][x] = 0
				matrix[y][0] = 0
			}
		}
	}

	for y := 1; y < height; y++ {
		for x := 1; x < width; x++ {
			if matrix[0][x] == 0 || matrix[y][0] == 0 {
				matrix[y][x] = 0
			}
		}
	}
	if clear_top {
		for x := range width {
			matrix[0][x] = 0
		}
	}

	if clear_left {
		for y := range height {
			matrix[y][0] = 0
		}
	}

}
