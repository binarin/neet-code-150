package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxDepth(t *testing.T) {
	// Example 1: root = [3,9,20,null,null,15,7], expected = 3
	assert.Equal(t, 3, maxDepth(buildTree([]*int{intPtr(3), intPtr(9), intPtr(20), nil, nil, intPtr(15), intPtr(7)})))
	// Example 2: root = [1,null,2], expected = 2
	assert.Equal(t, 2, maxDepth(buildTree([]*int{intPtr(1), nil, intPtr(2)})))
}
