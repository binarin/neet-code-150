// https://leetcode.com/problems/best-time-to-buy-and-sell-stock/
// Difficulty: Easy

package main

import (
	"fmt"
	"math"
)

func maxProfit(prices []int) int {
	minPrice := math.MaxInt
	best := 0
	for _, price := range prices {
		thisProfit := price - minPrice
		best = max(best, thisProfit)
		minPrice = min(minPrice, price)
	}
	return best
}

func main() {
	// Example 1: prices = [7,1,5,3,6,4], Output: 5
	prices := []int{7, 1, 5, 3, 6, 4}
	fmt.Println(maxProfit(prices))
}
