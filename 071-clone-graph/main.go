// https://leetcode.com/problems/clone-graph/
// Difficulty: Medium

package main

import "fmt"

// Definition for a Node.
type Node struct {
	Val       int
	Neighbors []*Node
}

func cloneGraph(node *Node) *Node {
	return nil
}

func main() {
	// Example 1: adjList = [[2,4],[1,3],[2,4],[1,3]]
	// Create 4 nodes
	node1 := &Node{Val: 1}
	node2 := &Node{Val: 2}
	node3 := &Node{Val: 3}
	node4 := &Node{Val: 4}

	// Set up neighbors
	node1.Neighbors = []*Node{node2, node4}
	node2.Neighbors = []*Node{node1, node3}
	node3.Neighbors = []*Node{node2, node4}
	node4.Neighbors = []*Node{node1, node3}

	result := cloneGraph(node1)
	fmt.Println(result)
}
