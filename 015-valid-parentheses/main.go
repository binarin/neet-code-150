package main

import (
	"fmt"
)

func main() {
	fmt.Println(isValid("()"))
}

func isValid(s string) bool {
	reverse := map[byte]byte{'}': '{', ']': '[', ')': '('}
	stack := []byte{}
	for i := 0; i < len(s); i++ {
		switch s[i] {
		case '(':
			fallthrough
		case '[':
			fallthrough
		case '{':
			stack = append(stack, s[i])
		case ')':
			fallthrough
		case ']':
			fallthrough
		case '}':
			if len(stack) == 0 || stack[len(stack)-1] != reverse[s[i]] {
				return false
			}
			stack = stack[0 : len(stack)-1]
		}
	}

	return len(stack) == 0
}
