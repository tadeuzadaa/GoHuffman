package main

import (
    
	"fmt"
    "os"
	"bufio"
    
	"go-huffman/decoder"
	"go-huffman/encoder"
	"go-huffman/frequency"
	"go-huffman/stats"
	"go-huffman/tree"
	"go-huffman/utils"
)

func main() {

    fmt.Println("================================")
	fmt.Println(" DEMONSTRAÇÃO DO CÓDIGO HUFFMAN")
	fmt.Println("================================")

    fmt.Println()

    // entrada do usuário
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Digite um texto: ")

	text, _ := reader.ReadString('\n')

	// remove quebra de linha
	text = text[:len(text)-1]

	fmt.Println()

	// frequências
	freq := frequency.Count(text)

	fmt.Println("===== FREQUÊNCIAS =====")

	for char, count := range freq {
		fmt.Printf("%c -> %d\n", char, count)
	}

	fmt.Println()

	// árvore
	root := tree.BuildTree(freq)

	fmt.Println("===== ÁRVORE =====")

	utils.PrintTree(root, 0)

	fmt.Println()

	// códigos
	codes := make(map[rune]string)

	encoder.GenerateCodes(
		root,
		"",
		codes,
	)

	fmt.Println("===== CÓDIGOS =====")

	for char, code := range codes {
		fmt.Printf("%c -> %s\n", char, code)
	}

	fmt.Println()

	// codificação
	encoded := encoder.Encode(
		text,
		codes,
	)

	fmt.Println("===== CODIFICADO =====")

	fmt.Println(encoded)

	fmt.Println()

	// decodificação
	decoded := decoder.Decode(
		encoded,
		root,
	)

	fmt.Println("===== DECODIFICADO =====")

	fmt.Println(decoded)

	fmt.Println()

	// estatísticas
	originalBits,
	compressedBits,
	economy := stats.Calculate(
		text,
		encoded,
	)

	fmt.Println("===== ESTATÍSTICAS =====")

	fmt.Printf(
		"Original: %d bits\n",
		originalBits,
	)

	fmt.Printf(
		"Comprimido: %d bits\n",
		compressedBits,
	)

	fmt.Printf(
		"Economia: %.2f%%\n",
		economy,
	)
}    