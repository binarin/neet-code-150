// https://leetcode.com/problems/task-scheduler/
// Difficulty: Medium

package main

import "fmt"

func leastInterval(tasks []byte, n int) int {
	return 0
}

func main() {
	// Example 1: tasks = ["A","A","A","B","B","B"], n = 2
	// Expected output: 8
	tasks := []byte{'A', 'A', 'A', 'B', 'B', 'B'}
	n := 2
	result := leastInterval(tasks, n)
	fmt.Println(result)
}
