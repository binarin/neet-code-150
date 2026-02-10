package main

import (
	"fmt"
	"slices"
)

func main() {
	for _, sol := range solveNQueens(4) {
		for _, ln := range sol {
			fmt.Println(ln)
		}
		fmt.Println()
	}
}

type Board struct {
	N     int
	Cells []byte
}

func (b Board) ToSolution() []string {
	result := make([]string, 0, b.N)
	for row := range b.N {
		bytes := slices.Clone(b.Cells[row*b.N : (row+1)*b.N])
		for i, b := range bytes {
			if b == 'X' {
				bytes[i] = '.'
			}
		}
		result = append(result, string(bytes))
	}
	return result
}

func (b Board) At(x, y int) byte {
	return b.Cells[y*b.N+x]
}

func (b Board) Put(x, y int, val byte) {
	b.Cells[y*b.N+x] = val
}

func newBoard(n int) Board {
	return Board{N: n, Cells: slices.Repeat([]byte{'.'}, n*n)}
}

func (b *Board) Clone() Board {
	cs := slices.Clone(b.Cells)
	return Board{N: b.N, Cells: cs}
}

func solveNQueens(n int) [][]string {
	var result [][]string

	var recur func(int, Board)
	recur = func(col int, b Board) {
		if col == n {
			result = append(result, b.ToSolution())
			return
		}
		for row := range n {
			if b.At(col, row) != '.' {
				continue
			}
			nb := b.Clone()
			for dx, dy := 1, 1; col+dx < n; dx, dy = dx+1, dy+1 {
				nb.Put(col+dx, row, 'X')
				if row-dy >= 0 && col+dx < n {
					nb.Put(col+dx, row-dy, 'X')
				}
				if row+dx < n && col+dy < n {
					nb.Put(col+dx, row+dy, 'X')
				}
			}
			nb.Put(col, row, 'Q')
			recur(col+1, nb)
		}
	}

	recur(0, newBoard(n))

	return result
}
