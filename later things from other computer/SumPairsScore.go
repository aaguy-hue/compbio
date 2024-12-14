package main

type Alignment []string

// Please do not remove package declarations because these are used by the autograder. If you need additional packages, then you may declare them above.

// Insert your SumPairsScore() function here, along with any subroutines that you need.
func SumPairsScore(align1 Alignment, align2 Alignment,
	idx1 int, idx2 int, match float64, mismatch float64, gap float64) float64 {
	finalScore := 0.0

	for i := 0; i < len(align1); i++ {
		for j := 0; j < len(align2); j++ {
			first := align1[i][idx1]
			second := align2[j][idx2]

			if first == '-' && second != '-' || first != '-' && second == '-' {
				finalScore -= gap
			} else if first != second {
				finalScore -= mismatch
			} else if first == second && first != '-' { // if there are two gaps, don't count it as a match
				finalScore += match
			}
		}
	}

	return finalScore
}
