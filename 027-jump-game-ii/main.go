package main

import (
	"fmt"
	"math"
	"slices"
)

func main() {
	fmt.Println(jump([]int{1, 3, 4}))
}

func jump(nums []int) int {
	var bests []int = slices.Repeat([]int{math.MaxInt}, len(nums))
	bests[0] = 0
	for i := 0; i < len(nums)-1; i++ {
		if bests[i] == math.MaxInt {
			return -1
		}
		for j := 1; j <= nums[i]; j++ {
			if i+j >= len(nums) {
				break
			}
			bests[i+j] = min(bests[i+j], bests[i]+1)
		}
	}
	if bests[len(bests)-1] == math.MaxInt {
		return -1
	}
	return bests[len(bests)-1]
}

func jumpBf(nums []int) int {
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

	if bests[len(bests)-1] == math.MaxInt {
		return -1
	}
	return bests[len(nums)-1]
}
