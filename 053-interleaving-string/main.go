package main

import (
	"fmt"
)

func main() {
	fmt.Println(isInterleave("aabcc", "dbbca", "aadbbcbcac"))
	// fmt.Println(isInterleave("a", "b", "a"))
}

func isInterleave(s1 string, s2 string, s3 string) bool {
	abs := func(a int) int {
		if a < 0 {
			a = -a
		}
		return a
	}

	var recur func(string, string, string, int) bool

	recur = func(ss1, ss2, ss3 string, delta int) bool {
		// fmt.Printf("'%s', '%s', '%s', %d\n", ss1, ss2, ss3, delta)

		if len(ss3) == 0 {
			return abs(delta) <= 1 && len(ss1) == 0 && len(ss2) == 0
		}
		if len(ss1) == 0 {
			return false
		}

		for i1, i3 := 0, 0; i1 < len(ss1) && i3 < len(ss3) && ss1[i1] == ss3[i3]; i1, i3 = i1+1, i3+1 {
			if recur(ss2, ss1[i1+1:], ss3[i3+1:], -delta+1) {
				return true
			}
		}

		return false
	}

	return recur(s1, s2, s3, 0) || recur(s2, s1, s3, 0)
}
