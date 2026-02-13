package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinDistance(t *testing.T) {
	assert.Equal(t, 3, minDistance("park", "spake"))
	assert.Equal(t, 3, minDistance("horse", "ros"))
	assert.Equal(t, 5, minDistance("intention", "execution"))
}
