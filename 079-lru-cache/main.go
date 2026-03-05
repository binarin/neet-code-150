// https://leetcode.com/problems/lru-cache/
// Difficulty: Medium

package main

import "fmt"

type LRUCache struct {
}

func Constructor(capacity int) LRUCache {
	return LRUCache{}
}

func (this *LRUCache) Get(key int) int {
	return 0
}

func (this *LRUCache) Put(key int, value int) {
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */

func main() {
	// Example 1:
	// Input: ["LRUCache", "put", "put", "get", "put", "get", "put", "get", "get", "get"]
	//        [[2], [1, 1], [2, 2], [1], [3, 3], [2], [4, 4], [1], [3], [4]]
	// Output: [null, null, null, 1, null, -1, null, -1, 3, 4]

	lRUCache := Constructor(2)
	lRUCache.Put(1, 1)       // cache is {1=1}
	lRUCache.Put(2, 2)       // cache is {1=1, 2=2}
	fmt.Println(lRUCache.Get(1)) // return 1
	lRUCache.Put(3, 3)       // LRU key was 2, evicts key 2, cache is {1=1, 3=3}
	fmt.Println(lRUCache.Get(2)) // returns -1 (not found)
	lRUCache.Put(4, 4)       // LRU key was 1, evicts key 1, cache is {4=4, 3=3}
	fmt.Println(lRUCache.Get(1)) // return -1 (not found)
	fmt.Println(lRUCache.Get(3)) // return 3
	fmt.Println(lRUCache.Get(4)) // return 4
}
