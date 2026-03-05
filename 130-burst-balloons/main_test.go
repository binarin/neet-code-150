package main

import "testing"

func TestMaxCoins(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{
			name: "Example 1",
			nums: []int{3, 1, 5, 8},
			want: 167,
		},
		{
			name: "Example 2",
			nums: []int{1, 5},
			want: 10,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := maxCoins(tt.nums)
			if got != tt.want {
				t.Errorf("maxCoins() = %v, want %v", got, tt.want)
			}
		})
	}
}
