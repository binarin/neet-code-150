package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLevelOrder(t *testing.T) {
	// Example 1: root = [3,9,20,null,null,15,7] -> [[3],[9,20],[15,7]]
	root1 := &TreeNode{
		Val:  3,
		Left: &TreeNode{Val: 9},
		Right: &TreeNode{
			Val:   20,
			Left:  &TreeNode{Val: 15},
			Right: &TreeNode{Val: 7},
		},
	}
	assert.Equal(t, [][]int{{3}, {9, 20}, {15, 7}}, levelOrder(root1))

	// Example 2: root = [1] -> [[1]]
	root2 := &TreeNode{Val: 1}
	assert.Equal(t, [][]int{{1}}, levelOrder(root2))

	// Example 3: root = [] -> []
	var root3 *TreeNode = nil
	assert.Equal(t, [][]int{}, levelOrder(root3))
}
