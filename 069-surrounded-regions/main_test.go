package main

import (
	"reflect"
	"testing"
)

func TestSolve(t *testing.T) {
	tests := []struct {
		name     string
		board    [][]byte
		expected [][]byte
	}{
		{
			name: "Example 1",
			board: [][]byte{
				{'X', 'X', 'X', 'X'},
				{'X', 'O', 'O', 'X'},
				{'X', 'X', 'O', 'X'},
				{'X', 'O', 'X', 'X'},
			},
			expected: [][]byte{
				{'X', 'X', 'X', 'X'},
				{'X', 'X', 'X', 'X'},
				{'X', 'X', 'X', 'X'},
				{'X', 'O', 'X', 'X'},
			},
		},
		{
			name: "Example 2",
			board: [][]byte{
				{'X'},
			},
			expected: [][]byte{
				{'X'},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			solve(tt.board)
			if !reflect.DeepEqual(tt.board, tt.expected) {
				t.Errorf("solve() = %v, want %v", tt.board, tt.expected)
			}
		})
	}
}
