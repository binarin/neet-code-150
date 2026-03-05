// https://leetcode.com/problems/lowest-common-ancestor-of-a-binary-search-tree/
// Difficulty: Medium

package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	return nil
}

func main() {
	// Example 1: root = [6,2,8,0,4,7,9,null,null,3,5], p = 2, q = 8
	// Expected output: 6
	node3 := &TreeNode{Val: 3}
	node5 := &TreeNode{Val: 5}
	node0 := &TreeNode{Val: 0}
	node4 := &TreeNode{Val: 4, Left: node3, Right: node5}
	node7 := &TreeNode{Val: 7}
	node9 := &TreeNode{Val: 9}
	node2 := &TreeNode{Val: 2, Left: node0, Right: node4}
	node8 := &TreeNode{Val: 8, Left: node7, Right: node9}
	root := &TreeNode{Val: 6, Left: node2, Right: node8}

	result := lowestCommonAncestor(root, node2, node8)
	if result != nil {
		fmt.Println(result.Val)
	} else {
		fmt.Println("nil")
	}
}
