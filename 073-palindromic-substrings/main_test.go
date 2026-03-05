package main

import "testing"

func TestCountSubstrings(t *testing.T) {
	tests := []struct {
		name     string
		s        string
		expected int
	}{
		{
			name:     "Example 1: abc",
			s:        "abc",
			expected: 3,
		},
		{
			name:     "Example 2: aaa",
			s:        "aaa",
			expected: 6,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := countSubstrings(tt.s)
			if result != tt.expected {
				t.Errorf("countSubstrings(%q) = %d, expected %d", tt.s, result, tt.expected)
			}
		})
	}
}
