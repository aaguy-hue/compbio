package kmers

type Kmer struct {
	Kmer      string
	QPosition int
}

func CreateKmers(genome string, k int) []Kmer {
	n := len(genome)
	substrings := make([]Kmer, n-k+1)

	for i := 0; i < len(genome)-k+1; i++ {
		substrings[i] = Kmer{Kmer: genome[i : i+k], QPosition: i}
	}

	return substrings
}
