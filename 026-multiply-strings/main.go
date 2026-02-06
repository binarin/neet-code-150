package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(multiply("2", "3"))
}

func strToSlice(s string) []byte {
	result := []byte{}
	for i := len(s) - 1; i >= 0; i-- {
		result = append(result, s[i]-'0')
	}
	return result
}

func addSlices(n1 []byte, n2 []byte) []byte {
	if len(n2) < len(n1) {
		n1, n2 = n2, n1
	}
	for len(n1) < len(n2) {
		n1 = append(n1, 0)
	}
	result := make([]byte, 0, len(n1)+1)
	var carry byte = 0
	for i := range n1 {
		dig := n1[i] + n2[i] + carry
		result = append(result, dig%10)
		carry = dig / 10
	}
	if carry > 0 {
		result = append(result, carry)
	}
	return result
}

func shr(n []byte, shift int) []byte {
	result := make([]byte, len(n)+shift)
	copy(result[shift:], n)
	return result
}

func mul1(num []byte, multiplier byte) []byte {
	result := []byte{}
	carry := byte(0)
	for _, digit := range num {
		val := digit*multiplier + carry
		result = append(result, val%10)
		carry = val / 10
	}
	if carry > 0 {
		result = append(result, carry)
	}
	return result
}

func multiply(num1 string, num2 string) string {
	n1 := strToSlice(num1)
	n2 := strToSlice(num2)

	result := []byte{}
	for power, digit := range n1 {
		result = addSlices(result, shr(mul1(n2, digit), power))
	}

	sr := ""
	for i := len(result) - 1; i >= 0; i-- {
		sr += string([]byte{result[i] + '0'})
	}

	sr = strings.TrimLeft(sr, "0")
	if len(sr) == 0 {
		sr = "0"
	}
	return sr
}
