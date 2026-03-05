package main

import "testing"

// Helper function to create a linked list with a cycle
// values: the node values
// pos: the index where the tail connects to (-1 means no cycle)
func createLinkedList(values []int, pos int) *ListNode {
	if len(values) == 0 {
		return nil
	}

	nodes := make([]*ListNode, len(values))
	for i, v := range values {
		nodes[i] = &ListNode{Val: v}
	}
	for i := 0; i < len(nodes)-1; i++ {
		nodes[i].Next = nodes[i+1]
	}
	// Create cycle if pos >= 0
	if pos >= 0 && pos < len(nodes) {
		nodes[len(nodes)-1].Next = nodes[pos]
	}
	return nodes[0]
}

func TestHasCycle(t *testing.T) {
	tests := []struct {
		name   string
		values []int
		pos    int
		want   bool
	}{
		{
			name:   "Example 1: cycle at pos 1",
			values: []int{3, 2, 0, -4},
			pos:    1,
			want:   true,
		},
		{
			name:   "Example 2: cycle at pos 0",
			values: []int{1, 2},
			pos:    0,
			want:   true,
		},
		{
			name:   "Example 3: no cycle",
			values: []int{1},
			pos:    -1,
			want:   false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			head := createLinkedList(tt.values, tt.pos)
			got := hasCycle(head)
			if got != tt.want {
				t.Errorf("hasCycle() = %v, want %v", got, tt.want)
			}
		})
	}
}
