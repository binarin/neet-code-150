package main

import (
	"fmt"
)

func main() {
	board := [][]byte{
		{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
		{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
		{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
		{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
		{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
		{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
	}
	fmt.Println(isValidSudoku(board))
}

func isValidSudoku(board [][]byte) bool {
	verticals := [9]map[byte]bool{}
	horizontals := [9]map[byte]bool{}
	squares := [9]map[byte]bool{}

	for x := range 9 {
		for y := range 9 {
			val := board[y][x]
			if val == '.' {
				continue
			}
			v_idx := y
			h_idx := x
			s_idx_y := (9*y + x) / 27
			s_idx_x := x / 3
			s_idx := s_idx_y*3 + s_idx_x
			if verticals[v_idx][val] || horizontals[h_idx][val] || squares[s_idx][val] {
				return false
			}
			if verticals[v_idx] == nil {
				verticals[v_idx] = map[byte]bool{}
			}
			if horizontals[h_idx] == nil {
				horizontals[h_idx] = map[byte]bool{}
			}
			if squares[s_idx] == nil {
				squares[s_idx] = map[byte]bool{}
			}
			verticals[v_idx][val] = true
			horizontals[h_idx][val] = true
			squares[s_idx][val] = true
		}
	}

	return true
}
