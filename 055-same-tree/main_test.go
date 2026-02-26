package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsSameTree(t *testing.T) {
	// Example 1: p = [1,2,3], q = [1,2,3] -> true
	p1 := &TreeNode{Val: 1, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 3}}
	q1 := &TreeNode{Val: 1, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 3}}
	assert.Equal(t, true, isSameTree(p1, q1))

	// Example 2: p = [1,2], q = [1,null,2] -> false
	p2 := &TreeNode{Val: 1, Left: &TreeNode{Val: 2}}
	q2 := &TreeNode{Val: 1, Right: &TreeNode{Val: 2}}
	assert.Equal(t, false, isSameTree(p2, q2))

	// Example 3: p = [1,2,1], q = [1,1,2] -> false
	p3 := &TreeNode{Val: 1, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 1}}
	q3 := &TreeNode{Val: 1, Left: &TreeNode{Val: 1}, Right: &TreeNode{Val: 2}}
	assert.Equal(t, false, isSameTree(p3, q3))
}
