// https://leetcode.com/problems/balanced-binary-tree/
// Difficulty: Easy

package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isBalanced(root *TreeNode) bool {
	abs := func(a int) int {
		if a < 0 {
			return -a
		} else {
			return a
		}
	}
	var recur func(*TreeNode) (int, bool)
	recur = func(node *TreeNode) (int, bool) {
		if node == nil {
			return 0, true
		}
		lDepth, lOk := recur(node.Left)
		if !lOk {
			return 0, false
		}
		rDepth, rOk := recur(node.Right)
		if !rOk {
			return 0, false
		}
		if abs(lDepth-rDepth) > 1 {
			return 0, false
		}
		return 1 + max(lDepth, rDepth), true
	}
	_, balanced := recur(root)
	return balanced
}

// buildTree builds a binary tree from a level-order slice representation
// where nil values represent missing nodes
func buildTree(vals []*int) *TreeNode {
	if len(vals) == 0 || vals[0] == nil {
		return nil
	}

	root := &TreeNode{Val: *vals[0]}
	queue := []*TreeNode{root}
	i := 1

	for i < len(vals) {
		node := queue[0]
		queue = queue[1:]

		if i < len(vals) && vals[i] != nil {
			node.Left = &TreeNode{Val: *vals[i]}
			queue = append(queue, node.Left)
		}
		i++

		if i < len(vals) && vals[i] != nil {
			node.Right = &TreeNode{Val: *vals[i]}
			queue = append(queue, node.Right)
		}
		i++
	}

	return root
}

func intPtr(v int) *int {
	return &v
}

func main() {
	// Example 1: root = [3,9,20,null,null,15,7]
	root := buildTree([]*int{intPtr(3), intPtr(9), intPtr(20), nil, nil, intPtr(15), intPtr(7)})
	result := isBalanced(root)
	fmt.Println(result) // Expected: true
}
