package main

import "testing"

func Test_wordBreak(t *testing.T) {
	tests := []struct {
		name     string
		s        string
		wordDict []string
		want     bool
	}{
		{
			name:     "example 1 - leetcode",
			s:        "leetcode",
			wordDict: []string{"leet", "code"},
			want:     true,
		},
		{
			name:     "example 2 - applepenapple",
			s:        "applepenapple",
			wordDict: []string{"apple", "pen"},
			want:     true,
		},
		{
			name:     "example 3 - catsandog",
			s:        "catsandog",
			wordDict: []string{"cats", "dog", "sand", "and", "cat"},
			want:     false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := wordBreak(tt.s, tt.wordDict); got != tt.want {
				t.Errorf("wordBreak() = %v, want %v", got, tt.want)
			}
		})
	}
}
