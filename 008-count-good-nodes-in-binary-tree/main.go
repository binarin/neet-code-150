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

func goodNodes(root *TreeNode) int {
	var recur func(cur *TreeNode, max_seen int) int
	recur = func(cur *TreeNode, max_seen int) int {
		if cur == nil {
			return 0
		}
		good_count := 0
		if cur.Val >= max_seen {
			good_count++
		}
		new_max := max(max_seen, cur.Val)
		good_count += recur(cur.Left, new_max)
		good_count += recur(cur.Right, new_max)
		return good_count
	}
	return recur(root, math.MinInt)
}

func intPtr(v int) *int {
	return &v
}

func buildTree(vals []*int) *TreeNode {
	if len(vals) == 0 || vals[0] == nil {
		return nil
	}
	root := &TreeNode{Val: *vals[0]}
	queue := []*TreeNode{root}
	i := 1
	for len(queue) > 0 && i < len(vals) {
		node := queue[0]
		queue = queue[1:]
		if i < len(vals) && vals[i] != nil {
			node.Left = &TreeNode{Val: *vals[i]}
			queue = append(queue, node.Left)
		}
		i++
		if i < len(vals) && vals[i] != nil {
			node.Right = &TreeNode{Val: *vals[i]}
			queue = append(queue, node.Right)
		}
		i++
	}
	return root
}

func main() {
	root := buildTree([]*int{intPtr(3), intPtr(1), intPtr(4), intPtr(3), nil, intPtr(1), intPtr(5)})
	fmt.Println(goodNodes(root))
}
