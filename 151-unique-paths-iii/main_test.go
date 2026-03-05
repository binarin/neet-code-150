package main

import "testing"

func Test_uniquePathsIII(t *testing.T) {
	tests := []struct {
		name string
		grid [][]int
		want int
	}{
		{
			name: "Example 1",
			grid: [][]int{{1, 0, 0, 0}, {0, 0, 0, 0}, {0, 0, 2, -1}},
			want: 2,
		},
		{
			name: "Example 2",
			grid: [][]int{{1, 0, 0, 0}, {0, 0, 0, 0}, {0, 0, 0, 2}},
			want: 4,
		},
		{
			name: "Example 3",
			grid: [][]int{{0, 1}, {2, 0}},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := uniquePathsIII(tt.grid); got != tt.want {
				t.Errorf("uniquePathsIII() = %v, want %v", got, tt.want)
			}
		})
	}
}
