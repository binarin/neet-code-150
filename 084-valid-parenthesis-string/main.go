// https://leetcode.com/problems/valid-parenthesis-string/
// Difficulty: Medium

package main

import (
	"fmt"
)

func checkValidStringSimple(s string) bool {
	// fmt.Println(s)
	openCount := 0
	for _, c := range s {
		if c == '(' {
			openCount++
		} else if openCount > 0 {
			openCount--
		} else {
			return false
		}
	}
	return openCount == 0
}
func checkValidString(s string) bool {
	cache := map[[2]int]bool{}
	var recur func(string, int) bool
	recur = func(rest string, openCount int) (result bool) {
		if len(rest) == 0 {
			return openCount == 0
		}
		cacheKey := [2]int{len(rest), openCount}
		if cached, ok := cache[cacheKey]; ok {
			return cached
		}
		defer (func() {
			cache[cacheKey] = result
		})()
		for idx := range len(rest) {
			switch rest[idx] {
			case '(':
				openCount++
				continue
			case ')':
				if openCount == 0 {
					return false
				}
				openCount--
			default:
				asEmpty := recur(rest[idx+1:], openCount)
				if asEmpty {
					return true
				}

				if openCount > 0 {
					asClose := recur(rest[idx+1:], openCount-1)
					if asClose {
						return true
					}
				}

				asOpen := recur(rest[idx+1:], openCount+1)
				return asOpen
			}
		}
		return openCount == 0
	}
	return recur(s, 0)
}

func main() {
	// Example 1: s = "()" -> true
	// fmt.Println(checkValidString("(((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((((("))
	// fmt.Println(checkValidString("(((((()*)(*)*))())())(()())())))((**)))))(()())()"))
	fmt.Println(checkValidString("**************************************************))))))))))))))))))))))))))))))))))))))))))))))))))"))
}
