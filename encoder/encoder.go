package encoder

import "go-huffman/model"

func GenerateCodes(
	node *model.Node,
	prefix string,
	codes map[rune]string,
){
	if node == nil{
		return
	}

	// encontrou caractere
	if node.IsLeaf(){
		codes[node.Char] = prefix
		return
	}

	// esquerda recebe 0
	GenerateCodes(
		node.Left,
		prefix+"0",
		codes,
	)

	// direita recebe 1
	GenerateCodes(
		node.Left,
		prefix+"1",
		codes,
	)
	
}

func Encode(text string, codes map[rune]string) string{
	encoded := ""

	for _, char := range text {
		encoded += codes[char]
	}

	return  encoded
}