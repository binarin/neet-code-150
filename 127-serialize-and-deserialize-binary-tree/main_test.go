package main

import (
	"reflect"
	"testing"
)

// Helper function to build tree from level-order array
func buildTree(vals []interface{}) *TreeNode {
	if len(vals) == 0 || vals[0] == nil {
		return nil
	}

	root := &TreeNode{Val: vals[0].(int)}
	queue := []*TreeNode{root}
	i := 1

	for len(queue) > 0 && i < len(vals) {
		node := queue[0]
		queue = queue[1:]

		if i < len(vals) && vals[i] != nil {
			node.Left = &TreeNode{Val: vals[i].(int)}
			queue = append(queue, node.Left)
		}
		i++

		if i < len(vals) && vals[i] != nil {
			node.Right = &TreeNode{Val: vals[i].(int)}
			queue = append(queue, node.Right)
		}
		i++
	}

	return root
}

// Helper function to convert tree to level-order array
func treeToArray(root *TreeNode) []interface{} {
	if root == nil {
		return []interface{}{}
	}

	result := []interface{}{}
	queue := []*TreeNode{root}

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		if node == nil {
			result = append(result, nil)
		} else {
			result = append(result, node.Val)
			queue = append(queue, node.Left)
			queue = append(queue, node.Right)
		}
	}

	// Remove trailing nils
	for len(result) > 0 && result[len(result)-1] == nil {
		result = result[:len(result)-1]
	}

	return result
}

func TestSerializeDeserialize(t *testing.T) {
	tests := []struct {
		name string
		vals []interface{}
	}{
		{
			name: "Example 1: [1,2,3,null,null,4,5]",
			vals: []interface{}{1, 2, 3, nil, nil, 4, 5},
		},
		{
			name: "Example 2: empty tree",
			vals: []interface{}{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			root := buildTree(tt.vals)
			ser := Constructor()
			deser := Constructor()
			data := ser.serialize(root)
			result := deser.deserialize(data)
			resultVals := treeToArray(result)

			if !reflect.DeepEqual(resultVals, tt.vals) {
				t.Errorf("serialize/deserialize() = %v, want %v", resultVals, tt.vals)
			}
		})
	}
}
