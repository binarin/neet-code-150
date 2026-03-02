package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindWords(t *testing.T) {
	board1 := [][]byte{
		{'o', 'a', 'a', 'n'},
		{'e', 't', 'a', 'e'},
		{'i', 'h', 'k', 'r'},
		{'i', 'f', 'l', 'v'},
	}
	words1 := []string{"oath", "pea", "eat", "rain"}
	assert.ElementsMatch(t, []string{"eat", "oath"}, findWords(board1, words1))

	board2 := [][]byte{
		{'a', 'b'},
		{'c', 'd'},
	}
	words2 := []string{"abcb"}
	assert.ElementsMatch(t, []string{}, findWords(board2, words2))

	board3 := [][]byte{
		{'o', 'a', 'a', 'n'},
		{'e', 't', 'a', 'e'},
		{'i', 'h', 'k', 'r'},
		{'i', 'f', 'l', 'v'},
	}
	words3 := []string{"oath", "pea", "eat", "rain", "hklf", "hf"}
	assert.ElementsMatch(t, []string{"oath", "eat", "hklf", "hf"}, findWords(board3, words3))
}
