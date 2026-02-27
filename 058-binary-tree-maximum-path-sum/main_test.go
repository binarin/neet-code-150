package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxPathSum(t *testing.T) {
	// Example 1: root = [1,2,3], expected = 6
	root1 := &TreeNode{
		Val:   1,
		Left:  &TreeNode{Val: 2},
		Right: &TreeNode{Val: 3},
	}
	assert.Equal(t, 6, maxPathSum(root1))

	// Example 2: root = [-10,9,20,null,null,15,7], expected = 42
	root2 := &TreeNode{
		Val:  -10,
		Left: &TreeNode{Val: 9},
		Right: &TreeNode{
			Val:   20,
			Left:  &TreeNode{Val: 15},
			Right: &TreeNode{Val: 7},
		},
	}
	assert.Equal(t, 42, maxPathSum(root2))
}
