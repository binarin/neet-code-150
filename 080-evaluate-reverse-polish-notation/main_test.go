package main

import "testing"

func Test_evalRPN(t *testing.T) {
	tests := []struct {
		name   string
		tokens []string
		want   int
	}{
		{
			name:   "Example 1: (2 + 1) * 3 = 9",
			tokens: []string{"2", "1", "+", "3", "*"},
			want:   9,
		},
		{
			name:   "Example 2: 4 + (13 / 5) = 6",
			tokens: []string{"4", "13", "5", "/", "+"},
			want:   6,
		},
		{
			name:   "Example 3: ((10 * (6 / ((9 + 3) * -11))) + 17) + 5 = 22",
			tokens: []string{"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"},
			want:   22,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := evalRPN(tt.tokens); got != tt.want {
				t.Errorf("evalRPN() = %v, want %v", got, tt.want)
			}
		})
	}
}
