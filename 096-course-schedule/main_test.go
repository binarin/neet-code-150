package main

import "testing"

func TestCanFinish(t *testing.T) {
	tests := []struct {
		name          string
		numCourses    int
		prerequisites [][]int
		want          bool
	}{
		{
			name:          "Example 1: simple dependency",
			numCourses:    2,
			prerequisites: [][]int{{1, 0}},
			want:          true,
		},
		{
			name:          "Example 2: circular dependency",
			numCourses:    2,
			prerequisites: [][]int{{1, 0}, {0, 1}},
			want:          false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := canFinish(tt.numCourses, tt.prerequisites)
			if got != tt.want {
				t.Errorf("canFinish(%d, %v) = %v, want %v", tt.numCourses, tt.prerequisites, got, tt.want)
			}
		})
	}
}
