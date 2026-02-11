package main

import (
	"fmt"
)

func main() {
	fmt.Println(insert([][]int{{1, 5}}, []int{2, 3}))
}

func insert(intervals [][]int, newInterval []int) [][]int {
	if len(intervals) == 0 {
		return [][]int{newInterval}
	}
	maybe_join := func(i1, i2 []int) (bool, []int) {
		if i1[0] > i2[0] {
			i1, i2 = i2, i1
		}
		if i1[1] >= i2[0] {
			return true, []int{i1[0], max(i1[1], i2[1])}
		} else {
			return false, nil
		}
	}
	result := make([][]int, 0, len(intervals)+1)
	for _, ival := range intervals {
		if newInterval != nil {
			intersects, joined := maybe_join(ival, newInterval)
			if intersects {
				newInterval = joined
				continue
			}
		}
		if newInterval != nil && newInterval[1] < ival[0] {
			result = append(result, newInterval)
			newInterval = nil
		}
		result = append(result, ival)
	}
	if newInterval != nil {
		result = append(result, newInterval)
	}
	return result
}
