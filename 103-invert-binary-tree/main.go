// https://leetcode.com/problems/invert-binary-tree/
// Difficulty: Easy

package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func invertTree(root *TreeNode) *TreeNode {
	return nil
}

func main() {
	// Example 1: root = [4,2,7,1,3,6,9]
	//        4
	//       / \
	//      2   7
	//     / \ / \
	//    1  3 6  9
	root := &TreeNode{
		Val: 4,
		Left: &TreeNode{
			Val:   2,
			Left:  &TreeNode{Val: 1},
			Right: &TreeNode{Val: 3},
		},
		Right: &TreeNode{
			Val:   7,
			Left:  &TreeNode{Val: 6},
			Right: &TreeNode{Val: 9},
		},
	}
	result := invertTree(root)
	fmt.Println(result) // Expected: [4,7,2,9,6,3,1]
}
