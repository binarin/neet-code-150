package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOrangesRotting(t *testing.T) {
	assert.Equal(t, 4, orangesRotting([][]int{{2, 1, 1}, {1, 1, 0}, {0, 1, 1}}))
	assert.Equal(t, -1, orangesRotting([][]int{{2, 1, 1}, {0, 1, 1}, {1, 0, 1}}))
	assert.Equal(t, 0, orangesRotting([][]int{{0, 2}}))
}
