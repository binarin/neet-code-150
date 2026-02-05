package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func diameterOfBinaryTree(root *TreeNode) int {
	// [2]int{diameter, max_depth}
	var recur func(root *TreeNode) [2]int
	recur = func(root *TreeNode) [2]int {
		if root == nil {
			return [2]int{0, 0}
		}
		left_dia_and_depth := recur(root.Left)
		right_dia_and_depth := recur(root.Right)
		this_dia := left_dia_and_depth[1] + right_dia_and_depth[1]
		return [2]int{max(this_dia, left_dia_and_depth[0], right_dia_and_depth[0]), 1 + max(left_dia_and_depth[1], right_dia_and_depth[1])}
	}

	return recur(root)[0]
}

func main() {
	// Example 1: root = [1,2,3,4,5], expected output: 3
	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val:   2,
			Left:  &TreeNode{Val: 4},
			Right: &TreeNode{Val: 5},
		},
		Right: &TreeNode{Val: 3},
	}
	fmt.Println(diameterOfBinaryTree(root))
}
