package kmers

type KmerSearchResult struct {
	Kmer       string // The k-mer that was found
	QPosition  int    // The position in the query sequence where the k-mer was found
	DbPosition int    // The position in the database where the k-mer was found
	Score      int    // The score of the k-mer
}

// This function searches for k-mers in a given database.
func SearchKmersInDb(kmers []Kmer, db string, k int) []KmerSearchResult {
	/*
		Note regarding optimization:
		- We could just have a nested loop here, but that would be O(n*m) where n is the number of k-mers and m is the length of the database.
		- Instead, we use a hash set to store the k-mers, which allows us to check for existence in O(1) time.
		- Generating the hash set is O(n), where n is the number of k-mers.
		- Then, we slide a window of size k over the database string and check if the substring exists in the hash set, which is O(m) where m is the length of the database.
		- This makes the overall complexity O(n + m), where n is the number of k-mers and m is the length of the database.
		- This takes more memory, but is much faster for large databases.
	*/
	kmerSet := make(map[string]bool)

	// Make hash set of k-mers for faster lookup
	for _, kmer := range kmers {
		kmerSet[kmer.Kmer] = true
	}

	foundKmers := make([]KmerSearchResult, 0)

	// Slide window over the database
	for i := 0; i <= len(db)-k; i++ {
		substr := db[i : i+k]
		if kmerSet[substr] {
			foundKmers = append(foundKmers, KmerSearchResult{
				Kmer:       substr,
				DbPosition: i,
				QPosition:  i,
			})
		}
	}

	return foundKmers
}
