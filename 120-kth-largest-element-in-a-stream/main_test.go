package main

import "testing"

func TestExample1(t *testing.T) {
	// Input: ["KthLargest", "add", "add", "add", "add", "add"]
	//        [[3, [4, 5, 8, 2]], [3], [5], [10], [9], [4]]
	// Output: [null, 4, 5, 5, 8, 8]

	kthLargest := Constructor(3, []int{4, 5, 8, 2})

	if got := kthLargest.Add(3); got != 4 {
		t.Errorf("Add(3) = %d, want 4", got)
	}
	if got := kthLargest.Add(5); got != 5 {
		t.Errorf("Add(5) = %d, want 5", got)
	}
	if got := kthLargest.Add(10); got != 5 {
		t.Errorf("Add(10) = %d, want 5", got)
	}
	if got := kthLargest.Add(9); got != 8 {
		t.Errorf("Add(9) = %d, want 8", got)
	}
	if got := kthLargest.Add(4); got != 8 {
		t.Errorf("Add(4) = %d, want 8", got)
	}
}

func TestExample2(t *testing.T) {
	// Input: ["KthLargest", "add", "add", "add", "add"]
	//        [[4, [7, 7, 7, 7, 8, 3]], [2], [10], [9], [9]]
	// Output: [null, 7, 7, 7, 8]

	kthLargest := Constructor(4, []int{7, 7, 7, 7, 8, 3})

	if got := kthLargest.Add(2); got != 7 {
		t.Errorf("Add(2) = %d, want 7", got)
	}
	if got := kthLargest.Add(10); got != 7 {
		t.Errorf("Add(10) = %d, want 7", got)
	}
	if got := kthLargest.Add(9); got != 7 {
		t.Errorf("Add(9) = %d, want 7", got)
	}
	if got := kthLargest.Add(9); got != 8 {
		t.Errorf("Add(9) = %d, want 8", got)
	}
}
