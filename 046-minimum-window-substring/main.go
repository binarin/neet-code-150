package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(minWindow("cabwefgewcwaefgcf", "cae"))
	// fmt.Println(minWindow("ADOBECODEBANC", "ABC"))
}

func minWindow(s string, t string) string {
	charIdx := func(v byte) byte {
		if 'a' <= v && v <= 'z' {
			return v - 'a'
		}
		if 'A' <= v && v <= 'Z' {
			return v - 'A' + 26
		}
		panic("Out of range char")
	}
	type Window [52]byte

	current := Window{}
	for i := 0; i < len(s); i++ {
		current[charIdx(s[i])]++
	}

	target, target_length := Window{}, len(t)
	for i := range target_length {
		c_idx := charIdx(t[i])
		target[c_idx]++
		if target[c_idx] > current[c_idx] {
			return ""
		}
	}

	shortest_len, shortest_window := math.MaxInt, ""

	// _dump_win := func(w Window) []string {
	// 	result := []string{}
	// 	for idx, count := range target {
	// 		if count > 0 {
	// 			var val byte
	// 			if idx < 26 {
	// 				val = byte(idx) + 'a'
	// 			} else {
	// 				val = byte(idx) + 'A'
	// 			}
	// 			result = append(result, fmt.Sprintf("%c: %d", val, w[idx]))
	// 		}
	// 	}
	// 	return result
	// }

	var recur func(int, int, Window)
	recur = func(left, right int, current Window) {
		for target[charIdx(s[left])] == 0 && right-left+1 >= target_length {
			left++
		}
		for target[charIdx(s[right])] == 0 && right-left+1 >= target_length {
			right--
		}

		substr := s[left : right+1]
		current_len := right - left + 1
		// fmt.Println(left, right, dump_win(current), shortest_len, shortest_window, current_len, substr)

		if current_len < target_length {
			panic("shouldn't be here")
		}

		if current_len < shortest_len {
			shortest_len, shortest_window = current_len, substr
		}

		left_char_idx := charIdx(s[left])
		if current[left_char_idx] > target[left_char_idx] {
			sub := current
			sub[left_char_idx]--
			recur(left+1, right, sub)
		}

		right_char_idx := charIdx(s[right])
		if current[right_char_idx] > target[right_char_idx] {
			sub := current
			sub[right_char_idx]--
			recur(left, right-1, sub)
		}
	}

	recur(0, len(s)-1, current)
	return shortest_window
}
