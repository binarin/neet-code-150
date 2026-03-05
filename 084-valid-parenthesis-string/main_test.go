package main

import "testing"

func Test_checkValidString(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want bool
	}{
		{
			name: "Example 1: simple parentheses",
			s:    "()",
			want: true,
		},
		{
			name: "Example 2: star as empty",
			s:    "(*)",
			want: true,
		},
		{
			name: "Example 3: star as left paren",
			s:    "(*))",
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := checkValidString(tt.s); got != tt.want {
				t.Errorf("checkValidString() = %v, want %v", got, tt.want)
			}
		})
	}
}
