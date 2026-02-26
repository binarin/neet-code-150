package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsValidBST(t *testing.T) {
	// Example 1: root = [2,1,3] -> true
	root1 := &TreeNode{
		Val:   2,
		Left:  &TreeNode{Val: 1},
		Right: &TreeNode{Val: 3},
	}
	assert.Equal(t, true, isValidBST(root1))

	// Example 2: root = [5,1,4,null,null,3,6] -> false
	root2 := &TreeNode{
		Val:  5,
		Left: &TreeNode{Val: 1},
		Right: &TreeNode{
			Val:   4,
			Left:  &TreeNode{Val: 3},
			Right: &TreeNode{Val: 6},
		},
	}
	assert.Equal(t, false, isValidBST(root2))
}
