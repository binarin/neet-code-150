package main

import (
	"fmt"
	"slices"
)

func main() {
	fmt.Println(subsetsWithDup([]int{4, 4, 4, 1, 4}))
}

func subsetsWithDup(nums []int) [][]int {
	frequencies := map[int]int{}

	for _, num := range nums {
		frequencies[num]++
	}

	freqs, freqVals := make([]int, 0, len(nums)), make([]int, 0, len(nums))
	for num, f := range frequencies {
		freqs = append(freqs, f)
		freqVals = append(freqVals, num)
	}
	// fmt.Println(freqs, freqVals)

	total := 1
	for _, f := range freqs {
		total *= f + 1
	}
	// fmt.Println(total)

	result := make([][]int, 0, total)

	for mask := range total {
		// fmt.Println("======= ", mask)
		subset := []int{}
		freqIdx := 0
		for mask > 0 {
			power := freqs[freqIdx] + 1
			occurences := mask % power
			// fmt.Printf("m %d, fi %d, pwr %d, occ %d, ss %v\n", mask, freqIdx, power, occurences, subset)
			subset = append(subset, slices.Repeat([]int{freqVals[freqIdx]}, occurences)...)
			freqIdx++
			mask /= power
		}

		result = append(result, subset)
	}

	return result
}
