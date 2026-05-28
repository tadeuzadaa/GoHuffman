package frequency

func Count (text string) map[rune]int {

   //freq recebe um mapa de 'caracter' como chave e 'frequencia' como valor
   //'B': 1
	freq := make(map[rune]int)

	for _, char := range text {
		freq[char]++
	}

	return freq
}