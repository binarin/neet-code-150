package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMerge(t *testing.T) {
	assert.Equal(t, [][]int{{1, 6}, {8, 10}, {15, 18}}, merge([][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}))
	assert.Equal(t, [][]int{{1, 5}}, merge([][]int{{1, 4}, {4, 5}}))
	assert.Equal(t, [][]int{{1, 7}}, merge([][]int{{4, 7}, {1, 4}}))
	assert.Equal(t, [][]int{{0, 5}}, merge([][]int{{4, 5}, {1, 4}, {0, 1}}))
}
