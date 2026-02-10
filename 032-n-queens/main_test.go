package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolveNQueens(t *testing.T) {
	assert.ElementsMatch(t, [][]string{{".Q..", "...Q", "Q...", "..Q."}, {"..Q.", "Q...", "...Q", ".Q.."}}, solveNQueens(4))
	assert.ElementsMatch(t, [][]string{{"Q"}}, solveNQueens(1))
}
