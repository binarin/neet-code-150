// https://leetcode.com/problems/reverse-linked-list/description/?envType=problem-list-v2&envId=plakya4j
// Difficulty: Easy

package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func buildList(vals []int) *ListNode {
	if len(vals) == 0 {
		return nil
	}
	head := &ListNode{Val: vals[0]}
	current := head
	for i := 1; i < len(vals); i++ {
		current.Next = &ListNode{Val: vals[i]}
		current = current.Next
	}
	return head
}

func listToSlice(head *ListNode) []int {
	var result []int
	for head != nil {
		result = append(result, head.Val)
		head = head.Next
	}
	return result
}

func main() {
	// Example 1: head = [1,2,3,4,5] -> [5,4,3,2,1]
	head := buildList([]int{1, 2, 3, 4, 5})
	result := reverseList(head)
	fmt.Println(listToSlice(result))
}

func reverseList(head *ListNode) *ListNode {
	return nil
}
