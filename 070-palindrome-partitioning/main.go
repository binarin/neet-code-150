// https://leetcode.com/problems/palindrome-partitioning/
// Difficulty: Medium

package main

import "fmt"

func partition(s string) [][]string {
	cache := map[string][][]string{}

	var recur func(string) [][]string
	recur = func(ss string) [][]string {
		if _, ok := cache[ss]; ok {
			return cache[ss]
		}
		if len(ss) == 0 {
			return [][]string{{}}
		}
		isPalindrome := func(l int) bool {
			if l == 1 {
				return true
			}
			left, right := 0, l-1
			for left <= right {
				if ss[left] != ss[right] {
					return false
				}
				left++
				right--
			}
			return true
		}
		result := [][]string{}
		for pLen := 1; pLen <= len(ss); pLen++ {
			if isPalindrome(pLen) {
				prefix := ss[0:pLen]
				for _, part := range partition(ss[pLen:]) {
					extended := append([]string{prefix}, part...)
					result = append(result, extended)
				}
			}
		}
		cache[ss] = result
		return result
	}

	return recur(s)
}

func main() {
	// Example 1: s = "aab"
	result := partition("aab")
	fmt.Println(result)
}
