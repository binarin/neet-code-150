// https://leetcode.com/problems/last-stone-weight/
// Difficulty: Easy
package main

import (
	"cmp"
	"container/heap"
	"fmt"
	"sort"
)

func main() {
	fmt.Println(lastStoneWeight([]int{2, 7, 4, 1, 8, 1}))
}

type MaxHeap[T cmp.Ordered] []T

func (h *MaxHeap[T]) Add(item T) {
	*h = append(*h, item)
	idx := len(*h) - 1
	// 1 11 12 111 112 121 122 1111 1112
	// 0  1  2  3    4   5   6   7    8
	for idx > 0 {
		parentIdx := (idx - 1) / 2 // 1,2 => 0, 3,4 => 1, 5,6 => 2, 7,8 = 3
		if (*h)[parentIdx] >= (*h)[idx] {
			return
		}
		(*h)[parentIdx], (*h)[idx] = (*h)[idx], (*h)[parentIdx]
		idx = parentIdx
	}
}

func (h *MaxHeap[T]) Next() T {
	result := (*h)[0]

	(*h)[0] = (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]

	idx := 0
	l := len(*h)
	for idx < l {
		leftIdx := (idx+1)*2 - 1
		rightIdx := (idx + 1) * 2
		largestIdx := idx
		if leftIdx < l && (*h)[leftIdx] > (*h)[largestIdx] {
			largestIdx = leftIdx
		}
		if rightIdx < l && (*h)[rightIdx] > (*h)[largestIdx] {
			largestIdx = rightIdx
		}
		if largestIdx == idx {
			break
		}
		(*h)[idx], (*h)[largestIdx] = (*h)[largestIdx], (*h)[idx]
		idx = largestIdx
	}

	return result
}

func (h *MaxHeap[T]) Peek() (T, error) {
	if len(*h) == 0 {
		var nothing T
		return nothing, fmt.Errorf("Can't Peek() into empty heap")
	}
	return (*h)[0], nil
}

func (h *MaxHeap[T]) Len() int {
	return len(*h)
}

type MaxHeapContainer []int

func (h MaxHeapContainer) Len() int           { return len(h) }
func (h MaxHeapContainer) Less(i, j int) bool { return h[j] < h[i] }
func (h MaxHeapContainer) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MaxHeapContainer) Push(x any) {
	*h = append(*h, x.(int))
}

func (h *MaxHeapContainer) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func lastStoneWeight(stones []int) int {
	if len(stones) == 0 {
		return 0
	}
	h := new(MaxHeapContainer)
	for _, s := range stones {
		*h = append(*h, s)
	}
	heap.Init(h)

	for h.Len() > 1 {
		fmt.Println(h)
		bigger := heap.Pop(h).(int)
		smaller := heap.Pop(h).(int)
		if bigger == smaller {
			continue
		}
		heap.Push(h, bigger-smaller)
	}

	if h.Len() == 0 {
		return 0
	} else {
		return (*h)[0]
	}
}

func lastStoneWeightOwnHeap(stones []int) int {
	if len(stones) == 0 {
		return 0
	}
	heap := new(MaxHeap[int])
	for _, s := range stones {
		heap.Add(s)
	}
	for heap.Len() > 1 {
		fmt.Println(heap)
		bigger := heap.Next()
		smaller := heap.Next()
		if bigger == smaller {
			continue
		}
		heap.Add(bigger - smaller)
	}

	if heap.Len() == 0 {
		return 0
	} else {
		result, err := heap.Peek()
		if err != nil {
			panic("shouldn't be here")
		}
		return result
	}
}

func lastStoneWeightBf(stones []int) int {
	if len(stones) == 0 {
		return 0
	}
	if len(stones) == 1 {
		return stones[0]
	}
	sort.Slice(stones, func(i, j int) bool {
		return stones[i] > stones[j]
	})
	fmt.Println(stones)
	x, y := stones[0], stones[1]
	if x == y {
		return lastStoneWeightBf(stones[2:])
	}
	rem := max(x-y, y-x)
	stones = stones[1:]
	stones[0] = rem
	return lastStoneWeightBf(stones)
}
