package blast

type DnaBlastParameters struct {
	k                   int     // Length of the words (k-mers/w-mers) the string is divided into
	X                   int     // AKA X-drop, the amount that the MSP score can drop before we stop extending
	S                   int     // Threshold score for MSPs
	Q                   float32 // Threshold p-value for MSPs
	mspCombineThreshold float32 // Threshold p-value to combine two MSPs
	match               int     // The reward for a match
	mismatch            int     // The penalty for a mismatch
	gapOpen             int     // The penalty for opening a gap
	gapExtend           int     // The penalty for extending a gap
}

type ProteinBlastParameters struct {
	k                   int     // Length of the words (k-mers/w-mers) the string is divided into
	scoreMatrix         [][]int // Scoring matrix, like BLOSOM62 or PAM250
	X                   int     // AKA X-drop, the amount that the MSP score can drop before we stop extending
	T                   int     // Threshold score for k-mers
	S                   int     // Threshold score for MSPs
	Q                   float32 // Threshold p-value for MSPs
	mspCombineThreshold float32 // Threshold p-value to combine two MSPs
	match               int     // The reward for a match
	mismatch            int     // The penalty for a mismatch
	gapOpen             int     // The penalty for opening a gap
	gapExtend           int     // The penalty for extending a gap
}
