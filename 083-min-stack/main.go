package main

import "fmt"

type MinStack struct {
}

func Constructor() MinStack {
	return MinStack{}
}

func (this *MinStack) Push(val int) {
}

func (this *MinStack) Pop() {
}

func (this *MinStack) Top() int {
	return 0
}

func (this *MinStack) GetMin() int {
	return 0
}

func main() {
	// Example 1:
	// Input: ["MinStack","push","push","push","getMin","pop","top","getMin"]
	//        [[],[-2],[0],[-3],[],[],[],[]]
	// Output: [null,null,null,null,-3,null,0,-2]

	minStack := Constructor()
	minStack.Push(-2)
	minStack.Push(0)
	minStack.Push(-3)
	fmt.Println(minStack.GetMin()) // expected: -3
	minStack.Pop()
	fmt.Println(minStack.Top())    // expected: 0
	fmt.Println(minStack.GetMin()) // expected: -2
}
