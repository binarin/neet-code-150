// https://leetcode.com/problems/kth-largest-element-in-a-stream/
// Difficulty: Easy

package main

import "fmt"

type KthLargest struct {
}

func Constructor(k int, nums []int) KthLargest {
	return KthLargest{}
}

func (this *KthLargest) Add(val int) int {
	return 0
}

func main() {
	// Example 1:
	// Input: ["KthLargest", "add", "add", "add", "add", "add"]
	//        [[3, [4, 5, 8, 2]], [3], [5], [10], [9], [4]]
	// Output: [null, 4, 5, 5, 8, 8]

	kthLargest := Constructor(3, []int{4, 5, 8, 2})
	fmt.Println(kthLargest.Add(3))  // expected: 4
	fmt.Println(kthLargest.Add(5))  // expected: 5
	fmt.Println(kthLargest.Add(10)) // expected: 5
	fmt.Println(kthLargest.Add(9))  // expected: 8
	fmt.Println(kthLargest.Add(4))  // expected: 8
}
