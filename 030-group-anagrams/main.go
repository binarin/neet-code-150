package main

import (
	"fmt"
	"maps"
	"slices"
)

func main() {
	fmt.Println(groupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"}))
}

func groupAnagrams(strs []string) [][]string {
	groups := map[string][]string{}
	for _, str := range strs {
		as_slice := make([]byte, len(str))
		copy(as_slice, str)
		slices.Sort(as_slice)
		var normalized string = string(as_slice)
		groups[normalized] = append(groups[normalized], str)
	}

	return slices.Collect(maps.Values(groups))
}
