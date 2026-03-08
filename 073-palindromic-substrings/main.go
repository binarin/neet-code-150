// https://leetcode.com/problems/palindromic-substrings/
// Difficulty: Medium

package main

import "fmt"

func countSubstrings(s string) int {
	total := 0
	for center := range len(s) {

		left, right := center, center
		for s[left] == s[right] {
			total++
			left--
			right++
			if left < 0 || right >= len(s) {
				break
			}
		}

		left, right = center, center+1
		for {
			if left < 0 || right >= len(s) {
				break
			}
			if s[left] != s[right] {
				break
			}
			total++
			left--
			right++
		}
	}

	// count even len
	// count odd len
	return total
}

func main() {
	fmt.Println(countSubstrings("aaa"))
}
