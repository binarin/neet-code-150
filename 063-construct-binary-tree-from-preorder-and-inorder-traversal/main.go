// https://leetcode.com/problems/construct-binary-tree-from-preorder-and-inorder-traversal/
// Difficulty: Medium
package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	// Example 1: preorder = [3,9,20,15,7], inorder = [9,3,15,20,7]
	// Output: [3,9,20,null,null,15,7]
	result := buildTree([]int{3, 9, 20, 15, 7}, []int{9, 3, 15, 20, 7})
	printTree(result)
}

func printTree(root *TreeNode) {
	if root == nil {
		fmt.Println("nil")
		return
	}
	// BFS to print level order
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		if node == nil {
			fmt.Print("null ")
		} else {
			fmt.Print(node.Val, " ")
			queue = append(queue, node.Left, node.Right)
		}
	}
	fmt.Println()
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	return nil
}
