package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestChange(t *testing.T) {
	assert.Equal(t, 4, change(5, []int{1, 2, 5}))
	assert.Equal(t, 0, change(3, []int{2}))
	assert.Equal(t, 1, change(10, []int{10}))
}

func TestChangeBf(t *testing.T) {
	assert.Equal(t, 4, changeBf(5, []int{1, 2, 5}))
	assert.Equal(t, 0, changeBf(3, []int{2}))
	assert.Equal(t, 1, changeBf(10, []int{10}))
}
