// https://leetcode.com/problems/find-median-from-data-stream/description/?envType=problem-list-v2&envId=plakya4j
// Difficulty: Hard

package main

import "fmt"

type MedianFinder struct {
}

func Constructor() MedianFinder {
	return MedianFinder{}
}

func (this *MedianFinder) AddNum(num int) {
}

func (this *MedianFinder) FindMedian() float64 {
	return 0.0
}

func main() {
	// Example 1:
	// Input: ["MedianFinder", "addNum", "addNum", "findMedian", "addNum", "findMedian"]
	//        [[], [1], [2], [], [3], []]
	// Output: [null, null, null, 1.5, null, 2.0]

	medianFinder := Constructor()
	medianFinder.AddNum(1)    // arr = [1]
	medianFinder.AddNum(2)    // arr = [1, 2]
	fmt.Println(medianFinder.FindMedian()) // expected: 1.5
	medianFinder.AddNum(3)    // arr = [1, 2, 3]
	fmt.Println(medianFinder.FindMedian()) // expected: 2.0
}
