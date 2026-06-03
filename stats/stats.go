package stats

func Calculate(original, encoded string) (int, int, float64) {

	originalBits := len([]rune(original)) * 8
	compressedBits := len(encoded)
	if originalBits == 0 {
		return originalBits, compressedBits, 0
	}

	economy := (1 - float64(compressedBits)/float64(originalBits)) * 100

	return originalBits, compressedBits, economy
}
