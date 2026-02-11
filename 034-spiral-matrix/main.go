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
	m := len(matrix)
	n := len(matrix[0])
	if m == 0 || n == 0 {
		return nil
	}
	if m == 1 {
		return matrix[0]
	}
	if n == 1 {
		result := make([]int, 0, m)
		for y := range matrix {
			result = append(result, matrix[y][0])
		}
		return result
	}
	x := 0
	y := 0
	hspan := n - 1
	vspan := m - 1
	result := []int{}
	used := 0
	total := n * m
	for {
		for range hspan {
			result = append(result, matrix[y][x])
			x++
			used++
		}
		if used == total {
			break
		}
		for range vspan {
			result = append(result, matrix[y][x])
			y++
			used++
		}
		if used == total {
			break
		}
		for range hspan {
			result = append(result, matrix[y][x])
			x--
			used++
		}
		if used == total {
			break
		}
		for range vspan {
			result = append(result, matrix[y][x])
			y--
			used++
		}
		if used == total {
			break
		}
		x++
		y++
		hspan--
		vspan--
	}
	return result
}
