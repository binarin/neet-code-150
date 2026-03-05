package main

import "testing"

func Test_mergeTriplets(t *testing.T) {
	tests := []struct {
		name     string
		triplets [][]int
		target   []int
		want     bool
	}{
		{
			name:     "example 1",
			triplets: [][]int{{2, 5, 3}, {1, 8, 4}, {1, 7, 5}},
			target:   []int{2, 7, 5},
			want:     true,
		},
		{
			name:     "example 2",
			triplets: [][]int{{3, 4, 5}, {4, 5, 6}},
			target:   []int{3, 2, 5},
			want:     false,
		},
		{
			name:     "example 3",
			triplets: [][]int{{2, 5, 3}, {2, 3, 4}, {1, 2, 5}, {5, 2, 3}},
			target:   []int{5, 5, 5},
			want:     true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mergeTriplets(tt.triplets, tt.target); got != tt.want {
				t.Errorf("mergeTriplets() = %v, want %v", got, tt.want)
			}
		})
	}
}
