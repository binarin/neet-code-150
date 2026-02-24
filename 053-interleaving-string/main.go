package main

import (
	"fmt"
	"strings"
)

func main() {
	// fmt.Println(isInterleave("aabcc", "dbbca", "aadbbcbcac"))
	// fmt.Println(isInterleave("a", "b", "a"))
	fmt.Println(isInterleave("ab", "bc", "babc"))
}

func isInterleave(s1 string, s2 string, s3 string) bool {
	use := func(...any) {

	}

	cache := map[[3]int]int{}
	cacheKey := func(idx1, idx2 int, swap bool) [3]int {
		swapVal := 0
		if swap {
			swapVal = 1
		}
		return [3]int{idx1, idx2, swapVal}
	}

	l1, l2, l3 := len(s1), len(s2), len(s3)

	var recur func(int, int, int, bool, int) bool

	recur = func(idx1, idx2, idx3 int, swap bool, level int) (result bool) {
		cidx := cacheKey(idx1, idx2, swap)
		if res, ok := cache[cidx]; ok {
			if res > 0 {
				return true
			} else if res < 0 {
				return false
			}
		}

		printPrefix := strings.Repeat("=", level) + " "
		ss1, ll1 := s1, l1
		ss2, ll2 := s2, l2

		if swap {
			// idx1 + idx2 = idx3
			printPrefix += "(SW:1) >"
			ss1, ll1 = s2, l2
			ss2, ll2 = s1, l1
		} else {
			printPrefix += "(SW:0) >"
		}

		defer (func() {
			val := -1
			if result {
				val = 1
			}
			cache[cidx] = val
			// fmt.Printf("%s returning %t, cache at %v set = %v\n", printPrefix, result, cidx, cache)
		})()

		use(ss2, printPrefix)
		// fmt.Printf("%s'%s'(%s[%d:]), '%s'(%s[%d:]), '%s'(%s[%d:]), %t, cidx %d, %v\n", printPrefix, ss1[idx1:], ss1, idx1, ss2[idx2:], ss2, idx2, s3[idx3:], s3, idx3, swap, cidx, cache)

		if idx3 == l3 {
			return idx1 == ll1 && idx2 == ll2
		}

		if idx1 == ll1 {
			return false
		}

		if idx2 == ll2 {
			if ss1[idx1:] == s3[idx3:] {
				return true
			} else {
				return false
			}
		}

		for delta := 0; idx1+delta < ll1 && idx3+delta < l3 && ss1[idx1+delta] == s3[idx3+delta]; delta++ {
			if recur(idx2, idx1+1+delta, idx3+1+delta, !swap, level+1) {
				return true
			}
		}

		return false
	}

	return recur(0, 0, 0, false, 0) || recur(0, 0, 0, true, 0)
}
