package main

import "testing"

func TestExample1(t *testing.T) {
	// Input: ["Trie", "insert", "search", "search", "startsWith", "insert", "search"]
	//        [[], ["apple"], ["apple"], ["app"], ["app"], ["app"], ["app"]]
	// Output: [null, null, true, false, true, null, true]

	trie := Constructor()
	trie.Insert("apple")

	if got := trie.Search("apple"); got != true {
		t.Errorf("Search(\"apple\") = %v, want true", got)
	}

	if got := trie.Search("app"); got != false {
		t.Errorf("Search(\"app\") = %v, want false", got)
	}

	if got := trie.StartsWith("app"); got != true {
		t.Errorf("StartsWith(\"app\") = %v, want true", got)
	}

	trie.Insert("app")

	if got := trie.Search("app"); got != true {
		t.Errorf("Search(\"app\") after insert = %v, want true", got)
	}
}
