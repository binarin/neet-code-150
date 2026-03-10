// https://leetcode.com/problems/evaluate-reverse-polish-notation/
// Difficulty: Medium

package main

import (
	"fmt"
	"strconv"
)

func evalRPN(tokens []string) int {
	stack := []int{}
	for _, tok := range tokens {
		if num, err := strconv.Atoi(tok); err == nil {
			stack = append(stack, num)
			continue
		}

		var result int
		op1, op2 := stack[len(stack)-2], stack[len(stack)-1]

		switch tok {
		case "+":
			result = op1 + op2
		case "-":
			result = op1 - op2
		case "*":
			result = op1 * op2
		case "/":
			result = op1 / op2
		}
		stack = stack[:len(stack)-2]
		stack = append(stack, result)
	}
	return stack[len(stack)-1]
}

func main() {
	// Example 1: tokens = ["2","1","+","3","*"], Output: 9
	tokens := []string{"2", "1", "+", "3", "*"}
	fmt.Println(evalRPN(tokens))
}
