package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Helper function to convert a ListNode to a slice of ints
func listToSlice(head *ListNode) []int {
	nums := []int{}
	current := head
	for current != nil {
		nums = append(nums, current.Val)
		current = current.Next
	}
	return nums
}

func TestMergeTwoLists(t *testing.T) {
	tests := []struct {
		name     string
		list1    []int
		list2    []int
		expected []int
	}{
		{
			name:     "Example 1",
			list1:    []int{1, 2, 4},
			list2:    []int{1, 3, 4},
			expected: []int{1, 1, 2, 3, 4, 4},
		},
		{
			name:     "Example 2",
			list1:    []int{},
			list2:    []int{},
			expected: []int{},
		},
		{
			name:     "Example 3",
			list1:    []int{},
			list2:    []int{0},
			expected: []int{0},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			list1 := makeList(tt.list1)
			list2 := makeList(tt.list2)
			result := mergeTwoLists(list1, list2)
			assert.Equal(t, tt.expected, listToSlice(result))
		})
	}
}
