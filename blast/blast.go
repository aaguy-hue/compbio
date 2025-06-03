package blast

import (
	"blast/kmers"
	"blast/msps"
	"fmt"
)

type BlastAlignment struct {
	alignment Alignment
	eValue    float64
}

/*
BLAST (Basic Local Alignment Search Tool) is a heuristic for comparing a query against a database

Note: This function ignores introns so it mainly works for prokaryotes 🥀.

Parameters:
  - query: The query string to be searched against the database.
  - db: The database to be searched.
  - k: The length of the words (k-mers) the string is divided into.
  - score: The scoring matrix, BLOSOM62 or PAM250, etc.
  - T: The threshold score for kmers.
  - S: The threshold score for MSPs.
  - Q: The threshold p-value for MSPs.
  - mspCombineThreshold: The threshold p-value to combine two MSPs.

Returns:
  - Array of high-scoring local alignments
*/
func DNABlast(query, db string, params DnaBlastParameters) ([]BlastAlignment, error) {
	// TODO: Make this more efficient by representing DNA as a byte array instead of a string since 2 bits are enough to represent a nucleotide (A, C, G, T).
	if len(query) == 0 {
		return []BlastAlignment{}, fmt.Errorf("No query given to BLAST.")
	} else if len(db) == 0 {
		panic("No database given to BLAST.")
	} else if params.k <= 0 {
		return []BlastAlignment{}, fmt.Errorf("k must be greater than 0, got %d", params.k)
	} else if params.S <= 0 {
		return []BlastAlignment{}, fmt.Errorf("S must be greater than 0, got %d", params.S)
	} else if params.Q <= 0 || params.Q > 1 {
		return []BlastAlignment{}, fmt.Errorf("Q must be between 0 and 1, got %f", params.Q)
	} else if params.mspCombineThreshold <= 0 || params.mspCombineThreshold > 1 {
		return []BlastAlignment{}, fmt.Errorf("mspCombineThreshold must be between 0 and 1, got %f", params.mspCombineThreshold)
	}

	words := kmers.CreateKmers(query, params.k)
	seeds := kmers.SearchKmersInDb(words, db, params.k)

	pairs := msps.CreateMaximalSegmentPairs(seeds, query, db, params.k, params.X)
	pairs = msps.TrimMsps(pairs, params.S)
	pairs = msps.EvaluateMsps(pairs, params.Q)
	pairs = msps.CombineMsps(pairs, params.mspCombineThreshold)

	alignments, eValues := LocalAlignMsps(pairs)

	var retval []BlastAlignment
	for i := 0; i < len(alignments); i++ {
		retval = append(retval, BlastAlignment{
			alignment: alignments[i],
			eValue:    eValues[i],
		})
	}
	return retval, nil
}

func DNABlastDefaults(query, db string) ([]BlastAlignment, error) {
	return DNABlast(query, db, DnaBlastParameters{
		k:                   11,
		X:                   20,
		S:                   20,
		Q:                   0.05,
		mspCombineThreshold: 0.05,
		match:               1,
		mismatch:            -1,
		gapOpen:             -2,
		gapExtend:           -1,
	})
}

func ProteinBlast(query, db string, score [][]int, params ProteinBlastParameters) ([]BlastAlignment, error) {
	if len(query) == 0 {
		return []BlastAlignment{}, fmt.Errorf("No query given to BLAST.")
	} else if len(db) == 0 {
		panic("No database given to BLAST.")
	}

	words := kmers.CreateKmers(query, params.k)
	// TODO: Use a scoring matrix to filter out low-scoring kmers.
	// This is only needed for protein BLAST, not DNA BLAST, so it's ok for now
	// highScoring := kmers.HighScoringKmers(words)\
	seeds := kmers.SearchKmersInDb(words, db)

	msps := MaximalSegmentPairs(seeds, db)
	msps = TrimMsps(msps, params.S)
	msps = EvaluateMsps(msps, params.Q)
	msps = CombineMsps(msps, params.mspCombineThreshold)

	alignments, eValues := LocalAlignMsps(msps)

	var retval []BlastAlignment
	for i := 0; i < len(alignments); i++ {
		retval = append(retval, BlastAlignment{
			alignment: alignments[i],
			eValue:    eValues[i],
		})
	}
	return retval, nil

}

func ProteinBlastDefaults(query, db string) ([]BlastAlignment, error) {
	return ProteinBlast(query, db, EmptyMatrix(), ProteinBlastParameters{
		k:                   3,
		T:                   20,
		S:                   20,
		Q:                   0.05,
		mspCombineThreshold: 0.05,
		scoreMatrix:         EmptyMatrix(),
	})
}

func EmptyMatrix() [][]int {
	return [][]int{}
}
