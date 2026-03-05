// https://leetcode.com/problems/binary-tree-right-side-view/description/?envType=problem-list-v2&envId=plakya4j
// Difficulty: Medium

package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rightSideView(root *TreeNode) []int {
	return nil
}

func main() {
	// Example 1: root = [1,2,3,null,5,null,4]
	//       1
	//      / \
	//     2   3
	//      \   \
	//       5   4
	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val:   2,
			Right: &TreeNode{Val: 5},
		},
		Right: &TreeNode{
			Val:   3,
			Right: &TreeNode{Val: 4},
		},
	}
	fmt.Println(rightSideView(root)) // Expected: [1,3,4]
}
