// https://leetcode.com/problems/two-sum-ii-input-array-is-sorted/
// Difficulty: Medium

package main

import (
	"fmt"
	"strings"
)

func twoSum(numbers []int, target int) []int {
	debug := false
	leftIdx, rightIdx := 0, len(numbers)-1

	mkRes := func(a, b int) []int {
		return []int{1 + a, 1 + b}
	}

	// we are looking in the range [leftIdx+1, rightIdx]
	// we are looking for index of: one that equals to need or the first one smaller
	// special consideration: if the smaller one have duplicates, we should pick the rightmost one
	// (ok, more optimal - not the leftmost)
	// (but maybe we can just check whether duplicate smaller ones sum up to target, and short-cut here)
	dump := func(header string, l, r, m int) {
		if !debug {
			return
		}
		valB, ptrB := strings.Builder{}, strings.Builder{}
		for i, v := range numbers {
			ptrS := ""
			if l == i {
				ptrS += "l"
			}
			if m == i {
				ptrS += "m"
			}
			if r == i {
				ptrS += "r"
			}
			valB.WriteString(fmt.Sprintf("%4d", v))
			ptrB.WriteString(fmt.Sprintf("%4s", ptrS))
		}
		fmt.Println(header)
		fmt.Println(valB.String())
		fmt.Println(ptrB.String())
	}

	for leftIdx < rightIdx {
		num1 := numbers[leftIdx]
		subL, subR, need := leftIdx+1, rightIdx, target-num1
		dump(fmt.Sprintf("== enter %d (at %d), need %d", numbers[leftIdx], leftIdx, need), leftIdx, rightIdx, -1)
		for subL <= subR {
			// 0001233344555789
			// t=6 ^   ^      ^
			subMid := subL + (subR-subL)/2
			dump("==== rtrim", subL, subR, subMid)
			if numbers[subMid] == need {
				return mkRes(leftIdx, subMid)
			}
			if numbers[subMid] > need {
				subR = subMid - 1
				continue
			}
			// we are going to discard left part + 1
			// as a special case, we should make sure that we won't discard duplicate numbers solution
			if subMid < subR && numbers[subMid] == numbers[subMid+1] && 2*numbers[subMid] == target {
				return mkRes(subMid, subMid+1)
			}
			// But we absolutely must to make progress with `subMid + 1`
			subL = subMid + 1
		}
		rightIdx = subR
		num1 = numbers[rightIdx]
		subL, subR, need = leftIdx, rightIdx-1, target-num1
		dump(fmt.Sprintf("== flip  %d (at %d), need %d", numbers[rightIdx], rightIdx, need), leftIdx, rightIdx, -1)
		for subL <= subR {
			// 0001233344555789
			// t=6 ^   ^      ^
			subMid := subL + (subR-subL)/2
			dump("==== ltrim", subL, subR, subMid)
			if numbers[subMid] == need {
				return mkRes(subMid, rightIdx)
			}
			if numbers[subMid] < need {
				subL = subMid + 1
				continue
			}
			if subL < subMid && numbers[subMid-1] == numbers[subMid] && 2*numbers[subMid] == target {
				return mkRes(subMid-1, subMid)
			}
			// But we absolutely must to make progress with `subMid + 1`
			subR = subMid - 1
		}
		leftIdx = subL
	}

	panic("precondition failed")
}

func main() {
	fmt.Println(twoSum([]int{
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		2, 3,
		9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9},
		5))
}
