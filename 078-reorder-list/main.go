// https://leetcode.com/problems/reorder-list/
// Difficulty: Medium

package main

import "fmt"

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func listLen(n *ListNode) int {
	length := 0
	for n != nil {
		length++
		n = n.Next
	}
	return length
}

func nReverse(n **ListNode) {
	var prev, cur *ListNode = nil, *n
	if cur == nil {
		return
	}
	if cur.Next == nil {
		return
	}
	for cur != nil {
		// fmt.Println("b", prev, cur)
		cur.Next, prev, cur = prev, cur, cur.Next
		// fmt.Println("a", prev, cur)
	}
	*n = prev
}

func drop(n *ListNode, count int) *ListNode {
	for n != nil && count > 0 {
		n = n.Next
		count--
	}
	return n
}

func reorderList(head *ListNode) {
	l := listLen(head)
	if l == 0 || l == 1 {
		return
	}
	snd := drop(head, l/2)
	// fmt.Println(l, listToSlice(snd))
	nReverse(&snd)
	// fmt.Println(listToSlice(snd))
	for {
		// fmt.Println("b", head, snd)
		if snd == nil || head.Next == snd || head == snd {
			break
		}
		// (1)->(2)->(3->nil)<-(4)<-(5)
		// ^1   ^2  ^3^3  ^2  ^1
		//
		// (1)->(2)->(3->nil)<-(4)
		//  ^1  2^    2^       1^
		//
		head.Next, snd.Next, head, snd = snd, head.Next, head.Next, snd.Next
	}
}

// Helper function to create a linked list from a slice
func createList(vals []int) *ListNode {
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

// Helper function to convert a linked list to a slice
func listToSlice(head *ListNode) []int {
	var result []int
	for head != nil {
		result = append(result, head.Val)
		head = head.Next
	}
	return result
}

func main() {
	// Example 1: head = [1,2,3,4] -> [1,4,2,3]
	head := createList([]int{1, 2, 3, 4})
	reorderList(head)
	fmt.Println(listToSlice(head))
}
