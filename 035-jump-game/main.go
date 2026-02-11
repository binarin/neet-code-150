package main

import (
	"fmt"
)

func main() {
	fmt.Println(canJump([]int{2, 3, 1, 1, 4}))
}

func canJump(nums []int) bool {
	can_reach := make([]bool, len(nums))
	can_reach[0] = true
	for i, val := range nums {
		if !can_reach[i] {
			return false
		}
		tail_length := len(nums) - i - 1
		// n=2, i=0 -> tail_length=1
		// di = min(val, 1)

		for di := 1; di <= min(val, tail_length); di++ {
			can_reach[i+di] = true
		}
	}
	return can_reach[len(nums)-1]
}
