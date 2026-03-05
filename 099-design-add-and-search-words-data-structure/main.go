// https://leetcode.com/problems/design-add-and-search-words-data-structure/description/?envType=problem-list-v2&envId=plakya4j
// Difficulty: Medium

package main

import "fmt"

type WordDictionary struct {
}

func Constructor() WordDictionary {
	return WordDictionary{}
}

func (this *WordDictionary) AddWord(word string) {
}

func (this *WordDictionary) Search(word string) bool {
	return false
}

func main() {
	// Example:
	// Input: ["WordDictionary","addWord","addWord","addWord","search","search","search","search"]
	//        [[],["bad"],["dad"],["mad"],["pad"],["bad"],[".ad"],["b.."]]
	// Output: [null,null,null,null,false,true,true,true]

	wordDictionary := Constructor()
	wordDictionary.AddWord("bad")
	wordDictionary.AddWord("dad")
	wordDictionary.AddWord("mad")
	fmt.Println(wordDictionary.Search("pad")) // expected: false
	fmt.Println(wordDictionary.Search("bad")) // expected: true
	fmt.Println(wordDictionary.Search(".ad")) // expected: true
	fmt.Println(wordDictionary.Search("b..")) // expected: true
}
