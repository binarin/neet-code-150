package main

import "testing"

func Test_maxProfit(t *testing.T) {
	tests := []struct {
		name   string
		prices []int
		want   int
	}{
		{
			name:   "Example 1: buy, sell, cooldown, buy, sell",
			prices: []int{1, 2, 3, 0, 2},
			want:   3,
		},
		{
			name:   "Example 2: single price",
			prices: []int{1},
			want:   0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxProfit(tt.prices); got != tt.want {
				t.Errorf("maxProfit() = %v, want %v", got, tt.want)
			}
		})
	}
}
