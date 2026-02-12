package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsSubtree(t *testing.T) {
	assert.Equal(t, true, isSubtree(buildTree([]int{3, 4, 5, 1, 2}), buildTree([]int{4, 1, 2})))
	assert.Equal(t, false, isSubtree(buildTree([]int{3, 4, 5, 1, 2, null, null, null, null, 0}), buildTree([]int{4, 1, 2})))
	assert.Equal(t, true, isSubtree(
		buildTree([]int{1, null, 1, null, 1, null, 1, null, 1, null, 1, null, 1, null, 1, null, 1, null, 1, null, 1, 2}),
		buildTree([]int{1, null, 1, null, 1, null, 1, null, 1, null, 1, 2}),
	))
}
