// https://leetcode.com/problems/time-based-key-value-store/
// Difficulty: Medium

package main

import "fmt"

type TimeMap struct {
}

func Constructor() TimeMap {
	return TimeMap{}
}

func (this *TimeMap) Set(key string, value string, timestamp int) {
}

func (this *TimeMap) Get(key string, timestamp int) string {
	return ""
}

func main() {
	// Example 1:
	// Input: ["TimeMap", "set", "get", "get", "set", "get", "get"]
	//        [[], ["foo", "bar", 1], ["foo", 1], ["foo", 3], ["foo", "bar2", 4], ["foo", 4], ["foo", 5]]
	// Output: [null, null, "bar", "bar", null, "bar2", "bar2"]

	timeMap := Constructor()
	timeMap.Set("foo", "bar", 1)
	fmt.Println(timeMap.Get("foo", 1))  // Expected: "bar"
	fmt.Println(timeMap.Get("foo", 3))  // Expected: "bar"
	timeMap.Set("foo", "bar2", 4)
	fmt.Println(timeMap.Get("foo", 4))  // Expected: "bar2"
	fmt.Println(timeMap.Get("foo", 5))  // Expected: "bar2"
}
