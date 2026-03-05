package main

import "testing"

func TestHammingWeight(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want int
	}{
		{
			name: "Example 1: n = 11",
			n:    11,
			want: 3,
		},
		{
			name: "Example 2: n = 128",
			n:    128,
			want: 1,
		},
		{
			name: "Example 3: n = 2147483645",
			n:    2147483645,
			want: 30,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := hammingWeight(tt.n)
			if got != tt.want {
				t.Errorf("hammingWeight(%d) = %d, want %d", tt.n, got, tt.want)
			}
		})
	}
}
