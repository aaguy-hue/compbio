package main

import "math/rand"

// GenerateRandomGenome takes a parameter length and returns
// a random DNA string of this length where each nucleotide has equal probability.
func GenerateRandomGenome(length int) string {
	a := make([]byte, length)

	for i := 0; i < length; i++ {
		n := rand.Intn(4)

		if n == 0 {
			a[i] = 'A'
		} else if n == 1 {
			a[i] = 'T'
		} else if n == 2 {
			a[i] = 'G'
		} else {
			a[i] = 'C'
		}
	}

	return string(a)
}
