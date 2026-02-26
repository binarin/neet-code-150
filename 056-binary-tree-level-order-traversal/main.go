package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	// Example 1: root = [3,9,20,null,null,15,7]
	root := &TreeNode{
		Val:  3,
		Left: &TreeNode{Val: 9},
		Right: &TreeNode{
			Val:   20,
			Left:  &TreeNode{Val: 15},
			Right: &TreeNode{Val: 7},
		},
	}
	fmt.Println(levelOrder(root))
}

func levelOrder(root *TreeNode) [][]int {
	result := [][]int{}
	if root == nil {
		return result
	}

	curLevelQueue := []*TreeNode{root}
	nextLevelQueue := []*TreeNode{}

	for len(curLevelQueue) > 0 {
		curLevelResult := []int{}
		for _, n := range curLevelQueue {
			curLevelResult = append(curLevelResult, n.Val)
			if n.Left != nil {
				nextLevelQueue = append(nextLevelQueue, n.Left)
			}
			if n.Right != nil {
				nextLevelQueue = append(nextLevelQueue, n.Right)
			}
		}
		curLevelQueue, nextLevelQueue = nextLevelQueue, nil
		result = append(result, curLevelResult)
	}

	return result
}
