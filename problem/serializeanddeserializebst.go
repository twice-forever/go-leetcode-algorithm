package problem

import (
	"strconv"
	"strings"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type Codecii struct {
}

func Constructoriii() Codecii {
	return Codecii{}
}

// Serializes a tree to a single string.
func (this *Codecii) serialize(root *TreeNode) string {
	var str string
	var addStr func(node *TreeNode)
	addStr = func(node *TreeNode) {
		if node == nil {
			return
		}
		str += strconv.Itoa(node.Val) + ","
		addStr(node.Left)
		addStr(node.Right)
	}
	addStr(root)
	return str
}

// Deserializes your encoded data to tree.
func (this *Codecii) deserialize(data string) *TreeNode {
	if data == "" {
		return nil
	}

	strSlice := strings.Split(data, ",")
	strSlice = strSlice[:len(strSlice)-1]

	var rebuild func(strs []string) *TreeNode
	rebuild = func(strs []string) *TreeNode {
		if len(strs) == 0 {
			return nil
		}

		node := &TreeNode{}
		node.Val, _ = strconv.Atoi(strs[0])
		for i, v := range strs {
			iv, _ := strconv.Atoi(v)
			if iv > node.Val {
				node.Left = rebuild(strs[1:i])
				node.Right = rebuild(strs[i:])
				return node
			}
		}
		node.Left = rebuild(strs[1:])
		return node
	}
	return rebuild(strSlice)
}

/**
 * Your Codec object will be instantiated and called as such:
 * ser := Constructor()
 * deser := Constructor()
 * tree := ser.serialize(root)
 * ans := deser.deserialize(tree)
 * return ans
 */
