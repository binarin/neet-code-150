package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func sliceToList(vals []int) *ListNode {
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
	fmt.Println(listToSlice(reverseKGroup(sliceToList([]int{1, 2, 3, 4, 5}), 2)))
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	if head == nil {
		return nil
	}

	cur := head.Next

	head.Next = nil
	var result_head *ListNode = &ListNode{Next: head}

	chunk_head := result_head
	chunk_tail := result_head.Next
	chunk_length := 1

	for {
		if cur == nil {
			break
		}
		if chunk_length == k {
			chunk_head = chunk_tail

			next := cur.Next
			cur.Next = nil
			chunk_head.Next = cur
			chunk_tail = cur
			chunk_length = 1

			cur = next
			continue
		}

		next := cur.Next
		cur.Next = chunk_head.Next
		chunk_head.Next = cur
		cur = next
		chunk_length++
	}

	var reverse func(*ListNode) *ListNode
	reverse = func(head *ListNode) *ListNode {
		if head == nil {
			return nil
		}
		if head.Next == nil {
			return head
		}
		prev := head
		cur := prev.Next
		prev.Next = nil
		for cur != nil {
			next := cur.Next
			cur.Next = prev
			prev = cur
			cur = next
		}
		return prev
	}

	if chunk_length != k {
		chunk_head.Next = reverse(chunk_head.Next)
	}

	return result_head.Next
}
