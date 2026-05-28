package utils

import (
	"fmt"
	"go-huffman/model"
	"strings"
)

func PrintTree(node *model.Node, level int){

	if node == nil{
		return
	}

	// direita
	PrintTree(node.Right, level+1)

	//espaço
	space := strings.Repeat("   ", level)

	//folha
	if node.IsLeaf(){
		fmt.Printf(
			"%s%c(%d)\n",
			space,
			node.Char,
			node.Frequency,
		)
	} else {
		fmt.Printf(
			"%s(%d)\n",
			space,
			node.Frequency,
		)
	}

	// esquerda
	PrintTree(node.Left, level+1)
}