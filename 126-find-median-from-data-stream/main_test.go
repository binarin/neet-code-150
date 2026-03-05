package main

import (
	"testing"
)

func TestExample1(t *testing.T) {
	// Input: ["MedianFinder", "addNum", "addNum", "findMedian", "addNum", "findMedian"]
	//        [[], [1], [2], [], [3], []]
	// Output: [null, null, null, 1.5, null, 2.0]

	medianFinder := Constructor()
	medianFinder.AddNum(1)
	medianFinder.AddNum(2)

	got := medianFinder.FindMedian()
	want := 1.5
	if got != want {
		t.Errorf("FindMedian() after [1,2] = %v, want %v", got, want)
	}

	medianFinder.AddNum(3)

	got = medianFinder.FindMedian()
	want = 2.0
	if got != want {
		t.Errorf("FindMedian() after [1,2,3] = %v, want %v", got, want)
	}
}
