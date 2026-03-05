// https://leetcode.com/problems/cheapest-flights-within-k-stops/
// Medium

package main

import "fmt"

func findCheapestPrice(n int, flights [][]int, src int, dst int, k int) int {
	return 0
}

func main() {
	// Example 1: n=4, flights=[[0,1,100],[1,2,100],[2,0,100],[1,3,600],[2,3,200]], src=0, dst=3, k=1
	// Expected: 700
	n := 4
	flights := [][]int{{0, 1, 100}, {1, 2, 100}, {2, 0, 100}, {1, 3, 600}, {2, 3, 200}}
	src := 0
	dst := 3
	k := 1
	fmt.Println(findCheapestPrice(n, flights, src, dst, k))
}
