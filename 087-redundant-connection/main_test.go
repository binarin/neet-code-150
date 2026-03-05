package main

import (
	"slices"
	"testing"
)

func TestFindRedundantConnection(t *testing.T) {
	tests := []struct {
		name   string
		edges  [][]int
		expect []int
	}{
		{
			name:   "example 1",
			edges:  [][]int{{1, 2}, {1, 3}, {2, 3}},
			expect: []int{2, 3},
		},
		{
			name:   "example 2",
			edges:  [][]int{{1, 2}, {2, 3}, {3, 4}, {1, 4}, {1, 5}},
			expect: []int{1, 4},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := findRedundantConnection(tt.edges)
			if !slices.Equal(got, tt.expect) {
				t.Errorf("findRedundantConnection(%v) = %v, want %v", tt.edges, got, tt.expect)
			}
		})
	}
}
