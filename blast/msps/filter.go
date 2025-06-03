package msps

import (
	"blast/util"
)

func TrimMsps(msps []MaximalSegmentPair, minScore int) []MaximalSegmentPair {
	var trimmed []MaximalSegmentPair
	for _, msp := range msps {
		if msp.Score >= minScore {
			trimmed = append(trimmed, msp)
		}
	}
	return trimmed
}

func EvaluateMsps(msps []MaximalSegmentPair, q float32) []MaximalSegmentPair {
	var evaluated []MaximalSegmentPair
	for _, msp := range msps {
		if util.CalculatePValue(msp.Score, len(msp.DbSequence)) <= q {
			msp.PValue = util.CalculatePValue(msp.Score, len(msp.DbSequence))
			evaluated = append(evaluated, msp)
		}
	}
	return evaluated
}
