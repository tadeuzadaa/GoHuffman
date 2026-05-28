package tree

import "go-huffman/model"

type PriorityQueue []*model.Node

func (p PriorityQueue) Len() int{
	// quantidade de elemento do heap
	return len(p)
}

func (p PriorityQueue) Less(i,j int) bool{
	// menor frequencia
	return p[i].Frequency < p[j].Frequency
}

func (p PriorityQueue) Swap (i, j int) {
	// inverte os elementos
	p[i], p[j] = p[j], p[i]
}

func (p *PriorityQueue)Push(x interface{}){
	// adciona o elemento e organiza
	node := x.(*model.Node)

	*p = append(*p, node)
}

func (p *PriorityQueue) Pop() interface{} {
	// jogar o menor pro final e remove
	old := *p

	n := len(old)

	item := old[n-1]
	
	*p = old[0 : n-1]

	return item
}


