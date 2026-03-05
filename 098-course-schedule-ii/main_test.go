package main

import (
	"testing"
)

func isValidOrder(numCourses int, prerequisites [][]int, order []int) bool {
	if len(order) != numCourses {
		return false
	}

	// Check all courses are present exactly once
	seen := make(map[int]bool)
	for _, course := range order {
		if course < 0 || course >= numCourses {
			return false
		}
		if seen[course] {
			return false
		}
		seen[course] = true
	}

	// Build position map
	pos := make(map[int]int)
	for i, course := range order {
		pos[course] = i
	}

	// Check all prerequisites are satisfied
	for _, prereq := range prerequisites {
		course, pre := prereq[0], prereq[1]
		if pos[pre] >= pos[course] {
			return false
		}
	}

	return true
}

func TestFindOrder(t *testing.T) {
	tests := []struct {
		name          string
		numCourses    int
		prerequisites [][]int
		wantEmpty     bool // true if expecting empty array (no valid order)
	}{
		{
			name:          "Example 1: 2 courses with one prerequisite",
			numCourses:    2,
			prerequisites: [][]int{{1, 0}},
			wantEmpty:     false,
		},
		{
			name:          "Example 2: 4 courses with multiple prerequisites",
			numCourses:    4,
			prerequisites: [][]int{{1, 0}, {2, 0}, {3, 1}, {3, 2}},
			wantEmpty:     false,
		},
		{
			name:          "Example 3: 1 course with no prerequisites",
			numCourses:    1,
			prerequisites: [][]int{},
			wantEmpty:     false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := findOrder(tt.numCourses, tt.prerequisites)
			if tt.wantEmpty {
				if len(got) != 0 {
					t.Errorf("findOrder() = %v, want empty array", got)
				}
			} else {
				if !isValidOrder(tt.numCourses, tt.prerequisites, got) {
					t.Errorf("findOrder() = %v, not a valid course order", got)
				}
			}
		})
	}
}
