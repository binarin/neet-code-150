package main

import "testing"

func Test_canPartition(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want bool
	}{
		{
			name: "example 1",
			nums: []int{1, 5, 11, 5},
			want: true,
		},
		{
			name: "example 2",
			nums: []int{1, 2, 3, 5},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := canPartition(tt.nums); got != tt.want {
				t.Errorf("canPartition() = %v, want %v", got, tt.want)
			}
		})
	}
}
