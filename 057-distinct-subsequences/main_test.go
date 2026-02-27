package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNumDistinct(t *testing.T) {
	assert.Equal(t, 3, numDistinct("rabbbit", "rabbit"))
	assert.Equal(t, 5, numDistinct("babgbag", "bag"))
}
