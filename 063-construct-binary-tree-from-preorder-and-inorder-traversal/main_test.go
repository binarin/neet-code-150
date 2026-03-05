package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func treesEqual(a, b *TreeNode) bool {
	if a == nil && b == nil {
		return true
	}
	if a == nil || b == nil {
		return false
	}
	return a.Val == b.Val && treesEqual(a.Left, b.Left) && treesEqual(a.Right, b.Right)
}

func TestBuildTree(t *testing.T) {
	// Example 1: preorder = [3,9,20,15,7], inorder = [9,3,15,20,7]
	// Output: [3,9,20,null,null,15,7]
	expected1 := &TreeNode{
		Val:  3,
		Left: &TreeNode{Val: 9},
		Right: &TreeNode{
			Val:   20,
			Left:  &TreeNode{Val: 15},
			Right: &TreeNode{Val: 7},
		},
	}
	assert.True(t, treesEqual(expected1, buildTree([]int{3, 9, 20, 15, 7}, []int{9, 3, 15, 20, 7})))

	// Example 2: preorder = [-1], inorder = [-1]
	// Output: [-1]
	expected2 := &TreeNode{Val: -1}
	assert.True(t, treesEqual(expected2, buildTree([]int{-1}, []int{-1})))
}
