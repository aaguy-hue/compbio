package main

func SimpsonsMap(allMaps map[string](map[string]int)) map[string]float64 {
	simpsonsMap := make(map[string]float64)
	for k, v := range allMaps {
		simpsonsMap[k] = SimpsonsIndex(v)
	}
	return simpsonsMap
}

func SimpsonsIndex(sample map[string]int) float64 {
	total := SampleTotal(sample)
	simpson := 0.0

	if total == 0 {
		panic("Error: Empty frequency map given to SimpsonsIndex()!")
	}

	for _, val := range sample {
		x := float64(val) / float64(total)
		simpson += x * x
	}
	return simpson
}

// SampleTotal takes a frequency map as input.
// It returns the sum of the counts for each string in a sample.
func SampleTotal(freqMap map[string]int) int {
	total := 0
	for _, val := range freqMap {
		total += val
	}
	return total
}
