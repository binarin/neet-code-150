package main

import (
	"fmt"
)

func main() {
	fmt.Println(orangesRotting([][]int{{2, 1, 0, 2}}))
}

const EMPTY = 0
const FRESH = 1
const ROTTEN = 2

func orangesRotting(grid [][]int) int {
	h := len(grid)
	w := len(grid[0])

	get := func(x int, y int) int {
		if x < 0 || y < 0 || x >= w || y >= h {
			return EMPTY
		}
		return grid[y][x]
	}
	is_rotten := func(x int, y int) bool {
		return get(x, y) == ROTTEN
	}

	step := 0
	for {
		seen_change := false
		seen_fresh := false
		updates := []int{}

		for x := range w {
			for y := range h {
				if grid[y][x] == FRESH {
					has_rotten_neigbhor := is_rotten(x-1, y) || is_rotten(x+1, y) || is_rotten(x, y-1) || is_rotten(x, y+1)
					if has_rotten_neigbhor {
						updates = append(updates, y*w+x)
						seen_change = true
					} else {
						seen_fresh = true
					}
				}
			}
		}
		for _, up := range updates {
			grid[up/w][up%w] = ROTTEN
		}
		fmt.Println(grid)
		if !seen_change {
			if seen_fresh {
				return -1
			} else {
				return step
			}
		}
		step++
	}
}
