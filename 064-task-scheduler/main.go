// https://leetcode.com/problems/task-scheduler/
// Difficulty: Medium

package main

import (
	"fmt"
	"maps"
	"slices"
)

func leastInterval(tasks []byte, n int) int {
	if n == 0 {
		return len(tasks)
	}
	counts := map[byte]int{}
	for _, t := range tasks {
		counts[t]++
	}

	var consumeOrder []byte
	reorder := func() {
		consumeOrder = slices.Collect(maps.Keys(counts))
		slices.SortFunc(consumeOrder, func(a, b byte) int {
			return counts[b] - counts[a]
		})
	}
	slotsUsed := 0
	var freeTaskSpots int
	for len(counts) > 0 {
		freeTaskSpots = n + 1
		emit := [26]bool{}
		reorder()
		for _, task := range consumeOrder {
			freeTaskSpots--
			emit[task-'A'] = true
			if counts[task] == 1 {
				delete(counts, task)
			} else {
				counts[task]--
			}
			if freeTaskSpots == 0 {
				break
			}
		}
		slotContents := make([]byte, 0, 26)
		for i, v := range emit {
			if !v {
				continue
			}
			slotContents = append(slotContents, byte(i)+'A')
		}
		fmt.Println(string(slotContents))
		slotsUsed++
	}

	return slotsUsed*(n+1) - freeTaskSpots

}

func main() {
	// Example 1: tasks = ["A","A","A","B","B","B"], n = 2
	// Expected output: 8
	// tasks, n := 2, []byte{'A', 'A', 'A', 'B', 'B', 'B'}
	n, tasks := 1, []byte{'B', 'C', 'D', 'A', 'A', 'A', 'A', 'G'}
	result := leastInterval(tasks, n)
	fmt.Println(result)
}
