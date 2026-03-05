package main

import "testing"

func Test_coinChange(t *testing.T) {
	tests := []struct {
		name   string
		coins  []int
		amount int
		want   int
	}{
		{
			name:   "example 1",
			coins:  []int{1, 2, 5},
			amount: 11,
			want:   3,
		},
		{
			name:   "example 2",
			coins:  []int{2},
			amount: 3,
			want:   -1,
		},
		{
			name:   "example 3",
			coins:  []int{1},
			amount: 0,
			want:   0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := coinChange(tt.coins, tt.amount); got != tt.want {
				t.Errorf("coinChange() = %v, want %v", got, tt.want)
			}
		})
	}
}
