package main

import (
	"fmt"
)

func main() {
	fmt.Println(combinationSum([]int{2, 3, 6, 7}, 7))
}

func combinationSum(candidates []int, target int) [][]int {
	result := [][]int{}
	type CacheKey [41]int // candidates position, remainder, frequencies for candidates ( 2 <= cand <= 40 )
	cache := map[CacheKey]bool{}

	make_cache_key := func(cs int, rem int, freqs *CacheKey) CacheKey {
		new_key := [41]int{}
		new_key[0] = cs
		new_key[1] = rem
		copy(new_key[2:], freqs[2:])
		return new_key
	}

	add_combination := func(freqs *CacheKey) {
		combination := []int{}
		for i := 2; i < 41; i++ {
			for rep := 0; rep < freqs[i]; rep++ {
				combination = append(combination, i)
			}
		}
		result = append(result, combination)
	}

	var recur func(int, int, CacheKey)
	recur = func(cs int, rem int, so_far CacheKey) {
		cache_key := make_cache_key(cs, rem, &so_far)
		if cache[cache_key] {
			return
		}
		cache[cache_key] = true

		if rem < 0 {
			return
		}
		if rem == 0 {
			add_combination(&cache_key)
			return
		}
		if cs == len(candidates) {
			return
		}

		recur(cs+1, rem, cache_key)
		cache_key[candidates[cs]] += 1
		recur(cs, rem-candidates[cs], cache_key)
	}
	recur(0, target, CacheKey{})
	return result
}
