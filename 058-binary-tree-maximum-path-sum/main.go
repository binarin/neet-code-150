package main

import (
	"fmt"
	"math"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	// Example 1: root = [1,2,3]
	root := &TreeNode{
		Val:   2,
		Left:  &TreeNode{Val: -1},
		Right: nil,
	}
	fmt.Println(maxPathSum(root))
}

func maxPathSum(root *TreeNode) (maxSum int) {
	maxSum = math.MinInt
	var recur func(*TreeNode) int
	recur = func(n *TreeNode) int {
		if n == nil {
			return 0
		}
		bestLeft, bestRight := recur(n.Left), recur(n.Right)

		ownPathSum := n.Val + bestLeft + bestRight
		// fmt.Printf("At %d: bl %d, br %d, own %d, cur max %d\n", n.Val, bestLeft, bestRight, ownPathSum, maxSum)
		maxSum = max(maxSum, ownPathSum)
		return max(0, n.Val+max(bestLeft, bestRight))
	}

	recur(root)
	return
}
