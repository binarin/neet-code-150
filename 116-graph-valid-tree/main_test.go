package main

import "testing"

func Test_validTree(t *testing.T) {
	tests := []struct {
		name   string
		n      int
		edges  [][]int
		want   bool
	}{
		{
			name:  "Example 1",
			n:     5,
			edges: [][]int{{0, 1}, {0, 2}, {0, 3}, {1, 4}},
			want:  true,
		},
		{
			name:  "Example 2",
			n:     5,
			edges: [][]int{{0, 1}, {1, 2}, {2, 3}, {1, 3}, {1, 4}},
			want:  false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := validTree(tt.n, tt.edges); got != tt.want {
				t.Errorf("validTree() = %v, want %v", got, tt.want)
			}
		})
	}
}
