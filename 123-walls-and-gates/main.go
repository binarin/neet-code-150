// https://leetcode.com/problems/walls-and-gates/description/?envType=problem-list-v2&envId=plakya4j
// Difficulty: Medium

package main

import "fmt"

func wallsAndGates(rooms [][]int) {
}

func main() {
	rooms := [][]int{
		{2147483647, -1, 0, 2147483647},
		{2147483647, 2147483647, 2147483647, -1},
		{2147483647, -1, 2147483647, -1},
		{0, -1, 2147483647, 2147483647},
	}
	wallsAndGates(rooms)
	fmt.Println(rooms)
}
