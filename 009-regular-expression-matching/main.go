package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(isMatch("aa", "a"))
}

func isMatch(s string, p string) bool {
	// let's start greedy without backtracking
	s_pos := 0
	p_pos := 0

	is_star_pattern := func() bool {
		return p_pos+1 < len(p) && p[p_pos+1] == '*'
	}
	consume := func(char_p byte, min_occurs int, max_occurs int, s_pos *int) bool {
		occurs := 0
		for occurs < min_occurs && *s_pos < len(s) {
			if char_p == '.' || char_p == s[*s_pos] {
				occurs++
				*s_pos++
			} else {
				break
			}
		}

		if occurs < min_occurs {
			return false
		}
		for occurs < max_occurs && *s_pos < len(s) {
			if char_p == '.' || char_p == s[*s_pos] {
				occurs++
				*s_pos++
			} else {
				break
			}
		}
		return true // min_occurs is already satisfied, we’ve just greadily consumed the rest
	}

	for {
		if s_pos >= len(s) || p_pos >= len(p) {
			break
		}
		char_p := p[p_pos]
		if is_star_pattern() {
			if !consume(char_p, 0, math.MaxInt, &s_pos) {
				return false
			}
			p_pos += 2
		} else {
			if !consume(char_p, 1, 1, &s_pos) {
				return false
			}
			p_pos += 1
		}
	}

	return s_pos == len(s) && p_pos == len(p)
}
