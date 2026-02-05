package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMergeKLists(t *testing.T) {
	assert.Equal(t, []int{1, 1, 2, 3, 4, 4, 5, 6}, listToSlice(mergeKLists([]*ListNode{
		mkList([]int{1, 4, 5}),
		mkList([]int{1, 3, 4}),
		mkList([]int{2, 6}),
	})))
	assert.Equal(t, []int(nil), listToSlice(mergeKLists([]*ListNode{})))
	assert.Equal(t, []int(nil), listToSlice(mergeKLists([]*ListNode{nil})))
}
