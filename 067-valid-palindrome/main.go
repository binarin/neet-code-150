// https://leetcode.com/problems/valid-palindrome/
// Difficulty: Easy

package main

import (
	"fmt"
	"unicode"
)

func isPalindrome(s string) bool {
	clean := make([]rune, 0, len(s))
	for _, rune := range s {
		if unicode.IsNumber(rune) || unicode.IsLower(rune) {
			clean = append(clean, rune)
		} else if unicode.IsUpper(rune) {
			clean = append(clean, unicode.ToLower(rune))
		}
	}
	left, right := 0, len(clean)-1
	for left <= right {
		if clean[left] != clean[right] {
			return false
		}
		left++
		right--
	}
	return true
}

func main() {
	result := isPalindrome("A man, a plan, a canal: Panama")
	fmt.Println(result)
}
