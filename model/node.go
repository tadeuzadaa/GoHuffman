package model

type Node struct{
	Char 		rune
	Frequency 	int

	Left *Node
	Right *Node
}

func (n *Node) IsLeaf() bool {
	return n.Left == nil && n.Right == nil
}