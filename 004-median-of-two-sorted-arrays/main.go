package main

import "fmt"

func main() {
	fmt.Println(findMedianSortedArrays([]int{1, 3}, []int{2}))
	fmt.Println(findMedianSortedArrays([]int{1, 2}, []int{3, 4}))
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	merged := make([]int, 0, len(nums1)+len(nums2))
	p1 := 0
	p2 := 0
	for {
		if p1 == len(nums1) || p2 == len(nums2) {
			break
		}
		if nums1[p1] < nums2[p2] {
			merged = append(merged, nums1[p1])
			p1++
		} else {
			merged = append(merged, nums2[p2])
			p2++
		}
	}
	// flush either p1 or p2
	if p1 != len(nums1) {
		merged = append(merged, nums1[p1:]...)
	}
	if p2 != len(nums2) {
		merged = append(merged, nums2[p2:]...)
	}
	fmt.Println(merged)
	if len(merged) == 0 {
		return 0
	}
	if len(merged)%2 == 0 {
		return (float64(merged[len(merged)/2-1]) + float64(merged[len(merged)/2])) / 2.0
	}
	return float64(merged[len(merged)/2])
}
