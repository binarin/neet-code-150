package main

import "testing"

func TestLowestCommonAncestor(t *testing.T) {
	tests := []struct {
		name     string
		buildTree func() (*TreeNode, *TreeNode, *TreeNode)
		expected int
	}{
		{
			name: "Example 1: p=2, q=8, LCA=6",
			buildTree: func() (*TreeNode, *TreeNode, *TreeNode) {
				// root = [6,2,8,0,4,7,9,null,null,3,5], p = 2, q = 8
				node3 := &TreeNode{Val: 3}
				node5 := &TreeNode{Val: 5}
				node0 := &TreeNode{Val: 0}
				node4 := &TreeNode{Val: 4, Left: node3, Right: node5}
				node7 := &TreeNode{Val: 7}
				node9 := &TreeNode{Val: 9}
				node2 := &TreeNode{Val: 2, Left: node0, Right: node4}
				node8 := &TreeNode{Val: 8, Left: node7, Right: node9}
				root := &TreeNode{Val: 6, Left: node2, Right: node8}
				return root, node2, node8
			},
			expected: 6,
		},
		{
			name: "Example 2: p=2, q=4, LCA=2",
			buildTree: func() (*TreeNode, *TreeNode, *TreeNode) {
				// root = [6,2,8,0,4,7,9,null,null,3,5], p = 2, q = 4
				node3 := &TreeNode{Val: 3}
				node5 := &TreeNode{Val: 5}
				node0 := &TreeNode{Val: 0}
				node4 := &TreeNode{Val: 4, Left: node3, Right: node5}
				node7 := &TreeNode{Val: 7}
				node9 := &TreeNode{Val: 9}
				node2 := &TreeNode{Val: 2, Left: node0, Right: node4}
				node8 := &TreeNode{Val: 8, Left: node7, Right: node9}
				root := &TreeNode{Val: 6, Left: node2, Right: node8}
				return root, node2, node4
			},
			expected: 2,
		},
		{
			name: "Example 3: p=2, q=1, LCA=2",
			buildTree: func() (*TreeNode, *TreeNode, *TreeNode) {
				// root = [2,1], p = 2, q = 1
				node1 := &TreeNode{Val: 1}
				root := &TreeNode{Val: 2, Left: node1}
				return root, root, node1
			},
			expected: 2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			root, p, q := tt.buildTree()
			result := lowestCommonAncestor(root, p, q)
			if result == nil {
				t.Errorf("lowestCommonAncestor() = nil, expected %d", tt.expected)
				return
			}
			if result.Val != tt.expected {
				t.Errorf("lowestCommonAncestor() = %d, expected %d", result.Val, tt.expected)
			}
		})
	}
}
