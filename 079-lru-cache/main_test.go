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

func TestLRUCache_Example2(t *testing.T) {
	// Input: ["LRUCache","get","get","put","get","put","put","put","put","get","put"]
	//        [[1],[6],[8],[12,1],[2],[15,11],[5,2],[1,15],[4,2],[5],[15,15]]
	// Output: [null,-1,-1,null,-1,null,null,null,null,-1,null]

	lRUCache := Constructor(1)

	if got := lRUCache.Get(6); got != -1 {
		t.Errorf("Get(6) = %d, want -1", got)
	}

	if got := lRUCache.Get(8); got != -1 {
		t.Errorf("Get(8) = %d, want -1", got)
	}

	lRUCache.Put(12, 1) // cache is {12=1}

	if got := lRUCache.Get(2); got != -1 {
		t.Errorf("Get(2) = %d, want -1", got)
	}

	lRUCache.Put(15, 11) // evicts 12, cache is {15=11}
	lRUCache.Put(5, 2)   // evicts 15, cache is {5=2}
	lRUCache.Put(1, 15)  // evicts 5, cache is {1=15}
	lRUCache.Put(4, 2)   // evicts 1, cache is {4=2}

	if got := lRUCache.Get(5); got != -1 {
		t.Errorf("Get(5) = %d, want -1", got)
	}

	lRUCache.Put(15, 15) // evicts 4, cache is {15=15}
}
