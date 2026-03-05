package main

import "testing"

func Test_isNStraightHand(t *testing.T) {
	tests := []struct {
		name      string
		hand      []int
		groupSize int
		want      bool
	}{
		{
			name:      "Example 1: can rearrange into groups of 3",
			hand:      []int{1, 2, 3, 6, 2, 3, 4, 7, 8},
			groupSize: 3,
			want:      true,
		},
		{
			name:      "Example 2: cannot rearrange into groups of 4",
			hand:      []int{1, 2, 3, 4, 5},
			groupSize: 4,
			want:      false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isNStraightHand(tt.hand, tt.groupSize); got != tt.want {
				t.Errorf("isNStraightHand() = %v, want %v", got, tt.want)
			}
		})
	}
}
