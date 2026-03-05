package main

import "testing"

func TestIsPalindrome(t *testing.T) {
	tests := []struct {
		name     string
		s        string
		expected bool
	}{
		{
			name:     "Example 1: palindrome with mixed case and punctuation",
			s:        "A man, a plan, a canal: Panama",
			expected: true,
		},
		{
			name:     "Example 2: not a palindrome",
			s:        "race a car",
			expected: false,
		},
		{
			name:     "Example 3: single space (empty after filtering)",
			s:        " ",
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := isPalindrome(tt.s)
			if result != tt.expected {
				t.Errorf("isPalindrome(%q) = %v, want %v", tt.s, result, tt.expected)
			}
		})
	}
}
