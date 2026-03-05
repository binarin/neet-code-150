package main

import "testing"

func Test_isAnagram(t *testing.T) {
	tests := []struct {
		name string
		s    string
		t    string
		want bool
	}{
		{
			name: "example 1",
			s:    "anagram",
			t:    "nagaram",
			want: true,
		},
		{
			name: "example 2",
			s:    "rat",
			t:    "car",
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isAnagram(tt.s, tt.t); got != tt.want {
				t.Errorf("isAnagram() = %v, want %v", got, tt.want)
			}
		})
	}
}
