package encoder

import (
	"strings"

	"go-huffman/model"
)

func GenerateCodes(
	node *model.Node,
	prefix string,
	codes map[rune]string,
) {
	if node == nil {
		return
	}

	// encontrou caractere
	if node.IsLeaf() {
		if prefix == "" {
			prefix = "0"
		}
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
		node.Right,
		prefix+"1",
		codes,
	)

}

func Encode(text string, codes map[rune]string) string {
	var encoded strings.Builder

	for _, char := range text {
		encoded.WriteString(codes[char])
	}

	return encoded.String()
}
