package main

import (
	"fmt"
)

func main() {
	board := [][]byte{
		{'A', 'A', 'A', 'A', 'A', 'A'},
		{'A', 'A', 'A', 'B', 'A', 'A'},
		{'A', 'A', 'A', 'A', 'A', 'A'},
		{'A', 'A', 'A', 'A', 'A', 'A'},
		{'A', 'A', 'A', 'A', 'A', 'A'},
		{'A', 'A', 'A', 'A', 'A', 'A'},
	}
	// board := [][]byte{
	// 	{'A', 'B', 'C', 'E'},
	// 	{'S', 'F', 'C', 'S'},
	// 	{'A', 'D', 'E', 'E'},
	// }
	fmt.Println(exist(board, "AB"))
}

func use(...any) {

}

func exist(board [][]byte, word string) bool {
	height := len(board)
	width := len(board[0])

	char_at := func(x, y int, used map[int]bool) byte {
		if x < 0 || y < 0 || x >= width || y >= height {
			return '!'
		}
		if _, ok := used[y*width+x]; ok {
			return '!'
		}
		return board[y][x]
	}

	dump_state := func(cur_x, cur_y int, used map[int]bool, word string) {
		fmt.Println("Word", word)
		for y := range height {
			for x := range width {
				if x == cur_x && y == cur_y {
					fmt.Print("·")
				} else if _, ok := used[y*width+x]; ok {
					fmt.Print("!")
				} else {
					fmt.Print(string(board[y][x]))
				}
			}
			fmt.Println()
		}
		fmt.Println()
	}
	use(dump_state)

	var recur func(int, int, string, map[int]bool) bool
	recur = func(x, y int, subword string, used map[int]bool) bool {
		// dump_state(x, y, used, subword)
		if len(subword) == 0 {
			return true
		}
		deltas := [4][2]int{{0, -1}, {0, 1}, {-1, 0}, {1, 0}}
		for _, delta := range deltas {
			new_x, new_y := x+delta[0], y+delta[1]
			if char_at(new_x, new_y, used) != subword[0] {
				continue
			}
			used_idx := new_y*width + new_x
			used[used_idx] = true
			result := recur(new_x, new_y, subword[1:], used)
			delete(used, used_idx)
			if result {
				return true
			}
		}
		return false
	}

	for y := range height {
		for x := range width {
			if board[y][x] == word[0] {
				if recur(x, y, word[1:], map[int]bool{(y*width + x): true}) {
					return true
				}
			}
		}
	}

	return false
}
