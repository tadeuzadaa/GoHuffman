package decoder

import (
	"strings"

	"go-huffman/model"
)

func Decode(encoded string, root *model.Node) string {
	if root == nil {
		return ""
	}
	if root.IsLeaf() {
		var decoded strings.Builder
		for range encoded {
			decoded.WriteRune(root.Char)
		}
		return decoded.String()
	}

	var decoded strings.Builder
	current := root

	for _, bit := range encoded {

		// esquerda
		if bit == '0' {
			current = current.Left
		}

		// direita
		if bit == '1' {
			current = current.Right
		}
		if current == nil {
			return decoded.String()
		}

		//encontrou caractere
		if current.IsLeaf() {
			decoded.WriteRune(current.Char)

			current = root
		}
	}

	return decoded.String()
}
