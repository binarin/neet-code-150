package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDiameterOfBinaryTree(t *testing.T) {
	// Example 1: root = [1,2,3,4,5] => 3
	assert.Equal(t, 3, diameterOfBinaryTree(&TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val:   2,
			Left:  &TreeNode{Val: 4},
			Right: &TreeNode{Val: 5},
		},
		Right: &TreeNode{Val: 3},
	}))

	// Example 2: root = [1,2] => 1
	assert.Equal(t, 1, diameterOfBinaryTree(&TreeNode{
		Val:  1,
		Left: &TreeNode{Val: 2},
	}))
}
