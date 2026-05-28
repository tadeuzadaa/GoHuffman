package stats

func Calculate (original, encoded string) (int, int, float64){

	originalBits := len(original) 
	compressedBits := len(encoded)

	economy := (1 - float64(compressedBits) / float64(originalBits))

	return originalBits, compressedBits, economy
}