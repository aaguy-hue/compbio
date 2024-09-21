package main

// Minimizer takes a string text and an integer k as input.
// It returns the k-mer of text that is lexicographically minimum.
func Minimizer(text string, k int) string {
	kmers := KmerComposition(text, k)

	minimizer := kmers[0]

	for i := 1; i < len(kmers); i++ {
		kmer := kmers[i]
		//fmt.Println(minimizer, kmer, CompareStrings(minimizer, kmer))
		if CompareStrings(minimizer, kmer) == kmer {
			minimizer = kmer
		}
	}

	return minimizer
}

func CompareStrings(str1, str2 string) string {
	if len(str1) != len(str2) {
		panic("Strings passed into CompareStrings have unequal length")
	}

	for i := 0; i < len(str1); i++ {
		if str1[i] < str2[i] {
			return str1
		} else if str1[i] > str2[i] {
			return str2
		}
	}

	return str1
}
