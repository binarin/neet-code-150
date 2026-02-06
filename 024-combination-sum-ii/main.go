package main

import (
	"fmt"
)

func main() {
	fmt.Println(combinationSum2([]int{10, 1, 2, 7, 6, 1, 5}, 8))
}

func combinationSum2(candidates []int, target int) [][]int {
	const CandDomainMax = 50
	type Freqs [CandDomainMax + 1]int

	result := [][]int{}
	freqs := Freqs{}

	for _, val := range candidates {
		freqs[val]++
	}

	add_solution := func(fs Freqs) {
		solution := []int{}
		for val, count := range fs {
			for range count {
				solution = append(solution, val)
			}
		}
		result = append(result, solution)
	}

	var recur func(int, int, Freqs)
	recur = func(val int, so_far int, used Freqs) {
		if so_far == target {
			add_solution(used)
			return
		}
		if so_far > target {
			return
		}
		if val >= len(freqs) {
			return
		}
		recur(val+1, so_far, used)
		for range freqs[val] - used[val] {
			used[val]++
			so_far += val
			recur(val+1, so_far, used)
		}
	}

	recur(0, 0, Freqs{})

	return result
}
