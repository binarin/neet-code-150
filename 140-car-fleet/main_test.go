package main

import "testing"

func Test_carFleet(t *testing.T) {
	tests := []struct {
		name     string
		target   int
		position []int
		speed    []int
		want     int
	}{
		{
			name:     "Example 1",
			target:   12,
			position: []int{10, 8, 0, 5, 3},
			speed:    []int{2, 4, 1, 1, 3},
			want:     3,
		},
		{
			name:     "Example 2",
			target:   10,
			position: []int{3},
			speed:    []int{3},
			want:     1,
		},
		{
			name:     "Example 3",
			target:   100,
			position: []int{0, 2, 4},
			speed:    []int{4, 2, 1},
			want:     1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := carFleet(tt.target, tt.position, tt.speed); got != tt.want {
				t.Errorf("carFleet() = %v, want %v", got, tt.want)
			}
		})
	}
}
