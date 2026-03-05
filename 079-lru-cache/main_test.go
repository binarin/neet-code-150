package main

import (
	"testing"
)

func TestLRUCache_Example1(t *testing.T) {
	// Input: ["LRUCache", "put", "put", "get", "put", "get", "put", "get", "get", "get"]
	//        [[2], [1, 1], [2, 2], [1], [3, 3], [2], [4, 4], [1], [3], [4]]
	// Output: [null, null, null, 1, null, -1, null, -1, 3, 4]

	lRUCache := Constructor(2)
	lRUCache.Put(1, 1) // cache is {1=1}
	lRUCache.Put(2, 2) // cache is {1=1, 2=2}

	if got := lRUCache.Get(1); got != 1 {
		t.Errorf("Get(1) = %d, want 1", got)
	}

	lRUCache.Put(3, 3) // LRU key was 2, evicts key 2, cache is {1=1, 3=3}

	if got := lRUCache.Get(2); got != -1 {
		t.Errorf("Get(2) = %d, want -1 (not found)", got)
	}

	lRUCache.Put(4, 4) // LRU key was 1, evicts key 1, cache is {4=4, 3=3}

	if got := lRUCache.Get(1); got != -1 {
		t.Errorf("Get(1) = %d, want -1 (not found)", got)
	}

	if got := lRUCache.Get(3); got != 3 {
		t.Errorf("Get(3) = %d, want 3", got)
	}

	if got := lRUCache.Get(4); got != 4 {
		t.Errorf("Get(4) = %d, want 4", got)
	}
}
