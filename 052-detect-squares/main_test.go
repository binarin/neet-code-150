package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDetectSquares(t *testing.T) {
	// Example 1:
	// Input: ["DetectSquares", "add", "add", "add", "count", "count", "add", "count"]
	//        [[], [[3, 10]], [[11, 2]], [[3, 2]], [[11, 10]], [[14, 8]], [[11, 2]], [[11, 10]]]
	// Output: [null, null, null, null, 1, 0, null, 2]
	detectSquares := Constructor()
	detectSquares.Add([]int{3, 10})
	detectSquares.Add([]int{11, 2})
	detectSquares.Add([]int{3, 2})
	assert.Equal(t, 1, detectSquares.Count([]int{11, 10}))
	assert.Equal(t, 0, detectSquares.Count([]int{14, 8}))
	detectSquares.Add([]int{11, 2}) // Adding duplicate points is allowed
	assert.Equal(t, 2, detectSquares.Count([]int{11, 10}))
}
