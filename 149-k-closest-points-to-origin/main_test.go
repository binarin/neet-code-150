package main

import (
	"sort"
	"testing"
)

func sortPoints(points [][]int) {
	sort.Slice(points, func(i, j int) bool {
		if points[i][0] != points[j][0] {
			return points[i][0] < points[j][0]
		}
		return points[i][1] < points[j][1]
	})
}

func equalPoints(a, b [][]int) bool {
	if len(a) != len(b) {
		return false
	}
	aCopy := make([][]int, len(a))
	bCopy := make([][]int, len(b))
	copy(aCopy, a)
	copy(bCopy, b)
	sortPoints(aCopy)
	sortPoints(bCopy)
	for i := range aCopy {
		if aCopy[i][0] != bCopy[i][0] || aCopy[i][1] != bCopy[i][1] {
			return false
		}
	}
	return true
}

func TestKClosest(t *testing.T) {
	tests := []struct {
		name     string
		points   [][]int
		k        int
		expected [][]int
	}{
		{
			name:     "Example 1",
			points:   [][]int{{1, 3}, {-2, 2}},
			k:        1,
			expected: [][]int{{-2, 2}},
		},
		{
			name:     "Example 2",
			points:   [][]int{{3, 3}, {5, -1}, {-2, 4}},
			k:        2,
			expected: [][]int{{3, 3}, {-2, 4}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := kClosest(tt.points, tt.k)
			if !equalPoints(result, tt.expected) {
				t.Errorf("kClosest(%v, %d) = %v, expected %v", tt.points, tt.k, result, tt.expected)
			}
		})
	}
}
