package main

import "testing"

func TestReverseBits(t *testing.T) {
	tests := []struct {
		name     string
		n        int
		expected int
	}{
		{
			name:     "example 1",
			n:        43261596,
			expected: 964176192,
		},
		{
			name:     "example 2",
			n:        2147483644,
			expected: 1073741822,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := reverseBits(tt.n)
			if result != tt.expected {
				t.Errorf("reverseBits(%d) = %d, expected %d", tt.n, result, tt.expected)
			}
		})
	}
}
