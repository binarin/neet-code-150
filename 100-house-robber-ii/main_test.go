package main

import "testing"

func TestRob(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected int
	}{
		{
			name:     "Example 1",
			nums:     []int{2, 3, 2},
			expected: 3,
		},
		{
			name:     "Example 2",
			nums:     []int{1, 2, 3, 1},
			expected: 4,
		},
		{
			name:     "Example 3",
			nums:     []int{1, 2, 3},
			expected: 3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := rob(tt.nums)
			if result != tt.expected {
				t.Errorf("rob(%v) = %v, want %v", tt.nums, result, tt.expected)
			}
		})
	}
}
