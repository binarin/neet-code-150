package main

import "testing"

func TestIsBalanced(t *testing.T) {
	tests := []struct {
		name     string
		root     []*int
		expected bool
	}{
		{
			name:     "Example 1: balanced tree",
			root:     []*int{intPtr(3), intPtr(9), intPtr(20), nil, nil, intPtr(15), intPtr(7)},
			expected: true,
		},
		{
			name:     "Example 2: unbalanced tree",
			root:     []*int{intPtr(1), intPtr(2), intPtr(2), intPtr(3), intPtr(3), nil, nil, intPtr(4), intPtr(4)},
			expected: false,
		},
		{
			name:     "Example 3: empty tree",
			root:     []*int{},
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			root := buildTree(tt.root)
			result := isBalanced(root)
			if result != tt.expected {
				t.Errorf("isBalanced() = %v, expected %v", result, tt.expected)
			}
		})
	}
}
