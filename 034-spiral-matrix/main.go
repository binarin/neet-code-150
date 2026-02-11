package main

import (
	"fmt"
)

func main() {
	fmt.Println(spiralOrder([][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}))
}

type Dir int

const (
	Right Dir = iota
	Down
	Left
	Up
)

func spiralOrder(matrix [][]int) []int {
	height := len(matrix)
	width := len(matrix[0])

	left, top := 0, 0
	right, bottom := width-1, height-1
	used, total := 0, width*height
	result := make([]int, 0, total)

	for {
		x, y := left, top
		for x <= right {
			result = append(result, matrix[y][x])
			x++
			used++
		}
		top++
		if used == total {
			break
		}

		x, y = right, top
		for y <= bottom {
			result = append(result, matrix[y][x])
			y++
			used++
		}
		right--
		if used == total {
			break
		}

		x, y = right, bottom
		for x >= left {
			result = append(result, matrix[y][x])
			x--
			used++
		}
		bottom--
		if used == total {
			break
		}

		x, y = left, bottom
		for y >= top {
			result = append(result, matrix[y][x])
			y--
			used++
		}
		left++
		if used == total {
			break
		}
	}

	return result
}
