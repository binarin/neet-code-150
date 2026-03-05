package main

import "testing"

func Test_alienOrder(t *testing.T) {
	tests := []struct {
		name  string
		words []string
		want  string
	}{
		{
			name:  "example 1",
			words: []string{"wrt", "wrf", "er", "ett", "rftt"},
			want:  "wertf",
		},
		{
			name:  "example 2",
			words: []string{"z", "x"},
			want:  "zx",
		},
		{
			name:  "example 3 - invalid order",
			words: []string{"z", "x", "z"},
			want:  "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := alienOrder(tt.words); got != tt.want {
				t.Errorf("alienOrder() = %v, want %v", got, tt.want)
			}
		})
	}
}
