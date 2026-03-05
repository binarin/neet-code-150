package main

import "testing"

func Test_longestIncreasingPath(t *testing.T) {
	tests := []struct {
		name   string
		matrix [][]int
		want   int
	}{
		{
			name:   "example 1",
			matrix: [][]int{{9, 9, 4}, {6, 6, 8}, {2, 1, 1}},
			want:   4,
		},
		{
			name:   "example 2",
			matrix: [][]int{{3, 4, 5}, {3, 2, 6}, {2, 2, 1}},
			want:   4,
		},
		{
			name:   "example 3",
			matrix: [][]int{{1}},
			want:   1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestIncreasingPath(tt.matrix); got != tt.want {
				t.Errorf("longestIncreasingPath() = %v, want %v", got, tt.want)
			}
		})
	}
}
