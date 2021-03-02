package main

import "fmt"

// @desc 一个简单的 Node 节点数据结构

type Node struct {
	le   *Node
	data interface{}
	ri   *Node
}

func NewNode(left, right *Node) *Node {
	return &Node{left, nil, right}
}

func (n *Node) SetData(data interface{}) {
	n.data = data
}

func (n *Node) GetData() interface{} {
	return n.data
}

func main() {
	root := NewNode(nil, nil)
	root.SetData("root")

	leftChild := NewNode(nil, nil)
	leftChild.SetData("left child")

	rightChild := NewNode(nil, nil)
	rightChild.SetData("right child")

	root.le = leftChild
	root.ri = rightChild

	fmt.Println(root.GetData())
	fmt.Println(root.le.GetData())
	fmt.Println(root.ri.GetData())
}
