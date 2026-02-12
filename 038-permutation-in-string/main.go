package main

import (
	"fmt"
)

func main() {
	fmt.Println(checkInclusion("ab", "eidbaooo"))
}

func cv(c byte) byte {
	return byte(c) - 'a'
}

type Counts [26]int

func checkInclusion(s1 string, s2 string) bool {
	if len(s2) < len(s1) {
		return false
	}
	n := len(s1)
	var need, cur Counts

	for pos := range len(s1) {
		need[cv(s1[pos])]++
		cur[cv(s2[pos])]++
	}

	// ["ab"]    ["abcdef"]
	// pos=3: ["cd"] -> -"c"(2: pos - n + 1) +"e"(pos + 1)
	// n=2
	for pos := n - 1; pos < len(s2)-1 && cur != need; pos++ {
		cur[cv(s2[pos-n+1])]--
		cur[cv(s2[pos+1])]++
	}
	return cur == need
}
