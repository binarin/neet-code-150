package main

import (
	"fmt"
)

func main() {
	fmt.Println(letterCombinations("23"))
}

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{""}
	}
	var first string
	switch digits[0] {
	case '2':
		first = "abc"[:]
		break
	case '3':
		first = "def"[:]
		break
	case '4':
		first = "ghi"[:]
		break
	case '5':
		first = "jkl"[:]
		break
	case '6':
		first = "mno"[:]
		break
	case '7':
		first = "pqrs"[:]
		break
	case '8':
		first = "tuv"[:]
		break
	case '9':
		first = "wxyz"[:]
		break
	default:
		panic("not 2-9")
	}
	result := []string{}
	subcombinations := letterCombinations(digits[1:])
	for i := 0; i < len(first); i++ {
		for _, sub := range subcombinations {
			result = append(result, first[i:i+1]+sub)
		}
	}
	return result
}
