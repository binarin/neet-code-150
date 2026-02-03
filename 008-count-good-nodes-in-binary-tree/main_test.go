package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGoodNodes(t *testing.T) {
	// Example 1: root = [3,1,4,3,null,1,5], expected = 4
	assert.Equal(t, 4, goodNodes(buildTree([]*int{intPtr(3), intPtr(1), intPtr(4), intPtr(3), nil, intPtr(1), intPtr(5)})))
	// Example 2: root = [3,3,null,4,2], expected = 3
	assert.Equal(t, 3, goodNodes(buildTree([]*int{intPtr(3), intPtr(3), nil, intPtr(4), intPtr(2)})))
	// Example 3: root = [1], expected = 1
	assert.Equal(t, 1, goodNodes(buildTree([]*int{intPtr(1)})))
}
