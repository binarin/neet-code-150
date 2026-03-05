package main

import "testing"

func Test_swimInWater(t *testing.T) {
	tests := []struct {
		name string
		grid [][]int
		want int
	}{
		{
			name: "example 1",
			grid: [][]int{{0, 2}, {1, 3}},
			want: 3,
		},
		{
			name: "example 2",
			grid: [][]int{
				{0, 1, 2, 3, 4},
				{24, 23, 22, 21, 5},
				{12, 13, 14, 15, 16},
				{11, 17, 18, 19, 20},
				{10, 9, 8, 7, 6},
			},
			want: 16,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := swimInWater(tt.grid); got != tt.want {
				t.Errorf("swimInWater() = %v, want %v", got, tt.want)
			}
		})
	}
}
