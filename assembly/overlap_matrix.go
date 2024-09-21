package main

// OverlapScoringMatrix takes a collection of reads along with alignment penalties.
// It returns a matrix in which mtx[i][j] is the overlap alignment score of
// reads[i] with reads[j].

func OverlapScoringMatrix(reads []string, match, mismatch, gap float64) [][]float64 {
	n := len(reads)
	mtx := make([][]float64, n)

	for i := range mtx {
		mtx[i] = make([]float64, n)
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if i != j {
				mtx[i][j] = ScoreOverlapAlignment(reads[i], reads[j], match, mismatch, gap)
			} else {
				mtx[i][j] = 0
			}
		}
	}

	return mtx
}
