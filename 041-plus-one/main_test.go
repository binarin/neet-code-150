package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPlusOne(t *testing.T) {
	assert.Equal(t, []int{1, 2, 4}, plusOne([]int{1, 2, 3}))
	assert.Equal(t, []int{4, 3, 2, 2}, plusOne([]int{4, 3, 2, 1}))
	assert.Equal(t, []int{1, 0}, plusOne([]int{9}))
}
