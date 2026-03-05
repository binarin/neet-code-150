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
	// result := buildTree([]int{3, 9, 20, 15, 7}, []int{9, 3, 15, 20, 7})
	result := buildTree([]int{1, 2}, []int{1, 2})
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
	if len(preorder) == 0 {
		return nil
	}
	lookup := map[int]int{}
	for idx, val := range inorder {
		lookup[val] = idx
	}

	type sl [2]int
	var recur func(sl, sl) *TreeNode
	recur = func(pre, in sl) *TreeNode {
		if pre[1] == 0 {
			return nil
		}
		current := preorder[pre[0]]
		leftIdxAbs := lookup[current]
		leftNo := leftIdxAbs - in[0]
		// fmt.Printf("%3d %3d %3d %3d %3d\n", pre, in, current, leftIdxAbs, leftNo)
		node := new(TreeNode)
		node.Val = current
		node.Left = recur(sl{pre[0] + 1, leftNo}, sl{in[0], leftNo})
		node.Right = recur(sl{pre[0] + leftNo + 1, pre[1] - leftNo - 1}, sl{in[0] + leftNo + 1, in[1] - leftNo - 1})
		return node
	}

	return recur(sl{0, len(preorder)}, sl{0, len(inorder)})
}
