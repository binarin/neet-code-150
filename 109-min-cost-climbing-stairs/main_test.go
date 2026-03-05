package main

import "testing"

func Test_minCostClimbingStairs(t *testing.T) {
	tests := []struct {
		name string
		cost []int
		want int
	}{
		{
			name: "Example 1",
			cost: []int{10, 15, 20},
			want: 15,
		},
		{
			name: "Example 2",
			cost: []int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1},
			want: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minCostClimbingStairs(tt.cost); got != tt.want {
				t.Errorf("minCostClimbingStairs() = %v, want %v", got, tt.want)
			}
		})
	}
}
