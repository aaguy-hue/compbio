package main

// CountSharedKmers takes two strings and an integer k. It returns the number of
// k-mers that are shared by the two strings.
func CountSharedKmers(str1, str2 string, k int) int {
	freq1 := make(map[string]int)
	freq2 := make(map[string]int)

	for i := 0; i < len(str1)-k+1; i++ {
		kmer := str1[i : i+k]
		freq1[kmer] += 1
	}
	for i := 0; i < len(str2)-k+1; i++ {
		kmer := str2[i : i+k]
		freq2[kmer] += 1
	}

	return SumOfMinima(freq1, freq2)
}

func SumOfMinima(freq1, freq2 map[string]int) int {
	sum := 0
	for key := range freq1 {
		sum += Min(freq1[key], freq2[key])
	}
	return sum
}
