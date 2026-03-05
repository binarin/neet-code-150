package main

import "testing"

func Test_minCostConnectPoints(t *testing.T) {
	tests := []struct {
		name   string
		points [][]int
		want   int
	}{
		{
			name:   "example 1",
			points: [][]int{{0, 0}, {2, 2}, {3, 10}, {5, 2}, {7, 0}},
			want:   20,
		},
		{
			name:   "example 2",
			points: [][]int{{3, 12}, {-2, 5}, {-4, 1}},
			want:   18,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minCostConnectPoints(tt.points); got != tt.want {
				t.Errorf("minCostConnectPoints() = %v, want %v", got, tt.want)
			}
		})
	}
}
