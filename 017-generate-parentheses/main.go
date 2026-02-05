package main

import (
	"fmt"
	"maps"
	"slices"
)

func main() {
	fmt.Println(generateParenthesis(3))
}

func generateParenthesis(n int) []string {
	if n == 0 {
		return []string{""}
	}
	if n == 1 {
		return []string{"()"}
	}
	subexprs := generateParenthesis(n - 1)
	result := map[string]bool{}

	for _, se := range subexprs {
		for i := 0; i <= len(se); i++ {
			result[se[0:i]+"()"+se[i:]] = true
		}
	}
	return slices.Collect(maps.Keys(result))
}
