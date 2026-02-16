package main

import (
	"fmt"
)

func main() {
	fmt.Println(minWindow("ab", "A"))
	// fmt.Println(minWindow("ADOBECODEBANC", "ABC"))
}

func minWindow(s string, t string) string {
	charIdx := func(v byte) byte {
		if 'a' <= v && v <= 'z' {
			return v - 'a'
		}
		if 'A' <= v && v <= 'Z' {
			return v - 'A' + 26
		}
		panic("Out of range char")
	}
	count := func(s string) [52]byte {
		result := [52]byte{}
		for i := 0; i < len(s); i++ {
			result[charIdx(s[i])]++
		}
		return result
	}
	target := count(t)
	enough := func(ct [52]byte) bool {
		for i, v := range ct {
			if target[i] > v {
				return false
			}
		}
		return true
	}
	shortest := s + "!"
	for i := 0; i <= len(s)-len(t); i++ {
		for j := i + len(t); j <= len(s); j++ {
			substr := s[i:j]
			if enough(count(substr)) && len(substr) < len(shortest) {
				shortest = substr
			}
		}
	}
	if len(shortest) > len(s) {
		return ""
	}
	return shortest
}
