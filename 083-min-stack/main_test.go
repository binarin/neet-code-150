package main

import "testing"

func TestMinStack_Example1(t *testing.T) {
	// Input: ["MinStack","push","push","push","getMin","pop","top","getMin"]
	//        [[],[-2],[0],[-3],[],[],[],[]]
	// Output: [null,null,null,null,-3,null,0,-2]

	minStack := Constructor()
	minStack.Push(-2)
	minStack.Push(0)
	minStack.Push(-3)

	if got := minStack.GetMin(); got != -3 {
		t.Errorf("GetMin() = %v, want -3", got)
	}

	minStack.Pop()

	if got := minStack.Top(); got != 0 {
		t.Errorf("Top() = %v, want 0", got)
	}

	if got := minStack.GetMin(); got != -2 {
		t.Errorf("GetMin() = %v, want -2", got)
	}
}
