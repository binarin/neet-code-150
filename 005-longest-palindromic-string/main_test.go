package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsPalindrome(t *testing.T) {
	assert.Equal(t, "ccc", longestPalindrome("ccc"))
	assert.Equal(t, "bb", longestPalindrome("bb"))
	assert.Equal(t, "b", longestPalindrome("ba"))
	assert.Equal(t, "", longestPalindrome(""))
}
