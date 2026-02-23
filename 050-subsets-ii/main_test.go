package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSubsetsWithDup(t *testing.T) {
	assert.ElementsMatch(t, [][]int{{}, {1}, {1, 2}, {1, 2, 2}, {2}, {2, 2}}, subsetsWithDup([]int{1, 2, 2}))
	assert.ElementsMatch(t, [][]int{{}, {0}}, subsetsWithDup([]int{0}))
}
