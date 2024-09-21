package util

func KmerComposition(genome string, k int) []string {
	n := len(genome)
	substrings := make([]string, n-k+1)

	for i := 0; i < len(genome)-k+1; i++ {
		substrings[i] = genome[i : i+k]
	}

	return substrings
}
