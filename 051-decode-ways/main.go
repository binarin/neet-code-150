package main

import (
	"fmt"
)

func main() {
	fmt.Println(numDecodings("12"))
}

func numDecodings(s string) int {
	length := len(s)
	cache := map[int]int{}
	var recur func(int) int
	recur = func(start int) int {
		if val, ok := cache[start]; ok {
			return val
		}
		if start == length {
			return 1
		}
		first := s[start]
		if first == '0' {
			cache[start] = 0
			return 0
		}
		// decode head as single char
		decodings := recur(start + 1)
		if start+1 < length {
			second := s[start+1]
			num := (first-'0')*10 + (second - '0')
			if num <= 26 {
				decodings += recur(start + 2)
			}
		}
		cache[start] = decodings
		return decodings
	}

	return recur(0)
}
