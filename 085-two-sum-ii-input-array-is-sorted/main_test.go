package main

import (
	"reflect"
	"testing"
)

func TestTwoSum(t *testing.T) {
	tests := []struct {
		numbers  []int
		target   int
		expected []int
	}{
		{[]int{2, 7, 11, 15}, 9, []int{1, 2}},
		{[]int{2, 3, 4}, 6, []int{1, 3}},
		{[]int{-1, 0}, -1, []int{1, 2}},
	}

	for _, tt := range tests {
		result := twoSum(tt.numbers, tt.target)
		if !reflect.DeepEqual(result, tt.expected) {
			t.Errorf("twoSum(%v, %d) = %v, expected %v", tt.numbers, tt.target, result, tt.expected)
		}
	}
}
