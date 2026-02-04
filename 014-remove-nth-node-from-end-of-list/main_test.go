package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRemoveNthFromEnd(t *testing.T) {
	assert.Equal(t, []int{1, 2, 3, 5}, listToSlice(removeNthFromEnd(sliceToList([]int{1, 2, 3, 4, 5}), 2)))
	assert.Equal(t, []int{}, listToSlice(removeNthFromEnd(sliceToList([]int{1}), 1)))
	assert.Equal(t, []int{1}, listToSlice(removeNthFromEnd(sliceToList([]int{1, 2}), 1)))
}
