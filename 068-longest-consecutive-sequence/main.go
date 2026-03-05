// https://leetcode.com/problems/longest-consecutive-sequence/
// Difficulty: Medium

package main

import (
	"fmt"
	"maps"
	"slices"
	"strings"
)

type Sequences struct {
	Seen       map[int]bool
	StartToEnd map[int]int
	EndToStart map[int]int
}

func NewSequences() Sequences {
	return Sequences{
		Seen:       map[int]bool{},
		StartToEnd: map[int]int{},
		EndToStart: map[int]int{},
	}
}

func (ss Sequences) String() string {
	if len(ss.StartToEnd) == 0 {
		return "<EMPTY>"
	}
	result := strings.Builder{}
	starts := slices.Collect(maps.Keys(ss.StartToEnd))
	slices.Sort(starts)
	for i, k := range starts {
		if i != 0 {
			result.WriteString(", ")
		}
		result.WriteString(fmt.Sprintf("[%d,%d (len: %d)]", k, ss.StartToEnd[k], ss.StartToEnd[k]-k+1))
	}
	return result.String()
}

func longestConsecutive(nums []int) int {
	ss := NewSequences()
	longest := 0
	for _, num := range nums {
		if _, ok := ss.Seen[num]; ok {
			continue
		}
		ss.Seen[num] = true
		// fmt.Println(">>> Got", num, "longest", longest)
		// fmt.Println(ss)
		rightIntervalStart := num + 1
		rightIntervalEnd, hasRightInterval := ss.StartToEnd[rightIntervalStart]
		leftIntervalEnd := num - 1
		leftIntervalStart, hasLeftInteral := ss.EndToStart[num-1]
		if hasRightInterval && hasLeftInteral {
			delete(ss.EndToStart, leftIntervalEnd)
			delete(ss.StartToEnd, rightIntervalStart)
			ss.StartToEnd[leftIntervalStart] = rightIntervalEnd
			ss.EndToStart[rightIntervalEnd] = leftIntervalStart
			longest = max(longest, rightIntervalEnd-leftIntervalStart+1)
		} else if hasLeftInteral {
			ss.StartToEnd[leftIntervalStart] = num
			delete(ss.EndToStart, leftIntervalEnd)
			ss.EndToStart[num] = leftIntervalStart
			longest = max(longest, num-leftIntervalStart+1)
		} else if hasRightInterval {
			delete(ss.StartToEnd, rightIntervalStart)
			ss.StartToEnd[num] = rightIntervalEnd
			ss.EndToStart[rightIntervalEnd] = num
			longest = max(longest, rightIntervalEnd-num+1)
		} else {
			ss.StartToEnd[num] = num
			ss.EndToStart[num] = num
			longest = max(longest, 1)
		}
		// fmt.Println(ss)
		// fmt.Println("<<< Max", longest)
	}
	return longest
}

func main() {
	// nums := []int{100, 4, 200, 1, 3, 2}
	// nums := []int{1, 0, 1, 2}
	nums := []int{0, -1}
	fmt.Println(longestConsecutive(nums))
}
