package main

import "fmt"

func main() {
	var strs = []string{"abcabcbb", "bbbbb", "pwwkew"}
	for _, s := range strs {
		fmt.Println(s, " -> ", lengthOfLongestSubstring(s))
	}
}

func hasDupes(s string) bool {
	m := make(map[rune]int)
	for _, chr := range s {
		m[chr] += 1
	}

	for _, val := range m {
		if val > 1 {
			return true
		}
	}
	return false
}

func lengthOfLongestSubstring(s string) int {
	longest := 0
	for start := 0; start < len(s); start++ {
		for end := start + 1; end <= len(s); end++ {
			substr := s[start:end]
			if hasDupes(substr) {
				break
			}
			longest = max(longest, len(substr))
		}
	}
	return longest
}
