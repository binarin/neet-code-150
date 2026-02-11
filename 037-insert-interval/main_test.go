package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInsert(t *testing.T) {
	assert.Equal(t, [][]int{{1, 5}, {6, 9}}, insert([][]int{{1, 3}, {6, 9}}, []int{2, 5}))
	assert.Equal(t, [][]int{{1, 2}, {3, 10}, {12, 16}}, insert([][]int{{1, 2}, {3, 5}, {6, 7}, {8, 10}, {12, 16}}, []int{4, 8}))
	assert.Equal(t, [][]int{{1, 5}}, insert([][]int{{1, 5}}, []int{2, 3}))
}
