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
	newNodes := map[*Node]*Node{}
	ensure := func(oldNode *Node) *Node {
		if n, ok := newNodes[oldNode]; ok {
			return n
		}
		newNodes[oldNode] = &Node{Val: oldNode.Val}
		return newNodes[oldNode]
	}
	queue := []*Node{node}
	visited := map[*Node]bool{node: true}

	for len(queue) > 0 {
		cur := queue[0]
		fmt.Println(cur)
		queue = queue[1:]
		newNode := ensure(cur)
		for _, neighbor := range cur.Neighbors {
			if _, ok := visited[neighbor]; !ok {
				queue = append(queue, neighbor)
				visited[neighbor] = true
			}
			newNeighbor := ensure(neighbor)
			newNode.Neighbors = append(newNode.Neighbors, newNeighbor)
		}
	}

	return newNodes[node]
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
