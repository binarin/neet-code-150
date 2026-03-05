package main

import "testing"

func TestSearch(t *testing.T) {
	tests := []struct {
		name   string
		nums   []int
		target int
		want   int
	}{
		{
			name:   "Example 1: target exists",
			nums:   []int{-1, 0, 3, 5, 9, 12},
			target: 9,
			want:   4,
		},
		{
			name:   "Example 2: target does not exist",
			nums:   []int{-1, 0, 3, 5, 9, 12},
			target: 2,
			want:   -1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := search(tt.nums, tt.target)
			if got != tt.want {
				t.Errorf("search(%v, %d) = %d, want %d", tt.nums, tt.target, got, tt.want)
			}
		})
	}
}
