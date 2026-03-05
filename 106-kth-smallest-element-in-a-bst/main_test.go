package main

import "testing"

func TestKthSmallest(t *testing.T) {
	tests := []struct {
		name     string
		root     *TreeNode
		k        int
		expected int
	}{
		{
			name: "Example 1: k=1 in [3,1,4,null,2]",
			root: &TreeNode{
				Val: 3,
				Left: &TreeNode{
					Val:   1,
					Right: &TreeNode{Val: 2},
				},
				Right: &TreeNode{Val: 4},
			},
			k:        1,
			expected: 1,
		},
		{
			name: "Example 2: k=3 in [5,3,6,2,4,null,null,1]",
			root: &TreeNode{
				Val: 5,
				Left: &TreeNode{
					Val: 3,
					Left: &TreeNode{
						Val:  2,
						Left: &TreeNode{Val: 1},
					},
					Right: &TreeNode{Val: 4},
				},
				Right: &TreeNode{Val: 6},
			},
			k:        3,
			expected: 3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := kthSmallest(tt.root, tt.k)
			if result != tt.expected {
				t.Errorf("kthSmallest() = %v, expected %v", result, tt.expected)
			}
		})
	}
}
