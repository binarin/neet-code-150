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
	// Example 1: root = [2,1,3]
	root := &TreeNode{
		Val:   2,
		Left:  &TreeNode{Val: 1},
		Right: &TreeNode{Val: 3},
	}
	fmt.Println(isValidBST(root))
}

func isValidBST(root *TreeNode) bool {
	var recur func(*TreeNode) (int, int, bool)
	recur = func(n *TreeNode) (treeMin, treeMax int, treeValid bool) {
		if n == nil {
			return math.MaxInt, math.MinInt, true
		}

		treeMin, treeMax, treeValid = n.Val, n.Val, false

		defer (func() {
			fmt.Printf("Result for %d: %d, %d, %t\n", n.Val, treeMin, treeMax, treeValid)
		})()

		lMin, lMax, lValid := recur(n.Left)
		if !lValid || lMax >= n.Val {
			return
		}

		rMin, rMax, rValid := recur(n.Right)
		if !rValid || n.Val >= rMin {
			return
		}
		treeMin, treeMax, treeValid = min(lMin, treeMin), max(rMax, treeMax), true
		return

	}
	_, _, valid := recur(root)
	return valid
}
