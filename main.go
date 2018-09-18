package main

import (
	"fmt"
)

type Node struct {
	Val   int
	Data  string
	Left  *Node
	Right *Node
}

func New(k int, data string) *Node {
	n := Insert(nil, k, data)
	return n
}

func Insert(n *Node, k int, data string) *Node {
	if n == nil {
		return &Node{k, data, nil, nil}
	}
	if k < n.Val {
		n.Left = Insert(n.Left, k, data)
	} else {
		n.Right = Insert(n.Right, k, data)
	}
	return n
}

func (n *Node) String() string {
	if n == nil {
		return "()"
	}
	s := ""
	if n.Left != nil {
		s += n.Left.String() + " "
	}
	s += fmt.Sprint(n.Val)
	if n.Right != nil {
		s += " " + n.Right.String()
	}
	return "(" + s + ")"
}

// func Walk(n *Node) *Node {
// 	if n == nil {
// 		return n
// 	}
// 	if n.Left != nil {
// 		fmt.Println(n)
// 	}
// }

func main() {
	data := []map[int]string{
		map[int]string{1: "test1"},
		map[int]string{3: "test3"},
		map[int]string{7: "test7"},
		map[int]string{5: "test5"},
	}

	node := New(2, "test2")
	for _, d := range data {
		for k, v := range d {
			node = Insert(node, k, v)
		}
	}

	fmt.Println(node.String())
}
