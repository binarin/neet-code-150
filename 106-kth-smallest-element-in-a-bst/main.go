// https://leetcode.com/problems/kth-smallest-element-in-a-bst/
// Difficulty: Medium

package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func kthSmallest(root *TreeNode, k int) int {
	return 0
}

func main() {
	// Example 1: root = [3,1,4,null,2], k = 1
	// Tree structure:
	//     3
	//    / \
	//   1   4
	//    \
	//     2
	root := &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val:   1,
			Right: &TreeNode{Val: 2},
		},
		Right: &TreeNode{Val: 4},
	}
	k := 1
	result := kthSmallest(root, k)
	fmt.Println(result) // Expected: 1
}
