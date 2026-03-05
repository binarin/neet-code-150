package main

import "testing"

func Test_minEatingSpeed(t *testing.T) {
	tests := []struct {
		name   string
		piles  []int
		h      int
		expect int
	}{
		{
			name:   "Example 1",
			piles:  []int{3, 6, 7, 11},
			h:      8,
			expect: 4,
		},
		{
			name:   "Example 2",
			piles:  []int{30, 11, 23, 4, 20},
			h:      5,
			expect: 30,
		},
		{
			name:   "Example 3",
			piles:  []int{30, 11, 23, 4, 20},
			h:      6,
			expect: 23,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := minEatingSpeed(tt.piles, tt.h)
			if got != tt.expect {
				t.Errorf("minEatingSpeed(%v, %d) = %d, want %d", tt.piles, tt.h, got, tt.expect)
			}
		})
	}
}
