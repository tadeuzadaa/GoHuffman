package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"

	"go-huffman/decoder"
	"go-huffman/encoder"
	"go-huffman/frequency"
	"go-huffman/stats"
	"go-huffman/tree"
	"go-huffman/utils"
)

const quitCommand = ":q"

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		clearConsole()
		printHeader()

		text, err := readInput(reader, "Digite um texto")
		if err != nil {
			fmt.Printf("Erro ao ler entrada: %v\n", err)
			return
		}
		if text == quitCommand {
			fmt.Println("Encerrando programa.")
			return
		}
		if text == "" {
			fmt.Println("Entrada vazia. Informe um texto para codificar.")
			if !waitNextAction(reader) {
				fmt.Println("Encerrando programa.")
				return
			}
			continue
		}

		runDemo(text)

		if !waitNextAction(reader) {
			fmt.Println("Encerrando programa.")
			return
		}
	}
}

func runDemo(text string) {
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
	if root == nil {
		fmt.Println("Não foi possível montar a árvore de Huffman.")
		return
	}

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

func readInput(reader *bufio.Reader, label string) (string, error) {
	fmt.Printf("%s (ou %s para sair): ", label, quitCommand)
	text, err := reader.ReadString('\n')
	if err != nil && !errors.Is(err, io.EOF) {
		return "", err
	}

	text = strings.TrimSpace(text)

	if errors.Is(err, io.EOF) && text == "" {
		return quitCommand, nil
	}

	return text, nil
}

func waitNextAction(reader *bufio.Reader) bool {
	for {
		fmt.Printf("\nPressione Enter para novo texto ou digite %s para sair: ", quitCommand)
		next, err := reader.ReadString('\n')
		if err != nil && !errors.Is(err, io.EOF) {
			fmt.Printf("\nErro ao ler entrada: %v\n", err)
			return false
		}

		next = strings.TrimSpace(next)
		if next == "" {
			return true
		}
		if next == quitCommand {
			return false
		}

		fmt.Printf("Opção inválida: %q\n", next)

		if errors.Is(err, io.EOF) {
			return false
		}
	}
}

func printHeader() {
	fmt.Println("================================")
	fmt.Println(" DEMONSTRAÇÃO DO CÓDIGO HUFFMAN")
	fmt.Println("================================")
	fmt.Println()
}

func clearConsole() {
	fmt.Print("\033[H\033[2J")
}
