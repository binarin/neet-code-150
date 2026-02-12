package main

import (
	"fmt"
	"math"
)

const null = math.MinInt

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	root := buildTree([]int{1, null, 1, null, 1, null, 1, null, 1, null, 1, null, 1, null, 1, null, 1, null, 1, null, 1, 2})
	subRoot := buildTree([]int{1, null, 1, null, 1, null, 1, null, 1, null, 1, 2})
	fmt.Println(isSubtree(root, subRoot))
}

func buildTree(vals []int) *TreeNode {
	if len(vals) == 0 || vals[0] == null {
		return nil
	}
	root := &TreeNode{Val: vals[0]}
	queue := []*TreeNode{root}
	i := 1
	for len(queue) > 0 && i < len(vals) {
		node := queue[0]
		queue = queue[1:]
		if i < len(vals) && vals[i] != null {
			node.Left = &TreeNode{Val: vals[i]}
			queue = append(queue, node.Left)
		}
		i++
		if i < len(vals) && vals[i] != null {
			node.Right = &TreeNode{Val: vals[i]}
			queue = append(queue, node.Right)
		}
		i++
	}
	return root
}

func (t *TreeNode) GoString() string {
	if t == nil {
		return "nil"
	}
	return fmt.Sprintf("(%d %s %s)", t.Val, t.Left.GoString(), t.Right.GoString())
}

func treeDepth(n *TreeNode) int {
	if n == nil {
		return 0
	}
	return 1 + max(treeDepth(n.Left), treeDepth(n.Right))
}

func isSameTree(n1 *TreeNode, n2 *TreeNode) bool {
	if n1 == nil && n2 == nil {
		return true
	}
	if n1 == nil || n2 == nil {
		return false
	}
	if n1.Val != n2.Val {
		return false
	}
	return isSameTree(n1.Left, n2.Left) && isSameTree(n1.Right, n2.Right)
}

func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
	if root == nil {
		return false
	}
	subDepth := treeDepth(subRoot)

	var recur func(*TreeNode) (bool, int)
	recur = func(n *TreeNode) (bool, int) {
		if n == nil {
			return false, 0
		}
		found_l, depth_l := recur(n.Left)
		found_r, depth_r := recur(n.Right)
		if found_l || found_r {
			return true, 0
		}
		cur_depth := 1 + max(depth_l, depth_r)
		if cur_depth == subDepth && isSameTree(n, subRoot) {
			return true, 0
		}
		return false, cur_depth
	}

	result, _ := recur(root)
	return result
}
