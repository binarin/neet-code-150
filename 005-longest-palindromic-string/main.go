package main

import (
	"fmt"
)

func main() {
	for _, inp := range []string{"ccc", "bb", "babad", "cbbd"} {
		lp := longestPalindrome(inp)
		fmt.Printf("in: %s, out: %s\n", inp, lp)
	}
}

func reverse(s string) string {
	result := make([]byte, len(s))
	for i := 0; i < len(s); i++ {
		result[len(s)-1-i] = s[i]
	}
	return string(result)
}

func isPalindrome(s string) bool {
	/// len 0
	if len(s) == 1 {
		return true
	}
	// 0 1 2 3 4 5   (len: 6) -> [0:3] == [3:]
	if len(s)%2 == 0 {
		return s[0:len(s)/2] == reverse(s[len(s)/2:])
	}
	// 0 1 2 3 4 5 7 (len: 7, len/2: 3) -> [0:3] == [4:]
	// 0 1 2 (len: 3)
	return s[0:len(s)/2] == reverse(s[len(s)/2+1:])
}

func longestPalindrome(s string) string {
	result := ""
	if len(s) == 0 {
		return ""
	}
	if len(s) == 1 {
		return s
	}
	for i := 0; i < len(s); i++ {
		// even-length palindrome: if i==1, 0 (1) 2 3 4 5 => checking 12 0123
		for start, end := i, i+1; start >= 0 && end < len(s); start, end = start-1, end+1 {
			if s[start] != s[end] {
				break
			}
			if len(result) < end-start+1 {
				result = s[start : end+1]
			}
		}
		// odd-length: if i==2, 0 1 (2) 3 4 5 => start = 1, end = 3
		for start, end := i, i; start >= 0 && end < len(s); start, end = start-1, end+1 {
			if s[start] != s[end] {
				break
			}
			if len(result) < end-start+1 {
				result = s[start : end+1]
			}
		}
	}
	return result
}
