package main

import "math/rand"

// ShuffleStrings takes a collection of strings patterns as input.
// It returns a random shuffle of the strings.
func ShuffleStrings(patterns []string) []string {
	amountOfStrings := len(patterns)
	permutations := rand.Perm(amountOfStrings)
	shuffledStrings := make([]string, amountOfStrings)

	for i := 0; i < len(permutations); i++ {
		index := permutations[i]
		shuffledStrings[i] = patterns[index]
	}

	return shuffledStrings
}
