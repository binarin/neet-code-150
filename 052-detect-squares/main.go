package main

import "fmt"

type DetectSquares struct {
	XLookup map[int]map[int]int
	YLookup map[int]map[int]int
}

func Constructor() DetectSquares {
	return DetectSquares{
		XLookup: map[int]map[int]int{},
		YLookup: map[int]map[int]int{},
	}
}

func (d *DetectSquares) Add(point []int) {
	if _, ok := d.XLookup[point[0]]; !ok {
		d.XLookup[point[0]] = map[int]int{}
	}
	if _, ok := d.YLookup[point[1]]; !ok {
		d.YLookup[point[1]] = map[int]int{}
	}
	d.XLookup[point[0]][point[1]]++
	d.YLookup[point[1]][point[0]]++
}

func abs(a int) int {
	if a < 0 {
		return -a
	} else {
		return a
	}
}

func (d *DetectSquares) Count(point []int) int {
	// given count(cx, cy)
	// for all points in row cy -> (p1x, cy), p1x /= cx (also need p1count)
	// for all points in col cx -> (cx, p2y), p2y /= cy (also need p2count)
	// if exists p3 = (p1x, p2y): result += p1count * p2count * p3count
	result := 0
	cx, cy := point[0], point[1]
	for p1x, p1count := range d.YLookup[cy] {
		if p1x == cx {
			continue
		}
		side := abs(p1x - cx)
		for p2y, p2count := range d.XLookup[cx] {
			if abs(cy-p2y) != side {
				continue
			}
			if p3count, ok := d.XLookup[p1x][p2y]; ok {
				result += p1count * p2count * p3count
			}
		}
	}
	return result
}

/**
 * Your DetectSquares object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(point);
 * param_2 := obj.Count(point);
 */

func main() {
	// Example 1
	detectSquares := Constructor()
	detectSquares.Add([]int{5, 10})
	detectSquares.Add([]int{10, 5})
	detectSquares.Add([]int{10, 10})
	// fmt.Println(detectSquares.Count([]int{5, 5})) // expected: 1
	detectSquares.Add([]int{3, 0})
	detectSquares.Add([]int{8, 0})
	detectSquares.Add([]int{8, 5})
	// fmt.Println(detectSquares.Count([]int{3, 5})) // expected: 1
	detectSquares.Add([]int{9, 0})
	detectSquares.Add([]int{9, 8})
	detectSquares.Add([]int{1, 8})
	// fmt.Println(detectSquares.Count([]int{1, 0})) // expected: 1
	detectSquares.Add([]int{0, 0})
	detectSquares.Add([]int{8, 0})
	detectSquares.Add([]int{8, 8})
	fmt.Println(detectSquares.Count([]int{0, 8})) // expected: 2
}
