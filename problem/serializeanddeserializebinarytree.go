package problem

import (
	"fmt"
	"strconv"
	"strings"
)

var (
	TestNode = &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
		},
		Right: &TreeNode{
			Val: 3,
			Left: &TreeNode{
				Val: 4,
				Left: &TreeNode{
					Val: 6,
				},
				Right: &TreeNode{
					Val: 7,
				},
			},
			Right: &TreeNode{
				Val: 5,
			},
		},
	}
)

type Codec struct {
}

func Constructorii() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	if root == nil {
		return "null,"
	}
	var str string
	nodeSlice := []*TreeNode{root}
	for len(nodeSlice) > 0 {
		var tmp []*TreeNode
		for _, v := range nodeSlice {
			if v == nil {
				str += "null,"
				continue
			}
			str += fmt.Sprintf("%d,", v.Val)
			tmp = append(tmp, v.Left)
			tmp = append(tmp, v.Right)
		}
		nodeSlice = tmp
	}
	return str
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	strSlice := strings.Split(data, ",")
	strSlice = strSlice[:len(strSlice)-1]
	if strSlice[0] == "null" {
		return nil
	}

	root := &TreeNode{}
	val := root
	var nodeSlice []*TreeNode
	for i := 0; i < len(strSlice); i++ {
		val.Val, _ = strconv.Atoi(strSlice[i])
		for si := range nodeSlice {
			if i%2 != 0 {
				if strSlice[i] != "null" {
					nodeSlice[si].Left = val
				}
				break
			} else {
				if strSlice[i] != "null" {
					nodeSlice[si].Right = val
				}
				nodeSlice = nodeSlice[1:]
				break
			}
		}
		if strSlice[i] != "null" {
			nodeSlice = append(nodeSlice, val)
		}
		val = &TreeNode{}
	}
	return root
}

func RunTest() {
	codec := Constructorii()

	fmt.Println(codec.serialize(TestNode))
	fmt.Println(codec.deserialize(codec.serialize(TestNode)))
	fmt.Println(codec.serialize(codec.deserialize(codec.serialize(TestNode))))
}
