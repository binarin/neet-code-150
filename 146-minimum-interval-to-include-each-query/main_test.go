package main

import (
	"reflect"
	"testing"
)

func TestMinInterval(t *testing.T) {
	tests := []struct {
		intervals [][]int
		queries   []int
		expected  []int
	}{
		{
			intervals: [][]int{{1, 4}, {2, 4}, {3, 6}, {4, 4}},
			queries:   []int{2, 3, 4, 5},
			expected:  []int{3, 3, 1, 4},
		},
		{
			intervals: [][]int{{2, 3}, {2, 5}, {1, 8}, {20, 25}},
			queries:   []int{2, 19, 5, 22},
			expected:  []int{2, -1, 4, 6},
		},
	}

	for _, tt := range tests {
		result := minInterval(tt.intervals, tt.queries)
		if !reflect.DeepEqual(result, tt.expected) {
			t.Errorf("minInterval(%v, %v) = %v, expected %v", tt.intervals, tt.queries, result, tt.expected)
		}
	}
}
