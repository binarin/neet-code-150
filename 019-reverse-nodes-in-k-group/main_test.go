package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReverseKGroup(t *testing.T) {
	assert.Equal(t, []int{2, 1, 4, 3, 5}, listToSlice(reverseKGroup(sliceToList([]int{1, 2, 3, 4, 5}), 2)))
	assert.Equal(t, []int{3, 2, 1, 4, 5}, listToSlice(reverseKGroup(sliceToList([]int{1, 2, 3, 4, 5}), 3)))
}
