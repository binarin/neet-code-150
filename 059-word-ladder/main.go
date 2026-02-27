package main

import (
	"fmt"
	"slices"
)

func main() {
	result := ladderLength("hit", "cog", []string{"hot", "dot", "dog", "lot", "log", "cog"})
	fmt.Println(result)
}

func ladderLength(beginWord string, endWord string, wordList []string) int {
	use := func(...any) {
	}
	wl := len(beginWord)
	use(wl)
	lookup := map[string][]int{}

	patternize := func(s string) (result []string) {
		for idx := range len(s) {
			result = append(result, s[:idx]+"X"+s[idx+1:])
		}
		return
	}

	endWordIdx := -1
	for idx, word := range wordList {
		if word == endWord {
			endWordIdx = idx
		}
		for _, pat := range patternize(word) {
			lookup[pat] = append(lookup[pat], idx)
		}
	}
	if endWordIdx < 0 {
		return 0
	}

	wordList = append(wordList, beginWord)
	visited := slices.Repeat([]bool{false}, len(wordList))
	queue := [][2]int{{1, len(wordList) - 1}}

	// DFS, queue is FIFO - so cost in queue is non-decreasing
	for len(queue) > 0 {
		cost, idx := queue[0][0], queue[0][1]
		queue = queue[1:]
		if idx == endWordIdx {
			return cost
		}

		if visited[idx] {
			continue
		}
		fmt.Printf("==> Picked %s(%d) - cost %d\n", wordList[idx], idx, cost)

		visited[idx] = true
		for _, pattern := range patternize(wordList[idx]) {
			candidates, ok := lookup[pattern]
			fmt.Printf("Candidates for %s: %v\n", pattern, candidates)
			if !ok {
				continue
			}
			for _, candIdx := range candidates {
				if visited[candIdx] {
					continue
				}
				queue = append(queue, [2]int{cost + 1, candIdx})
			}
		}
	}
	return 0
}
