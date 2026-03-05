package main

import "testing"

func TestCopyRandomList(t *testing.T) {
	tests := []struct {
		name   string
		pairs  [][2]int // [val, randomIndex] where -1 means null
		expect string
	}{
		{
			name:   "Example 1: [[7,null],[13,0],[11,4],[10,2],[1,0]]",
			pairs:  [][2]int{{7, -1}, {13, 0}, {11, 4}, {10, 2}, {1, 0}},
			expect: "[[7,null],[13,0],[11,4],[10,2],[1,0]]",
		},
		{
			name:   "Example 2: [[1,1],[2,1]]",
			pairs:  [][2]int{{1, 1}, {2, 1}},
			expect: "[[1,1],[2,1]]",
		},
		{
			name:   "Example 3: [[3,null],[3,0],[3,null]]",
			pairs:  [][2]int{{3, -1}, {3, 0}, {3, -1}},
			expect: "[[3,null],[3,0],[3,null]]",
		},
		{
			name:   "Empty list",
			pairs:  [][2]int{},
			expect: "[]",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			head := buildList(tt.pairs)
			result := copyRandomList(head)
			got := listToString(result)
			if got != tt.expect {
				t.Errorf("copyRandomList() = %v, want %v", got, tt.expect)
			}

			// Verify it's a deep copy (no pointers to original nodes)
			if head != nil && result != nil {
				origNodes := make(map[*Node]bool)
				for curr := head; curr != nil; curr = curr.Next {
					origNodes[curr] = true
				}
				for curr := result; curr != nil; curr = curr.Next {
					if origNodes[curr] {
						t.Errorf("Result contains pointer to original node")
					}
					if curr.Random != nil && origNodes[curr.Random] {
						t.Errorf("Result Random pointer points to original node")
					}
				}
			}
		})
	}
}
