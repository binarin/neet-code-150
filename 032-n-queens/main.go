package main

import (
	"fmt"
	"log"
	"strings"
)

func main() {
	for _, sol := range solveNQueens(4) {
		for _, ln := range sol {
			fmt.Println(ln)
		}
		fmt.Println()
	}
}

func solveNQueens(n int) [][]string {
	if n == 0 {
		log.Fatal("n == 0")
	}
	if n == 1 {
		return [][]string{{"Q"}}
	}

	var result [][]string
	sub_solutions := solveNQueens(n - 1)
	// sub_solutions: (n-1)x(n-1)
	// 0 <= insert_at <= n - 1
	// 0 <= insert_at < n: OK!
	for insert_at := range n {
	sub_sols:
		for _, ss := range sub_solutions {
			var sol []string
			for i := range n {
				if i == insert_at {
					sol = append(sol, "Q"+strings.Repeat(".", n-1))
					continue
				}

				sub_i := i
				if i > insert_at {
					sub_i--
				}

				if sub_i >= n-1 {
					continue
				}

				if n >= 3 {
					if i < insert_at {
						x := insert_at - i - 1
						if ss[sub_i][x] == 'Q' {
							continue sub_sols
						}
					}

					if i > insert_at {
						x := i - insert_at - 1
						if ss[sub_i][x] == 'Q' {
							continue sub_sols
						}
					}
				}

				sol = append(sol, "."+ss[sub_i])
			}

			result = append(result, sol)
		}
	}

	return result
}
