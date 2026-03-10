// https://leetcode.com/problems/min-stack/
// Difficulty: Medium

package main

import (
	"fmt"
	"math"
)

type MinStack struct {
	Stack [][2]int
	Min   int
}

func Constructor() MinStack {
	return MinStack{Min: math.MaxInt}
}

func (this *MinStack) Push(val int) {
	this.Stack = append(this.Stack, [2]int{val, this.Min})
	this.Min = min(this.Min, val)
}

func (this *MinStack) Pop() {
	this.Min = this.Stack[len(this.Stack)-1][1]
	this.Stack = this.Stack[0 : len(this.Stack)-1]
}

func (this *MinStack) Top() int {
	return this.Stack[len(this.Stack)-1][0]
}

func (this *MinStack) GetMin() int {
	return this.Min
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
