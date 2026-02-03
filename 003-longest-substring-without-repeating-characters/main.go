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
	if len(s) == 0 {
		return 0
	}
	if len(s) == 1 {
		return 1
	}
	current_set := make(map[byte]bool)
	max_length := 0

	head := 0
	current_set[s[head]] = true

	tail := 0

outer:
	for {
		// move tail while can
		for {
			if tail == len(s)-1 {
				// already at the end
				break outer
			}
			if _, present := current_set[s[tail+1]]; present {
				// tail is blocked
				break
			}
			current_set[s[tail+1]] = true
			tail++
			max_length = max(max_length, tail-head+1)
		}

		// move head until tail unblocked
		conflicting_char := s[tail+1]
		for {
			if head == len(s)-1 {
				break outer
			}
			head_char := s[head]
			delete(current_set, s[head])
			head++
			if head_char == conflicting_char {
				break
			}
		}
	}

	return max_length
}

func lengthOfLongestSubstringBrute(s string) int {
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
