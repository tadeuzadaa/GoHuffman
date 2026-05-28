package decoder

import "go-huffman/model"

func Decode( encoded string, root *model.Node) string{
	
	decoded := ""
	current := root

	for _, bit := range encoded{

		// esquerda
		if bit == '0'{
			current = current.Left
		}

		// direita
		if bit == '1'{
			current = current.Right
		}

		//encontrou caractere
		if current.IsLeaf(){

			decoded += string(current.Char)

			current = root
		}
	}

	return decoded
}