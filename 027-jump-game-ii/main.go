package main

import (
	"fmt"
	"math"
	"slices"
)

func main() {
	fmt.Println(jump([]int{2, 3, 1, 1, 4}))
}

func jump(nums []int) int {
	var bests []int = slices.Repeat([]int{math.MaxInt}, len(nums))

	var recur func(int, int)
	recur = func(pos int, jumps int) {
		if pos >= len(nums) {
			return
		}
		if bests[pos] <= jumps {
			return
		}
		bests[pos] = jumps
		for i := 1; i <= nums[pos]; i++ {
			jump_target := pos + i
			if jump_target >= len(nums) {
				break
			}
			recur(jump_target, jumps+1)
		}
	}
	recur(0, 0)

	return bests[len(nums)-1]
}
