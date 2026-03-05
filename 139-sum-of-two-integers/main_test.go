package main

import "testing"

func TestGetSum(t *testing.T) {
	tests := []struct {
		a        int
		b        int
		expected int
	}{
		{1, 2, 3},
		{2, 3, 5},
	}

	for _, tt := range tests {
		result := getSum(tt.a, tt.b)
		if result != tt.expected {
			t.Errorf("getSum(%d, %d) = %d; expected %d", tt.a, tt.b, result, tt.expected)
		}
	}
}
