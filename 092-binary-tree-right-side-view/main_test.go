package main

import (
	"reflect"
	"testing"
)

func TestRightSideView(t *testing.T) {
	tests := []struct {
		name     string
		root     *TreeNode
		expected []int
	}{
		{
			name: "Example 1: root = [1,2,3,null,5,null,4]",
			root: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val:   2,
					Right: &TreeNode{Val: 5},
				},
				Right: &TreeNode{
					Val:   3,
					Right: &TreeNode{Val: 4},
				},
			},
			expected: []int{1, 3, 4},
		},
		{
			name: "Example 2: root = [1,2,3,4,null,null,null,5]",
			root: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val:  4,
						Left: &TreeNode{Val: 5},
					},
				},
				Right: &TreeNode{Val: 3},
			},
			expected: []int{1, 3, 4, 5},
		},
		{
			name: "Example 3: root = [1,null,3]",
			root: &TreeNode{
				Val:   1,
				Right: &TreeNode{Val: 3},
			},
			expected: []int{1, 3},
		},
		{
			name:     "Example 4: root = []",
			root:     nil,
			expected: []int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := rightSideView(tt.root)
			if result == nil {
				result = []int{}
			}
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("rightSideView() = %v, expected %v", result, tt.expected)
			}
		})
	}
}
