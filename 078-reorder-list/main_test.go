package main

import (
	"reflect"
	"testing"
)

func TestReorderList(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected []int
	}{
		{
			name:     "Example 1: [1,2,3,4] -> [1,4,2,3]",
			input:    []int{1, 2, 3, 4},
			expected: []int{1, 4, 2, 3},
		},
		{
			name:     "Example 2: [1,2,3,4,5] -> [1,5,2,4,3]",
			input:    []int{1, 2, 3, 4, 5},
			expected: []int{1, 5, 2, 4, 3},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			head := createList(tt.input)
			reorderList(head)
			result := listToSlice(head)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("reorderList() = %v, want %v", result, tt.expected)
			}
		})
	}
}
