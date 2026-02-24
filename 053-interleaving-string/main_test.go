package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsInterleave(t *testing.T) {
	assert.Equal(t, true, isInterleave("aabcc", "dbbca", "aadbbcbcac"))
	assert.Equal(t, false, isInterleave("aabcc", "dbbca", "aadbbbaccc"))
	assert.Equal(t, true, isInterleave("", "", ""))
}
