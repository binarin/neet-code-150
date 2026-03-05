package main

import (
	"reflect"
	"testing"
)

func TestCountBits(t *testing.T) {
	tests := []struct {
		name     string
		n        int
		expected []int
	}{
		{
			name:     "Example 1",
			n:        2,
			expected: []int{0, 1, 1},
		},
		{
			name:     "Example 2",
			n:        5,
			expected: []int{0, 1, 1, 2, 1, 2},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := countBits(tt.n)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("countBits(%d) = %v, expected %v", tt.n, result, tt.expected)
			}
		})
	}
}
