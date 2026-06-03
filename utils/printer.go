package utils

import (
	"fmt"

	"go-huffman/model"
)

func PrintTree(node *model.Node, level int) {
	_ = level // compatibilidade com a assinatura usada em main.go
	if node == nil {
		return
	}
	printNode(node, "", "", true)
}

func printNode(node *model.Node, prefix, edge string, isTail bool) {
	if node == nil {
		return
	}

	label := fmt.Sprintf("(%d)", node.Frequency)
	if node.IsLeaf() {
		label = fmt.Sprintf("%c(%d)", node.Char, node.Frequency)
	}

	if edge == "" {
		fmt.Println(label)
	} else {
		connector := "├──"
		if isTail {
			connector = "└──"
		}
		fmt.Printf("%s%s %s: %s\n", prefix, connector, edge, label)
	}

	children := []struct {
		node *model.Node
		edge string
	}{}
	if node.Left != nil {
		children = append(children, struct {
			node *model.Node
			edge string
		}{node: node.Left, edge: "0"})
	}
	if node.Right != nil {
		children = append(children, struct {
			node *model.Node
			edge string
		}{node: node.Right, edge: "1"})
	}

	nextPrefix := prefix
	if edge != "" {
		if isTail {
			nextPrefix += "    "
		} else {
			nextPrefix += "│   "
		}
	}

	for i, child := range children {
		printNode(child.node, nextPrefix, child.edge, i == len(children)-1)
	}
}
