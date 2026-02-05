package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCombinationSum(t *testing.T) {
	assert.ElementsMatch(t, [][]int{{2, 2, 3}, {7}}, combinationSum([]int{2, 3, 6, 7}, 7))
	assert.ElementsMatch(t, [][]int{{2, 2, 2, 2}, {2, 3, 3}, {3, 5}}, combinationSum([]int{2, 3, 5}, 8))
	assert.ElementsMatch(t, [][]int{}, combinationSum([]int{2}, 1))
}
