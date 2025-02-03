package main

// import (
// 	"fmt"
// 	"example.com/m/v2/util"
// )

// type BlastAlignment struct {
// 	alignment Alignment
// 	eValue    float64
// }

// // BLAST (Basic Local Alignment Search Tool) is a heuristic for comparing a query against a database
// // Note: This function ignores introns so it only works for prokaryotes.
// //
// // Parameters:
// // - query: The query string to be searched against the database.
// // - db: The database to be searched.
// // - k: The length of the words (k-mers) the string is divided into.
// // - score: The scoring matrix.
// // - T: The threshold value for kmers.
// // - S: The threshold value for MSPs.
// // - Q: The threshold p-value for MSPs.
// // - mspCombineThreshold: The threshold p-value to combine two MSPs.
// //
// // Returns:
// // - Array of high-scoring local alignments
// func BLAST(query, db string, k int, score [][]int, T, S, Q, mspCombineThreshold int) ([]BlastAlignment, error) {
// 	if len(query) == 0 {
// 		return []BlastAlignment{}, fmt.Errorf("No query given to BLAST.")
// 	} else if len(db) == 0 {
// 		panic("No database given to BLAST.")
// 	}

// 	words := util.KmerComposition(query, k)
// 	highScoring := HighScoringKmers(words)
// 	seeds := SearchKmersInDb(highScoring, db)

// 	msps := MaximalSegmentPairs(seeds, db)
// 	msps = TrimMsps(msps, S)
// 	msps = EvaluateMsps(msps, Q)
// 	msps = CombineMsps(msps, mspCombineThreshold)

// 	alignments, eValues := LocalAlignMsps(msps)

// 	var retval []BlastAlignment
// 	for i := 0; i < len(alignments); i++ {
// 		retval[i] = BlastAlignment{
// 			alignment: alignments[i],
// 			eValue:    eValues[i],
// 		}
// 	}
// 	return retval, nil
// }
