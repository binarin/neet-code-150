package main

import (
	"reflect"
	"testing"
)

func TestReverseList(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected []int
	}{
		{
			name:     "Example 1: [1,2,3,4,5] -> [5,4,3,2,1]",
			input:    []int{1, 2, 3, 4, 5},
			expected: []int{5, 4, 3, 2, 1},
		},
		{
			name:     "Example 2: [1,2] -> [2,1]",
			input:    []int{1, 2},
			expected: []int{2, 1},
		},
		{
			name:     "Example 3: [] -> []",
			input:    []int{},
			expected: []int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			head := buildList(tt.input)
			result := reverseList(head)
			got := listToSlice(result)
			if !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("reverseList() = %v, want %v", got, tt.expected)
			}
		})
	}
}
