package main

import "testing"

func TestExample(t *testing.T) {
	// Input: ["WordDictionary","addWord","addWord","addWord","search","search","search","search"]
	//        [[],["bad"],["dad"],["mad"],["pad"],["bad"],[".ad"],["b.."]]
	// Output: [null,null,null,null,false,true,true,true]

	wordDictionary := Constructor()
	wordDictionary.AddWord("bad")
	wordDictionary.AddWord("dad")
	wordDictionary.AddWord("mad")

	// search("pad") -> false
	if got := wordDictionary.Search("pad"); got != false {
		t.Errorf("Search(\"pad\") = %v, want false", got)
	}

	// search("bad") -> true
	if got := wordDictionary.Search("bad"); got != true {
		t.Errorf("Search(\"bad\") = %v, want true", got)
	}

	// search(".ad") -> true
	if got := wordDictionary.Search(".ad"); got != true {
		t.Errorf("Search(\".ad\") = %v, want true", got)
	}

	// search("b..") -> true
	if got := wordDictionary.Search("b.."); got != true {
		t.Errorf("Search(\"b..\") = %v, want true", got)
	}
}
