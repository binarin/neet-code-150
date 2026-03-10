// https://leetcode.com/problems/lru-cache/
// Difficulty: Medium

package main

import (
	"fmt"
	"maps"
	"slices"
	"strings"
)

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
	fmt.Println("constr 2", lRUCache)
	lRUCache.Put(1, 1) // cache is {1=1}
	fmt.Println("put 1", lRUCache)
	lRUCache.Put(2, 2) // cache is {1=1, 2=2}
	fmt.Println("put 2", lRUCache)
	fmt.Println("get 1(exp 1)", lRUCache.Get(1), lRUCache)
	lRUCache.Put(3, 3)
	fmt.Println("put 3(2 evicted, 1,3 left)", lRUCache)
	lRUCache.Put(4, 4)
	fmt.Println("put 4(1 evicted, 3,4 left)", lRUCache)
	fmt.Println("get 1(exp -1)", lRUCache.Get(1), lRUCache)
	fmt.Println("get 1(exp 3)", lRUCache.Get(3), lRUCache)
	fmt.Println("get 1(exp 4)", lRUCache.Get(4), lRUCache)
}

type LRUCache struct {
	Capacity int
	Values   map[int]int
	RingMap  map[int]*DLNode
	Ring     *DLNode
}

type DLNode struct {
	Key        int
	Prev, Next *DLNode
}

func Constructor(capacity int) LRUCache {
	cache := LRUCache{Capacity: capacity}
	cache.Ring = &DLNode{}
	cache.Ring.Prev, cache.Ring.Next = cache.Ring, cache.Ring
	cache.Values = map[int]int{}
	cache.RingMap = map[int]*DLNode{}
	return cache
}

func (c *LRUCache) EnsureUnlinked(key int) *DLNode {
	if n, ok := c.RingMap[key]; ok {
		n.Prev.Next, n.Next.Prev = n.Next, n.Prev
		n.Next, n.Prev = nil, nil
		return n
	} else {
		n := &DLNode{Key: key}
		c.RingMap[key] = n
		return n
	}
}

func (c *LRUCache) TrackUsage(key int) {
	n := c.EnsureUnlinked(key)
	c.Ring.Prev.Next, c.Ring.Prev, n.Prev, n.Next = n, n, c.Ring.Prev, c.Ring
}

func (c LRUCache) String() string {
	b := strings.Builder{}
	fmt.Fprintf(&b, "\n[\tusage: %d/%d\n\tcache: ", len(c.Values), c.Capacity)
	for idx, key := range slices.Sorted(maps.Keys(c.Values)) {
		if idx != 0 {
			b.WriteString(", ")
		}
		fmt.Fprintf(&b, "%d=%d", key, c.Values[key])
	}
	fmt.Fprintf(&b, "\n\tMRU: %d, LRU: %d", c.Ring.Prev.Key, c.Ring.Next.Key)
	b.WriteString("\n]")

	return b.String()
}

func (c *LRUCache) Get(key int) int {
	val, ok := c.Values[key]
	if !ok {
		return -1
	}
	c.TrackUsage(key)
	return val
}

func (c *LRUCache) Put(key int, value int) {
	if _, ok := c.Values[key]; !ok && len(c.Values) == c.Capacity {
		n := c.EnsureUnlinked(c.Ring.Next.Key)
		delete(c.Values, n.Key)
		delete(c.RingMap, n.Key)
	}
	c.TrackUsage(key)
	c.Values[key] = value
}
