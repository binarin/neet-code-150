package main

import (
	"slices"
	"testing"

	"github.com/stretchr/testify/assert"
)

func normalizeCombinations(combos [][]int) [][]int {
	result := make([][]int, len(combos))
	for i, combo := range combos {
		sorted := make([]int, len(combo))
		copy(sorted, combo)
		slices.Sort(sorted)
		result[i] = sorted
	}
	slices.SortFunc(result, func(a, b []int) int {
		for i := 0; i < len(a) && i < len(b); i++ {
			if a[i] != b[i] {
				return a[i] - b[i]
			}
		}
		return len(a) - len(b)
	})
	return result
}

func TestCombinationSum2(t *testing.T) {
	got := combinationSum2([]int{10, 1, 2, 7, 6, 1, 5}, 8)
	want := [][]int{{1, 1, 6}, {1, 2, 5}, {1, 7}, {2, 6}}
	assert.Equal(t, normalizeCombinations(want), normalizeCombinations(got))

	got = combinationSum2([]int{2, 5, 2, 1, 2}, 5)
	want = [][]int{{1, 2, 2}, {5}}
	assert.Equal(t, normalizeCombinations(want), normalizeCombinations(got))
}
