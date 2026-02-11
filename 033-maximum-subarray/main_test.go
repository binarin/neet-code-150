package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func check(t *testing.T, nums []int) {
	expected := maxSubArrayBf(nums)
	assert.Equal(t, expected, maxSubArray(nums), fmt.Sprintf("input: %v", nums))
}

func TestMaxSubArray(t *testing.T) {
	check(t, []int{1, 2, -1, -2, 2, 1, -2, 1, 4, -5, 4})
	check(t, []int{-2, -1, -2})
	check(t, []int{-1, 0, -2})
	check(t, []int{-2, -1})
	check(t, []int{-1})
	check(t, []int{-2, 1, -3, 4, -1, 2, 1, -5, 4})
	check(t, []int{1})
	check(t, []int{5, 4, -1, 7, 8})
}
