package main

import "testing"

func TestTimeMap_Example1(t *testing.T) {
	// Input: ["TimeMap", "set", "get", "get", "set", "get", "get"]
	//        [[], ["foo", "bar", 1], ["foo", 1], ["foo", 3], ["foo", "bar2", 4], ["foo", 4], ["foo", 5]]
	// Output: [null, null, "bar", "bar", null, "bar2", "bar2"]

	timeMap := Constructor()

	timeMap.Set("foo", "bar", 1)

	if got := timeMap.Get("foo", 1); got != "bar" {
		t.Errorf("Get(foo, 1) = %q, want %q", got, "bar")
	}

	if got := timeMap.Get("foo", 3); got != "bar" {
		t.Errorf("Get(foo, 3) = %q, want %q", got, "bar")
	}

	timeMap.Set("foo", "bar2", 4)

	if got := timeMap.Get("foo", 4); got != "bar2" {
		t.Errorf("Get(foo, 4) = %q, want %q", got, "bar2")
	}

	if got := timeMap.Get("foo", 5); got != "bar2" {
		t.Errorf("Get(foo, 5) = %q, want %q", got, "bar2")
	}
}
