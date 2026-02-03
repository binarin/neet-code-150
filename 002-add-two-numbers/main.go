package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func ln(val int, next *ListNode) *ListNode {
	return &ListNode{Val: val, Next: next}
}

func main() {
	n342 := ln(2, ln(4, ln(3, nil)))
	n465 := ln(5, ln(6, ln(4, nil)))
	print_list(addTwoNumbers(n342, n465))
}

func print_list(lst *ListNode) {
	var items = []int{}
	for lst != nil {
		items = append(items, lst.Val)
		lst = lst.Next
	}
	fmt.Println(items)
}

func nreverse(lst *ListNode) *ListNode {
	if lst == nil {
		return nil
	}
	if lst.Next == nil {
		return lst
	}

	prev := lst
	cur := lst.Next
	prev.Next = nil

	for cur != nil {
		save_next := cur.Next
		cur.Next = prev
		prev = cur
		cur = save_next
	}

	return prev
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var prev *ListNode = nil
	var prev_carry int = 0
	for l1 != nil || l2 != nil {
		var l1_v, l2_v int
		if l1 != nil {
			l1_v = l1.Val
		}
		if l2 != nil {
			l2_v = l2.Val
		}
		var val = l1_v + l2_v + prev_carry
		prev = &ListNode{Val: val % 10, Next: prev}
		prev_carry = val / 10
		if l1 != nil {
			l1 = l1.Next
		}
		if l2 != nil {
			l2 = l2.Next
		}
	}
	if prev_carry > 0 {
		prev = &ListNode{Val: prev_carry, Next: prev}
	}
	return nreverse(prev)
}
