package main

import "testing"

func TestLeastInterval(t *testing.T) {
	tests := []struct {
		name     string
		tasks    []byte
		n        int
		expected int
	}{
		{
			name:     "Example 1: tasks = [A,A,A,B,B,B], n = 2",
			tasks:    []byte{'A', 'A', 'A', 'B', 'B', 'B'},
			n:        2,
			expected: 8,
		},
		{
			name:     "Example 2: tasks = [A,C,A,B,D,B], n = 1",
			tasks:    []byte{'A', 'C', 'A', 'B', 'D', 'B'},
			n:        1,
			expected: 6,
		},
		{
			name:     "Example 3: tasks = [A,A,A,B,B,B], n = 3",
			tasks:    []byte{'A', 'A', 'A', 'B', 'B', 'B'},
			n:        3,
			expected: 10,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := leastInterval(tt.tasks, tt.n)
			if result != tt.expected {
				t.Errorf("leastInterval(%v, %d) = %d; expected %d", tt.tasks, tt.n, result, tt.expected)
			}
		})
	}
}
