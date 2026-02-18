package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(largestRectangleArea([]int{2, 1, 5, 6, 2, 3}))
}

type Stack [][2]int

func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

func (s *Stack) Push(val, pos int) {
	*s = append(*s, [2]int{val, pos})
}

func (s *Stack) TopVal() int {
	return (*s)[len(*s)-1][0]
}

func (s *Stack) TopPos() int {
	return (*s)[len(*s)-1][1]
}

func (s *Stack) Pop() {
	*s = (*s)[:len(*s)-1]
}

func (s *Stack) String() string {
	parts := make([]string, len(*s))
	for i, v := range *s {
		parts[i] = fmt.Sprintf("%d:%d", v[0], v[1])
	}
	return strings.Join(parts, ", ")
}

func use(...any) {
}

func largestRectangleArea(heights []int) int {
	stack := Stack{}

	maxArea := 0
	handleRect := func(height, left, right int) {
		area := height * (right - left)
		rectStr := fmt.Sprintf("%d:%d→%d(%d)", height, left, right, area)
		use(rectStr)
		if area > maxArea {
			maxArea = area
			// fmt.Println("Improved", rectStr)
		} else {
			// fmt.Println("Discarded", rectStr)
		}
	}

outer:
	for i := range heights {
		// fmt.Println("At", i, ":", stack.String())
		cur := heights[i]
		if stack.IsEmpty() {
			stack.Push(cur, i)
			continue
		}
		if stack.TopVal() < cur {
			stack.Push(cur, i)
			continue
		} else if stack.TopVal() == cur {
			continue
		} else {
			bestFallback := i
			for !stack.IsEmpty() {
				if stack.TopVal() > cur {
					handleRect(stack.TopVal(), stack.TopPos(), i)
					bestFallback = stack.TopPos()
					stack.Pop()
				} else if stack.TopVal() == cur {
					continue outer
				} else {
					break
				}
			}
			stack.Push(cur, bestFallback)
		}

	}

	for !stack.IsEmpty() {
		handleRect(stack.TopVal(), stack.TopPos(), len(heights))
		stack.Pop()
	}
	return maxArea
}
