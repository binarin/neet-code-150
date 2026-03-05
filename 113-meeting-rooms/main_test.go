package main

import "testing"

func Test_canAttendMeetings(t *testing.T) {
	tests := []struct {
		name      string
		intervals [][]int
		want      bool
	}{
		{
			name:      "Example 1: overlapping meetings",
			intervals: [][]int{{0, 30}, {5, 10}, {15, 20}},
			want:      false,
		},
		{
			name:      "Example 2: non-overlapping meetings",
			intervals: [][]int{{7, 10}, {2, 4}},
			want:      true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := canAttendMeetings(tt.intervals); got != tt.want {
				t.Errorf("canAttendMeetings() = %v, want %v", got, tt.want)
			}
		})
	}
}
