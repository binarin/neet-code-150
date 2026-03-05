package main

import "testing"

// Helper function to build graph from adjacency list
func buildGraph(adjList [][]int) *Node {
	if len(adjList) == 0 {
		return nil
	}

	nodes := make([]*Node, len(adjList))
	for i := range nodes {
		nodes[i] = &Node{Val: i + 1}
	}

	for i, neighbors := range adjList {
		for _, neighbor := range neighbors {
			nodes[i].Neighbors = append(nodes[i].Neighbors, nodes[neighbor-1])
		}
	}

	return nodes[0]
}

// Helper function to convert graph to adjacency list
func graphToAdjList(node *Node) [][]int {
	if node == nil {
		return [][]int{}
	}

	visited := make(map[*Node]bool)
	nodes := make([]*Node, 0)

	// BFS to collect all nodes
	queue := []*Node{node}
	visited[node] = true
	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]
		nodes = append(nodes, curr)

		for _, neighbor := range curr.Neighbors {
			if !visited[neighbor] {
				visited[neighbor] = true
				queue = append(queue, neighbor)
			}
		}
	}

	// Sort nodes by value
	for i := 0; i < len(nodes)-1; i++ {
		for j := i + 1; j < len(nodes); j++ {
			if nodes[i].Val > nodes[j].Val {
				nodes[i], nodes[j] = nodes[j], nodes[i]
			}
		}
	}

	result := make([][]int, len(nodes))
	for i, n := range nodes {
		result[i] = make([]int, 0)
		for _, neighbor := range n.Neighbors {
			result[i] = append(result[i], neighbor.Val)
		}
	}

	return result
}

// Helper function to verify that the clone is a deep copy (different pointers)
func isDeepCopy(original, clone *Node) bool {
	if original == nil && clone == nil {
		return true
	}
	if original == nil || clone == nil {
		return false
	}

	visitedOriginal := make(map[*Node]bool)
	visitedClone := make(map[*Node]bool)

	var collectNodes func(node *Node, visited map[*Node]bool) []*Node
	collectNodes = func(node *Node, visited map[*Node]bool) []*Node {
		if node == nil || visited[node] {
			return nil
		}
		visited[node] = true
		result := []*Node{node}
		for _, neighbor := range node.Neighbors {
			result = append(result, collectNodes(neighbor, visited)...)
		}
		return result
	}

	originalNodes := collectNodes(original, visitedOriginal)
	cloneNodes := collectNodes(clone, visitedClone)

	// Check no pointers are shared
	for _, o := range originalNodes {
		for _, c := range cloneNodes {
			if o == c {
				return false
			}
		}
	}

	return true
}

func equalAdjList(a, b [][]int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if len(a[i]) != len(b[i]) {
			return false
		}
		for j := range a[i] {
			if a[i][j] != b[i][j] {
				return false
			}
		}
	}
	return true
}

func TestCloneGraph(t *testing.T) {
	tests := []struct {
		name    string
		adjList [][]int
		want    [][]int
	}{
		{
			name:    "Example 1: 4 nodes connected graph",
			adjList: [][]int{{2, 4}, {1, 3}, {2, 4}, {1, 3}},
			want:    [][]int{{2, 4}, {1, 3}, {2, 4}, {1, 3}},
		},
		{
			name:    "Example 2: single node with no neighbors",
			adjList: [][]int{{}},
			want:    [][]int{{}},
		},
		{
			name:    "Example 3: empty graph",
			adjList: [][]int{},
			want:    [][]int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			original := buildGraph(tt.adjList)
			clone := cloneGraph(original)

			got := graphToAdjList(clone)
			if !equalAdjList(got, tt.want) {
				t.Errorf("cloneGraph() = %v, want %v", got, tt.want)
			}

			if !isDeepCopy(original, clone) {
				t.Errorf("cloneGraph() did not return a deep copy")
			}
		})
	}
}
