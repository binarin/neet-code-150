package main

import "testing"

func Test_findMin(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{
			name: "example 1: rotated 3 times",
			nums: []int{3, 4, 5, 1, 2},
			want: 1,
		},
		{
			name: "example 2: rotated 4 times",
			nums: []int{4, 5, 6, 7, 0, 1, 2},
			want: 0,
		},
		{
			name: "example 3: rotated n times (no rotation)",
			nums: []int{11, 13, 15, 17},
			want: 11,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findMin(tt.nums); got != tt.want {
				t.Errorf("findMin() = %v, want %v", got, tt.want)
			}
		})
	}
}
