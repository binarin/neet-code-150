// https://leetcode.com/problems/implement-trie-prefix-tree/description/?envType=problem-list-v2&envId=plakya4j
// Difficulty: Medium

package main

import "fmt"

type Trie struct {
}

func Constructor() Trie {
	return Trie{}
}

func (this *Trie) Insert(word string) {
}

func (this *Trie) Search(word string) bool {
	return false
}

func (this *Trie) StartsWith(prefix string) bool {
	return false
}

func main() {
	// Example 1:
	// Input: ["Trie", "insert", "search", "search", "startsWith", "insert", "search"]
	//        [[], ["apple"], ["apple"], ["app"], ["app"], ["app"], ["app"]]
	// Output: [null, null, true, false, true, null, true]

	trie := Constructor()
	trie.Insert("apple")
	fmt.Println(trie.Search("apple"))    // Expected: true
	fmt.Println(trie.Search("app"))      // Expected: false
	fmt.Println(trie.StartsWith("app"))  // Expected: true
	trie.Insert("app")
	fmt.Println(trie.Search("app"))      // Expected: true
}
