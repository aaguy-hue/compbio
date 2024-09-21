package main

// GreedyAssembler takes a collection of strings and returns a genome whose
// k-mer composition is these strings. It makes the following assumptions.
// 1. "Perfect coverage" -- every k-mer is detected
// 2. No errors in reads
// 3. Every read has equal length (k)
// 4. DNA is single-stranded
func GreedyAssembler(reads []string) string {
	if len(reads) == 0 {
		panic("Error: No reads passed into GreedyAssembler")
	}

	genome := reads[0]
	reads = reads[1:]
	k := len(reads[0]) // length of each k-mer
	AssertSameLength(reads)

	for len(reads) > 0 {
		n := len(genome)
		startReq := genome[0 : k-1]
		endReq := genome[n-k+1 : n]

		for i := 0; i < len(reads); i++ {
			kmer := reads[i]
			if kmer[1:k] == startReq {
				genome = string(kmer[0]) + genome
				reads = append(reads[:i], reads[i+1:]...)
			}

			if len(reads) < 1 {
				break
			}

			if kmer[0:k-1] == endReq {
				genome = genome + string(kmer[k-1])
				reads = append(reads[:i], reads[i+1:]...)
			}
		}
	}

	return genome
}

// func main() {
// 	str := "CC CC CC CC CC CC"
// 	fmt.Println(GreedyAssembler(strings.Split(str, " ")))
// }
