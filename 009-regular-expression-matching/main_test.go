package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsMatch(t *testing.T) {
	assert.Equal(t, true, isMatch("a", "ab*"))
	assert.Equal(t, false, isMatch("aaa", "ab*a"))
	assert.Equal(t, false, isMatch("aa", "a"))
	assert.Equal(t, true, isMatch("aa", "aa"))
	assert.Equal(t, true, isMatch("aa", "a*"))
	assert.Equal(t, true, isMatch("ab", ".*"))
	assert.Equal(t, false, isMatch("ab", ".*c"))
	assert.Equal(t, true, isMatch("abc", ".*c"))
	assert.Equal(t, true, isMatch("aaa", "a*a"))
}
