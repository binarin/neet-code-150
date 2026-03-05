// https://leetcode.com/problems/serialize-and-deserialize-binary-tree/
// Difficulty: Hard

package main

import "fmt"

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Codec struct {
}

func Constructor() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	return ""
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	return nil
}

func main() {
	// Example 1: root = [1,2,3,null,null,4,5]
	//       1
	//      / \
	//     2   3
	//        / \
	//       4   5
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 3}
	root.Right.Left = &TreeNode{Val: 4}
	root.Right.Right = &TreeNode{Val: 5}

	ser := Constructor()
	deser := Constructor()
	data := ser.serialize(root)
	fmt.Printf("Serialized: %s\n", data)
	ans := deser.deserialize(data)
	fmt.Printf("Deserialized root val: %v\n", ans)
}
