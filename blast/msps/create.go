package msps

import (
	"blast/kmers"
)

func CreateMaximalSegmentPairs(seeds []kmers.KmerSearchResult, query, db string, k int, xDropoff int) []MaximalSegmentPair {
	msps := make([]MaximalSegmentPair, 0)

	for _, seed := range seeds {
		msps = append(msps, CreateMsp(seed, query, db, k, xDropoff))
	}

	return msps
}

func CreateMsp(seed kmers.KmerSearchResult, query, db string, k int, xDropoff int) MaximalSegmentPair {
	// At the beginning, for DNA blast there will be exact matches, so we can use the kmer length as the score.
	score := k
	maxScore := score

	dbStart := seed.DbPosition
	dbEnd := seed.DbPosition + k - 1
	qStart := seed.QPosition
	qEnd := seed.QPosition + k - 1

	bestDbStart, bestDbEnd := dbStart, dbEnd
	bestQStart, bestQEnd := qStart, qEnd

	// Start by going left
	for dbStart > 0 && qStart > 0 {
		dbStart--
		qStart--

		if db[dbStart] == query[qStart] {
			score += 1
		} else {
			score -= 1
		}

		if score > maxScore {
			maxScore = score
			bestDbStart, bestQStart = dbStart, qStart
		}

		if score < maxScore-xDropoff {
			break
		}
	}

	// Go to the right
	for dbEnd < len(db)-1 && qEnd < len(query)-1 {
		dbEnd++
		qEnd++
		if db[dbEnd] == query[qEnd] {
			score += 1
		} else {
			score -= 1
		}

		if score > maxScore {
			maxScore = score
			bestDbEnd, bestQEnd = dbEnd, qEnd
		}

		if score < maxScore-xDropoff {
			break
		}
	}

	return MaximalSegmentPair{
		DbStart:    bestDbStart,
		DbEnd:      bestDbEnd,
		QStart:     bestQStart,
		QEnd:       bestQEnd,
		Score:      maxScore,
		DbSequence: db[bestDbStart : bestDbEnd+1],
		QSequence:  query[bestQStart : bestQEnd+1],
	}
}
