package main

import (
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestJump(t *testing.T) {
	assert.Equal(t, 2, jump([]int{2, 3, 1, 1, 4}))
	assert.Equal(t, 2, jump([]int{2, 3, 0, 1, 4}))
}

func TestJumpRandom(t *testing.T) {
	for i := 0; i < 10; i++ {
		length := rand.Intn(20) + 1
		nums := make([]int, length)
		for j := 0; j < length; j++ {
			nums[j] = rand.Intn(5)
		}
		// Ensure we can reach the end
		nums[0] = max(nums[0], 1)

		expected := jumpBf(nums)
		actual := jump(nums)
		assert.Equal(t, expected, actual, "input: %v", nums)
	}
}
