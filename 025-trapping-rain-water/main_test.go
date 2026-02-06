package main

import (
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTrap(t *testing.T) {
	assert.Equal(t, 6, trap([]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}))
	assert.Equal(t, 9, trap([]int{4, 2, 0, 3, 2, 5}))
}

func TestTrapRandom(t *testing.T) {
	for i := 0; i < 10; i++ {
		length := rand.Intn(21) + 20
		height := make([]int, length)
		for j := 0; j < length; j++ {
			height[j] = rand.Intn(41)
		}
		expected := trapBf(height)
		assert.Equalf(t, expected, trap(height), "height: %v", height)
	}
}
