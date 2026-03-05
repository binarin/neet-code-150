// https://leetcode.com/problems/meeting-rooms/
// Difficulty: Easy

package main

import "fmt"

func canAttendMeetings(intervals [][]int) bool {
	return false
}

func main() {
	// Example 1: intervals = [[0,30],[5,10],[15,20]], expected: false
	intervals := [][]int{{0, 30}, {5, 10}, {15, 20}}
	fmt.Println(canAttendMeetings(intervals))
}
