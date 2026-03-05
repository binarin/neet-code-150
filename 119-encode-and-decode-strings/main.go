// https://leetcode.com/problems/encode-and-decode-strings/description/?envType=problem-list-v2&envId=plakya4j
// Difficulty: Medium

package main

import "fmt"

type Codec struct {
}

// Encodes a list of strings to a single string.
func (codec *Codec) Encode(strs []string) string {
	return ""
}

// Decodes a single string to a list of strings.
func (codec *Codec) Decode(strs string) []string {
	return nil
}

// Your Codec object will be instantiated and called as such:
// var codec Codec
// codec.Decode(codec.Encode(strs));

func main() {
	var codec Codec

	// Example 1
	strs1 := []string{"Hello", "World"}
	encoded1 := codec.Encode(strs1)
	decoded1 := codec.Decode(encoded1)
	fmt.Printf("Input: %v\n", strs1)
	fmt.Printf("Encoded: %q\n", encoded1)
	fmt.Printf("Decoded: %v\n", decoded1)
}
