// https://leetcode.com/problems/gas-station/
// Difficulty: Medium

package main

import (
	"fmt"
	"math"
)

func canCompleteCircuit(gas []int, cost []int) int {
	surplus, minSurplus := 0, math.MaxInt
	for i := range gas {
		surplus += gas[i] - cost[i]
		minSurplus = min(minSurplus, surplus)
	}
	if surplus < 0 {
		return -1
	}
	for i := range gas {
		if minSurplus >= 0 {
			return i
		}
		minSurplus -= gas[i] - cost[i]
	}
	return -1
}

func canCompleteCircuitBf(gas []int, cost []int) int {
outer:
	for start := range gas {
		surplus := 0
		for cur := range gas {
			cur = (start + cur) % len(gas)
			surplus += gas[cur]
			surplus -= cost[cur]
			if surplus < 0 {
				continue outer
			}
		}
		return start
	}
	return -1
}

func main() {
	gas := []int{1, 2, 3, 4, 5}
	cost := []int{3, 4, 5, 1, 2}
	fmt.Println(canCompleteCircuit(gas, cost))
}
