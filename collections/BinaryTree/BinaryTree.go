package BinaryTree

import (
	"bytes"
	"math"
	"strconv"
)

type Node struct {
	Val   int
	Left  *Node
	Right *Node
}

func InitBinaryTree(val int) *Node {
	return &Node{
		val,
		nil,
		nil,
	}
}
func (t *Node) Find(val int) *Node {
	if t == nil {
		return t
	}
	if val > t.Val {
		return t.Right.Find(val)
	} else if val < t.Val {
		return t.Left.Find(val)
	} else {
		return t
	}
}

func (t *Node) FindMax() *Node {
	if t == nil {
		return nil
	}
	if t.Right != nil {
		return t.Right.FindMax()
	} else {
		return t
	}
}

func (t *Node) FindMin() *Node {
	if t == nil {
		return nil
	}
	if t.Left != nil {
		return t.Left.FindMin()
	} else {
		return t
	}
}

func (t *Node) Del(val int) *Node {

	if t == nil {
		return t
	}
	if t.Val == val {
		if t.Right == nil && t.Left == nil {
			return nil
		}
		if t.Right == nil {
			return t.Left
		} else if t.Left == nil {
			return t.Right
		} else {
			tmp := t.Left.Right
			t.Left.Right = t.Right
			t.Left.Right.FindMin().Left = tmp
			return t.Left
		}
	} else if t.Val > val {
		t.Left = t.Left.Del(val)
		return t
	} else {
		t.Right = t.Right.Del(val)
		return t
	}
}

func (t *Node) Insert(val int) {
	if val > t.Val {
		if t.Right == nil {
			t.Right = &Node {
				val,
				nil,
				nil,
			}
		} else {
			t.Right.Insert(val)
		}
	} else if val < t.Val {
		if t.Left == nil {
			t.Left = &Node {
				val,
				nil,
				nil,
			}
		} else {
			t.Left.Insert(val)
		}
	}
}

func (t *Node) Height() float64 {
	if t == nil {
		return 0
	}
	return 1 + math.Max(float64(t.Left.Height()), float64(t.Right.Height()))
}

func (t *Node) String() string {
	height := t.Height()
	var currentNodes, nextNodes []*Node
	var result []string
	currentNodes = append(currentNodes, t)
	isAllEmpty := false
	for !isAllEmpty {
		isAllEmpty = true
		var buf bytes.Buffer
		for i := range currentNodes {
			tmp := currentNodes[i]
			if tmp != nil {
				buf.WriteString(strconv.Itoa(currentNodes[i].Val))
				for j := 0; float64(j) < math.Pow(float64(2), float64(height)) / float64(len(currentNodes)); j++ {
					buf.WriteString("    ")
				}
				nextNodes = append(nextNodes, tmp.Left)
				nextNodes = append(nextNodes, tmp.Right)
				isAllEmpty = false
			} else {
				buf.WriteString(" ")
				for j := 0; float64(j) < math.Pow(float64(2), float64(height)) / float64(len(currentNodes)); j++ {
					buf.WriteString("    ")
				}
				nextNodes = append(nextNodes, nil)
				nextNodes = append(nextNodes, nil)
			}
		}
		result = append(result, buf.String())
		currentNodes = nextNodes
		nextNodes = nil
	}
	var buf bytes.Buffer
	for i := range result {
		for j := 0; float64(j) < math.Pow(float64(2), float64(len(result) - i)); j++ {
			buf.WriteString(" ")
		}
		buf.WriteString(result[i])
		buf.WriteString("\n")
	}
	return buf.String()
}
