package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(lastStoneWeightBf([]int{2, 7, 4, 1, 8, 1}))
}

func lastStoneWeight(stones []int) int {
	return lastStoneWeightBf(stones)
}

func lastStoneWeightBf(stones []int) int {
	if len(stones) == 0 {
		return 0
	}
	if len(stones) == 1 {
		return stones[0]
	}
	sort.Slice(stones, func(i, j int) bool {
		return stones[i] > stones[j]
	})
	fmt.Println(stones)
	x, y := stones[0], stones[1]
	if x == y {
		return lastStoneWeightBf(stones[2:])
	}
	rem := max(x-y, y-x)
	stones = stones[1:]
	stones[0] = rem
	return lastStoneWeightBf(stones)
}
