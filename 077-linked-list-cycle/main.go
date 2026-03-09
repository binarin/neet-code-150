// https://leetcode.com/problems/linked-list-cycle/
// Difficulty: Easy

package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}
	rabbit, hare := head, head
	for {
		for range 2 {
			rabbit = rabbit.Next
			if rabbit == nil {
				return false
			}
			if rabbit == hare {
				return true
			}
		}
		hare = hare.Next
	}
	return false
}

func main() {
	// Example 1: head = [3,2,0,-4], pos = 1 -> true
	// Create nodes
	node0 := &ListNode{Val: 3}
	node1 := &ListNode{Val: 2}
	node2 := &ListNode{Val: 0}
	node3 := &ListNode{Val: -4}
	// Link nodes
	node0.Next = node1
	node1.Next = node2
	node2.Next = node3
	// Create cycle: tail connects to node at index 1
	node3.Next = node1

	result := hasCycle(node0)
	fmt.Println(result)
}
