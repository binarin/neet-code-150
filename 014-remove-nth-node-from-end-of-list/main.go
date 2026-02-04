package main

import (
	"fmt"
	"math"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func sliceToList(s []int) *ListNode {
	dummy := &ListNode{}
	cur := dummy
	for _, v := range s {
		cur.Next = &ListNode{Val: v}
		cur = cur.Next
	}
	return dummy.Next
}

func listToSlice(head *ListNode) []int {
	result := []int{}
	for head != nil {
		result = append(result, head.Val)
		head = head.Next
	}
	return result
}

func main() {
	result := removeNthFromEnd(sliceToList([]int{1, 2, 3, 4, 5}), 2)
	fmt.Println(listToSlice(result))
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	type RecurRes struct {
		n        *ListNode
		from_end int
	}
	var recur func(*ListNode) RecurRes
	recur = func(cur *ListNode) RecurRes {
		if cur == nil {
			return RecurRes{n: nil, from_end: 0}
		}
		sub_res := recur(cur.Next)
		from_end := sub_res.from_end + 1
		if from_end == n {
			return RecurRes{n: sub_res.n, from_end: math.MinInt}
		} else {
			cur.Next = sub_res.n
			return RecurRes{n: cur, from_end: from_end}
		}
	}

	r := recur(head)
	return r.n
}
