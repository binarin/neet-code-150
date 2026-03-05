package main

import (
	"reflect"
	"testing"
)

// Helper function to convert tree to level-order slice for comparison
func treeToSlice(root *TreeNode) []any {
	if root == nil {
		return []any{}
	}
	result := []any{}
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		if node == nil {
			result = append(result, nil)
		} else {
			result = append(result, node.Val)
			queue = append(queue, node.Left, node.Right)
		}
	}
	// Trim trailing nils
	for len(result) > 0 && result[len(result)-1] == nil {
		result = result[:len(result)-1]
	}
	return result
}

func TestInvertTree(t *testing.T) {
	tests := []struct {
		name     string
		root     *TreeNode
		expected []any
	}{
		{
			name: "Example 1: [4,2,7,1,3,6,9] -> [4,7,2,9,6,3,1]",
			root: &TreeNode{
				Val: 4,
				Left: &TreeNode{
					Val:   2,
					Left:  &TreeNode{Val: 1},
					Right: &TreeNode{Val: 3},
				},
				Right: &TreeNode{
					Val:   7,
					Left:  &TreeNode{Val: 6},
					Right: &TreeNode{Val: 9},
				},
			},
			expected: []any{4, 7, 2, 9, 6, 3, 1},
		},
		{
			name: "Example 2: [2,1,3] -> [2,3,1]",
			root: &TreeNode{
				Val:   2,
				Left:  &TreeNode{Val: 1},
				Right: &TreeNode{Val: 3},
			},
			expected: []any{2, 3, 1},
		},
		{
			name:     "Example 3: [] -> []",
			root:     nil,
			expected: []any{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := invertTree(tt.root)
			resultSlice := treeToSlice(result)
			if !reflect.DeepEqual(resultSlice, tt.expected) {
				t.Errorf("got %v, want %v", resultSlice, tt.expected)
			}
		})
	}
}
