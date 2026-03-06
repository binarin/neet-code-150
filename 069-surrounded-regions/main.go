// https://leetcode.com/problems/surrounded-regions/
// Difficulty: Medium

package main

import "fmt"

func solve(board [][]byte) {
	type pt [2]int
	queue := []pt{}
	keep := map[pt]bool{}
	height, width := len(board), len(board[0])

	enqueue := func(c pt) {
		if c[0] < 0 || c[0] >= width || c[1] < 0 || c[1] >= height {
			return
		}
		if board[c[1]][c[0]] == 'O' {
			if _, ok := keep[c]; !ok {
				queue = append(queue, c)
				keep[c] = true
			}
		}
	}

	for x := range width {
		enqueue(pt{x, 0})
		enqueue(pt{x, height - 1})
	}

	for y := range height {
		enqueue(pt{0, y})
		enqueue(pt{width - 1, y})
	}

	deltas := [4]pt{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
	for len(queue) > 0 {
		c := queue[0]
		queue = queue[1:]
		for _, d := range deltas {
			enqueue(pt{c[0] + d[0], c[1] + d[1]})
		}
	}

	for y := range height {
		for x := range width {
			if board[y][x] == 'O' && !keep[pt{x, y}] {
				board[y][x] = 'X'
			}
		}
	}
}

func main() {
	board := [][]byte{
		{'X', 'X', 'X', 'X'},
		{'X', 'O', 'O', 'X'},
		{'X', 'X', 'O', 'X'},
		{'X', 'O', 'X', 'X'},
	}
	solve(board)
	fmt.Println(board)
}
