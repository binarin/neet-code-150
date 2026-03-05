package main

import (
	"reflect"
	"testing"
)

func TestWallsAndGates(t *testing.T) {
	tests := []struct {
		name     string
		rooms    [][]int
		expected [][]int
	}{
		{
			name: "Example 1",
			rooms: [][]int{
				{2147483647, -1, 0, 2147483647},
				{2147483647, 2147483647, 2147483647, -1},
				{2147483647, -1, 2147483647, -1},
				{0, -1, 2147483647, 2147483647},
			},
			expected: [][]int{
				{3, -1, 0, 1},
				{2, 2, 1, -1},
				{1, -1, 2, -1},
				{0, -1, 3, 4},
			},
		},
		{
			name: "Example 2",
			rooms: [][]int{
				{-1},
			},
			expected: [][]int{
				{-1},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			wallsAndGates(tt.rooms)
			if !reflect.DeepEqual(tt.rooms, tt.expected) {
				t.Errorf("wallsAndGates() = %v, want %v", tt.rooms, tt.expected)
			}
		})
	}
}
