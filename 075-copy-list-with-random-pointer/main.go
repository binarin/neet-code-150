// https://leetcode.com/problems/copy-list-with-random-pointer/
// Difficulty: Medium

package main

import "fmt"

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func copyRandomList(head *Node) *Node {
	newCons := map[*Node]*Node{}
	var ensure func(*Node) *Node
	ensure = func(n *Node) *Node {
		if n == nil {
			return nil
		}
		if cons, ok := newCons[n]; ok {
			return cons
		}
		newCons[n] = &Node{Val: n.Val}
		newCons[n].Next = ensure(n.Next)
		newCons[n].Random = ensure(n.Random)
		return newCons[n]
	}
	return ensure(head)
}

// buildList creates a linked list from a slice of [val, randomIndex] pairs
// where randomIndex is the 0-based index of the node that Random points to, or -1 for nil
func buildList(pairs [][2]int) *Node {
	if len(pairs) == 0 {
		return nil
	}

	nodes := make([]*Node, len(pairs))
	for i, p := range pairs {
		nodes[i] = &Node{Val: p[0]}
	}

	for i := 0; i < len(nodes); i++ {
		if i < len(nodes)-1 {
			nodes[i].Next = nodes[i+1]
		}
		randomIdx := pairs[i][1]
		if randomIdx >= 0 {
			nodes[i].Random = nodes[randomIdx]
		}
	}

	return nodes[0]
}

// listToString converts a linked list to string representation
func listToString(head *Node) string {
	if head == nil {
		return "[]"
	}

	// First pass: build index map
	nodeToIdx := make(map[*Node]int)
	idx := 0
	for curr := head; curr != nil; curr = curr.Next {
		nodeToIdx[curr] = idx
		idx++
	}

	// Second pass: build string
	result := "["
	for curr := head; curr != nil; curr = curr.Next {
		if curr != head {
			result += ","
		}
		randomStr := "null"
		if curr.Random != nil {
			randomStr = fmt.Sprintf("%d", nodeToIdx[curr.Random])
		}
		result += fmt.Sprintf("[%d,%s]", curr.Val, randomStr)
	}
	result += "]"
	return result
}

func main() {
	// Example 1: head = [[7,null],[13,0],[11,4],[10,2],[1,0]]
	head := buildList([][2]int{{7, -1}, {13, 0}, {11, 4}, {10, 2}, {1, 0}})
	result := copyRandomList(head)
	fmt.Println(listToString(result))
}
