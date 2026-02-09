package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPermute(t *testing.T) {
	assert.ElementsMatch(t, [][]int{{1, 2, 3}, {1, 3, 2}, {2, 1, 3}, {2, 3, 1}, {3, 1, 2}, {3, 2, 1}}, permute([]int{1, 2, 3}))
	assert.ElementsMatch(t, [][]int{{0, 1}, {1, 0}}, permute([]int{0, 1}))
	assert.ElementsMatch(t, [][]int{{1}}, permute([]int{1}))
}
