package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestThreeSum(t *testing.T) {
	assert.ElementsMatch(t, [][]int{{-1, -1, 2}, {-1, 0, 1}}, threeSum([]int{-1, 0, 1, 2, -1, -4}))
	assert.ElementsMatch(t, [][]int{}, threeSum([]int{0, 1, 1}))
	assert.ElementsMatch(t, [][]int{{0, 0, 0}}, threeSum([]int{0, 0, 0}))
}
