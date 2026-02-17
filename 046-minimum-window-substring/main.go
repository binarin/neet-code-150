package main

import (
	"fmt"
	"maps"
	"math"
	"strings"
)

func main() {
	fmt.Println(minWindow("a", "a"))
	// fmt.Println(minWindow("cabwefgewcwaefgcf", "cae"))
	// fmt.Println(minWindow("ADOBECODEBANC", "ABC"))
}

func use(...any) {
}

func minWindow(s string, t string) string {
	if len(s) < len(t) || len(s) == 0 || len(t) == 0 {
		return ""
	}

	targetFreq := map[byte]int{}
	targetLen := len(t)
	for i := range targetLen {
		targetFreq[t[i]]++
	}

	shortestLen, shortestWindow := math.MaxInt, ""

	// invariant: chopLeft + currentLen + chopRight = haystackLen
	chopLeft, chopRight, haystackLen := 0, len(s), len(s)

	extra := map[byte]int{}

	freq_fmt := func(f map[byte]int) string {
		result := []string{}
		for k, v := range f {
			result = append(result, fmt.Sprintf("%c:%d", k, v))
		}
		return strings.Join(result, ",")
	}
	use(freq_fmt)

	visualize := func(prefix string) {
		// fmt.Printf("%s: %s→%s←%s (extra: %s)\n", prefix, s[:chopLeft], s[chopLeft:haystackLen-chopRight], s[haystackLen-chopRight:], freq_fmt(extra))
	}

	needed := maps.Clone(targetFreq)

	extendRight := func() bool {
		for len(needed) > 0 && chopRight > 0 {
			// chopLeft + currentLen = chopLeft + haystackLen - chopLeft - chopRight = haystackLen - chopRight
			addCharIdx := haystackLen - chopRight
			chopRight--
			if _, ok := targetFreq[s[addCharIdx]]; ok {
				if _, ok := needed[s[addCharIdx]]; ok {
					needed[s[addCharIdx]]--
					if needed[s[addCharIdx]] == 0 {
						delete(needed, s[addCharIdx])
					}
				} else {
					extra[s[addCharIdx]]++
				}
			}
		}
		return len(needed) == 0
	}

	trimLeft := func() bool {
		for chopLeft+targetLen <= haystackLen {
			dropCharIdx := chopLeft
			if _, ok := targetFreq[s[dropCharIdx]]; ok {
				if val, ok := extra[s[dropCharIdx]]; ok {
					if val > 1 {
						extra[s[dropCharIdx]]--
					} else {
						delete(extra, s[dropCharIdx])
					}
				} else {
					return true
				}
			}
			chopLeft++

		}
		return false
	}

	maybeImprove := func() {
		currentLen := haystackLen - chopLeft - chopRight
		if currentLen < shortestLen {
			shortestLen, shortestWindow = currentLen, s[chopLeft:haystackLen-chopRight]
		}
	}

	for {
		visualize("LO")
		if !extendRight() {
			break
		}
		visualize("EX")
		if !trimLeft() {
			break
		}
		visualize("TR")
		maybeImprove()
		needed[s[chopLeft]]++
		chopLeft++
	}

	return shortestWindow
}
