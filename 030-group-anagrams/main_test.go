package main

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func normalizeGroups(groups [][]string) [][]string {
	for _, g := range groups {
		sort.Strings(g)
	}
	sort.Slice(groups, func(i, j int) bool {
		if len(groups[i]) != len(groups[j]) {
			return len(groups[i]) < len(groups[j])
		}
		for k := 0; k < len(groups[i]); k++ {
			if groups[i][k] != groups[j][k] {
				return groups[i][k] < groups[j][k]
			}
		}
		return false
	})
	return groups
}

func TestGroupAnagrams(t *testing.T) {
	// Example 1
	result1 := groupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"})
	expected1 := [][]string{{"bat"}, {"nat", "tan"}, {"ate", "eat", "tea"}}
	assert.Equal(t, normalizeGroups(expected1), normalizeGroups(result1))

	// Example 2
	result2 := groupAnagrams([]string{""})
	expected2 := [][]string{{""}}
	assert.Equal(t, normalizeGroups(expected2), normalizeGroups(result2))

	// Example 3
	result3 := groupAnagrams([]string{"a"})
	expected3 := [][]string{{"a"}}
	assert.Equal(t, normalizeGroups(expected3), normalizeGroups(result3))
}
