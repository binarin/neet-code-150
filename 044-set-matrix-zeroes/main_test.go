package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSetZeroes(t *testing.T) {
	matrix1 := [][]int{{1, 1, 1}, {1, 0, 1}, {1, 1, 1}}
	expected1 := [][]int{{1, 0, 1}, {0, 0, 0}, {1, 0, 1}}
	setZeroes(matrix1)
	assert.Equal(t, expected1, matrix1)

	matrix2 := [][]int{{0, 1, 2, 0}, {3, 4, 5, 2}, {1, 3, 1, 5}}
	expected2 := [][]int{{0, 0, 0, 0}, {0, 4, 5, 0}, {0, 3, 1, 0}}
	setZeroes(matrix2)
	assert.Equal(t, expected2, matrix2)
}
