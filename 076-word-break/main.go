// https://leetcode.com/problems/word-break/
// Difficulty: Medium

package main

import "fmt"

type Trie struct {
	IsLeaf  bool
	Letters [26]*Trie
}

func charIdx(c byte) int {
	return int(c - 'a')
}

func (t *Trie) Add(s string) {
	for idx := range len(s) {
		cid := charIdx(s[idx])
		if t.Letters[cid] == nil {
			t.Letters[cid] = &Trie{}
		}
		t = t.Letters[cid]
	}
	t.IsLeaf = true
}

func wordBreak(s string, wordDict []string) bool {
	breakableAfter := map[int]bool{}
	trieRoot := &Trie{}
	for _, word := range wordDict {
		trieRoot.Add(word)
	}
	var recur func(int) bool
	recur = func(start int) (result bool) {
		if cache, ok := breakableAfter[start]; ok {
			return cache
		}
		defer (func() { breakableAfter[start] = result })()

		curS, curT := start, trieRoot
		for {
			if curS == len(s) {
				return curT.IsLeaf
			}
			if curT.IsLeaf {
				result := recur(curS)
				if result {
					return true
				}
			}
			next := curT.Letters[charIdx(s[curS])]
			if next == nil {
				return false
			}
			curS++
			curT = next
		}
	}
	return recur(0)
}

func main() {
	// Example 1: s = "leetcode", wordDict = ["leet","code"]
	// Expected output: true
	result := wordBreak("leetcode", []string{"leet", "code"})
	fmt.Println(result)
}
