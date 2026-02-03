package main

import (
	"fmt"
)

func main() {
	for _, inp := range []string{"babad", "cbbd"} {
		lp := longestPalindrome(inp)
		fmt.Printf("in: %s, out: %s\n", inp, lp)
	}
}


func reverse(s string) string {
	result := make([]byte, len(s))
	for i := 0; i < len(s); i++ {
		result[len(s) - 1 - i] = s[i]
	}
	return string(result)
}

func isPalindrome(s string) bool {
	/// len 0
	if len(s) == 1 {
		return true
	}
	// 0 1 2 3 4 5   (len: 6) -> [0:3] == [3:]
	if len(s) % 2 == 0 {
		return s[0:len(s) / 2] == reverse(s[len(s) / 2:])
	}
	// 0 1 2 3 4 5 7 (len: 7, len/2: 3) -> [0:3] == [4:]
	// 0 1 2 (len: 3)
	return s[0:len(s) / 2] == reverse(s[len(s)/2 + 1:])
}

func longestPalindrome(s string) string {
	result := ""
	if len(s) == 0 {
		return ""
	}
	if len(s) == 1 {
		return s
	}
	// let’s start with the brute-force approach
	for i := 0; i < len(s) - 1; i++ {
		for j := i + 1; j <= len(s); j++ {
			if isPalindrome(s[i:j]) && len(result) < j - i {
				result = s[i:j]
			}
		}
	}
	return result
}
