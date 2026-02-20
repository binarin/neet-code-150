package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSubsets(t *testing.T) {
	assert.Equal(t, [][]int{{}, {1}, {2}, {1, 2}, {3}, {1, 3}, {2, 3}, {1, 2, 3}}, subsets([]int{1, 2, 3}))
	assert.Equal(t, [][]int{{}, {0}}, subsets([]int{0}))
}
