package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func mkList(vals []int) *ListNode {
	dummy := &ListNode{}
	cur := dummy
	for _, v := range vals {
		cur.Next = &ListNode{Val: v}
		cur = cur.Next
	}
	return dummy.Next
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
	lists := []*ListNode{
		mkList([]int{1, 4, 5}),
		mkList([]int{1, 3, 4}),
		mkList([]int{2, 6}),
	}
	fmt.Println(listToSlice(mergeKLists(lists)))
}

func mergeLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	if l2.Val < l1.Val {
		l1, l2 = l2, l1
	}
	l1.Next = mergeLists(l1.Next, l2)
	return l1
}

func mergeListsIter(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}

	var head, cur *ListNode
	if l2.Val < l1.Val {
		l1, l2 = l2, l1
	}
	head = l1
	cur = head
	l1 = l1.Next

	for {
		if l1 == nil {
			cur.Next = l2
			break
		}
		if l2 == nil {
			cur.Next = l1
			break
		}

		if l2.Val < l1.Val {
			l1, l2 = l2, l1
		}
		cur.Next = l1
		cur = l1
		l1 = l1.Next
	}
	return head

}

func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	if len(lists) == 1 {
		return lists[0]
	}
	// ----------- -----------
	left := mergeKLists(lists[0 : len(lists)/2])
	right := mergeKLists(lists[len(lists)/2:])
	return mergeListsIter(left, right)
}
