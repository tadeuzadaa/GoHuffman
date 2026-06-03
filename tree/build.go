package tree

import (
	"container/heap"
	"go-huffman/model"
)

func BuildTree(freq map[rune]int) *model.Node {

	p := &PriorityQueue{}
	heap.Init(p)

	// adiciona caracter no heap
	for char, frequency := range freq {
		node := &model.Node{
			Char:      char,
			Frequency: frequency,
		}
		heap.Push(p, node)
	}

	//monta a arvore
	for p.Len() > 1 {
		left := heap.Pop(p).(*model.Node)
		right := heap.Pop(p).(*model.Node)

		parent := &model.Node{
			Frequency: left.Frequency + right.Frequency,
			Left:      left,
			Right:     right,
		}
		heap.Push(p, parent)
	}
	if p.Len() == 0 {
		return nil
	}

	// raiz final
	return heap.Pop(p).(*model.Node)
}
