// https://leetcode.com/problems/word-search-ii/
// Difficulty: Hard
package main

import (
	"fmt"
	"strings"
)

func main() {
	// board := [][]byte{
	// 	{'o', 'a', 'a', 'n'},
	// 	{'e', 't', 'a', 'e'},
	// 	{'i', 'h', 'k', 'r'},
	// 	{'i', 'f', 'l', 'v'},
	// }
	// words := []string{"oath", "pea", "eat", "rain"}
	// fmt.Println(findWords(board, words))

	// board2 := [][]byte{
	// 	{'o', 'a', 'a', 'n'},
	// 	{'e', 't', 'a', 'e'},
	// 	{'i', 'h', 'k', 'r'},
	// 	{'i', 'f', 'l', 'v'},
	// }
	// words2 := []string{"oath", "pea", "eat", "rain", "hklf", "hf"}
	// fmt.Println(findWords(board2, words2))

	board3 := [][]byte{
		{'a', 'b', 'c'},
		{'a', 'e', 'd'},
		{'a', 'f', 'g'},
	}
	words3 := []string{"abcdefg", "gfedcbaaa", "eaabcdgfa", "befa", "dgc", "ade"}
	fmt.Println(findWords(board3, words3))
}

const debug = false

type TrieNode struct {
	IsLeaf    bool
	Letters   [26]*TrieNode
	WordCount int
}

func (node *TrieNode) RenderTrie() (result []string) {
	if node == nil {
		return
	}

	for idx, sub := range node.Letters {
		if sub == nil || sub.WordCount == 0 {
			continue
		}
		subLines := sub.RenderTrie()
		char := string(byte(idx + 'a'))
		for i := range subLines {
			if i == 0 {
				subLines[i] = char + subLines[i]
			} else {
				subLines[i] = " " + subLines[i]
			}
		}
		result = append(result, subLines...)

	}
	if node.IsLeaf {
		result = append(result, "X")
	}
	return
}

func (node *TrieNode) PrettyString(prefix string) string {
	result := []string{}
	for _, s := range node.RenderTrie() {
		result = append(result, prefix+s)
	}
	return strings.Join(result, "\n")
}

func (node *TrieNode) AddWord(w string) {
	for idx := range len(w) {
		node.WordCount++
		letter := w[idx] - 'a'
		if node.Letters[letter] == nil {
			node.Letters[letter] = new(TrieNode)
		}
		node = node.Letters[letter]
	}
	node.IsLeaf = true
	node.WordCount++
}

func (node *TrieNode) RemoveWord(w string) {
	for idx := range len(w) {
		node.WordCount--
		node = node.Letters[w[idx]-'a']
	}
	node.WordCount--
	node.IsLeaf = false
}

func (node *TrieNode) String() string {
	chars := ""
	if node.IsLeaf {
		chars += "X"
	}
	for i := range 26 {
		if node.Letters[i] != nil {
			chars += string(byte(i + 'a'))
		}
	}
	return fmt.Sprintf("[%-2d %-5t - %s]", node.WordCount, node.IsLeaf, chars)
}

func BuildTrie(words []string) *TrieNode {
	root := new(TrieNode)
	for _, w := range words {
		root.AddWord(w)
	}
	return root
}

const maxWordLength = 10

func findWords(board [][]byte, words []string) (result []string) {
	height, width := len(board), len(board[0])
	trieRoot := BuildTrie(words)

	boardIdxAt := func(c [2]int) byte {
		return board[c[1]][c[0]] - 'a'
	}
	boardAt := func(c [2]int) byte {
		return board[c[1]][c[0]]
	}
	deltas := [4][2]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}

	validNext := func(c [2]int) (result [][2]int) {
		for _, d := range deltas {
			x, y := c[0]+d[0], c[1]+d[1]
			if x >= 0 && x < width && y >= 0 && y < height {
				result = append(result, [2]int{x, y})
			}
		}
		return
	}

	var recur func(c [2]int, tNode *TrieNode, wordSoFar []byte, soFarLen int, visited map[[2]int]bool)
	recur = func(c [2]int, tNode *TrieNode, wordSoFar []byte, soFarLen int, visited map[[2]int]bool) {
		var firstPrefix, RetPrefix, prefix string
		if debug {
			firstPrefix = strings.Repeat(" ", soFarLen) + "> "
			RetPrefix = strings.Repeat(" ", soFarLen) + "< "
			prefix = strings.Repeat(" ", soFarLen) + " | "
			defer (func() { fmt.Println(RetPrefix) })()
		}
		if debug {
			fmt.Println(firstPrefix, c, "==", string(boardAt(c)), "["+string(wordSoFar[:soFarLen])+"]", visited)
			fmt.Println(prefix + tNode.String())
			fmt.Println(tNode.PrettyString(prefix))
		}
		if tNode == nil {
			return
		}

		if tNode.IsLeaf {
			w := string(wordSoFar[:soFarLen])
			if debug {
				fmt.Println(prefix, "Removing", w)
			}
			trieRoot.RemoveWord(w)
			result = append(result, w)
		}

		if tNode.WordCount == 0 {
			return
		}

		for _, newC := range validNext(c) {
			newTNode := tNode.Letters[boardIdxAt(newC)]
			if newTNode == nil {
				continue
			}
			if newTNode.WordCount == 0 {
				continue
			}
			if visited[newC] {
				continue
			}
			wordSoFar[soFarLen] = boardAt(newC)
			visited[newC] = true
			recur(newC, newTNode, wordSoFar, soFarLen+1, visited)
			visited[newC] = false
		}
	}

	// words are up to 10 letters, there is up to 3000 words in a list
outer:
	for y := range height {
		for x := range width {
			c := [2]int{x, y}
			subTrie := trieRoot.Letters[boardIdxAt(c)]
			if subTrie != nil {
				if debug {
					fmt.Println("vvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvv", c)
					fmt.Println(trieRoot.PrettyString("> "))
					fmt.Println("^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^")
				}
				wordSoFar := make([]byte, maxWordLength)
				wordSoFar[0] = boardAt(c)
				visited := map[[2]int]bool{}
				visited[c] = true
				recur(c, subTrie, wordSoFar, 1, visited)
			}
			if len(result) == len(words) {
				break outer
			}
		}
	}

	return result
}

// Example 1:

// Input: board = [["o","a","a","n"],["e","t","a","e"],["i","h","k","r"],["i","f","l","v"]], words = ["oath","pea","eat","rain"]
// Output: ["eat","oath"]

// Example 2:

// Input: board = [["a","b"],["c","d"]], words = ["abcb"]
// Output: []
