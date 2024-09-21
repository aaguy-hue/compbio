package main

import (
	"math/rand"
)

func EuclidGCD(a, b int) int {
	// from mathoperations.go
	for a != b {
		if a > b {
			a = a - b
		} else {
			b = b - a
		}
	}
	return a
}

func RelativelyPrime(a, b int) bool {
	if EuclidGCD(a, b) == 1 {
		return true
	}
	return false
}

func RelativelyPrimeProbability(lowerBound, upperBound, numPairs int) float64 {
	if lowerBound == upperBound {
		if lowerBound == 1 {
			return 1
		} else {
			return 0
		}
	}

	var count float64 = 0
	for i := 0; i < numPairs; i++ {
		num := rand.Intn(upperBound-lowerBound+1) + lowerBound
		num2 := rand.Intn(upperBound-lowerBound+1) + lowerBound
		if RelativelyPrime(num, num2) {
			count++
		}
	}
	return count / float64(numPairs)
}

// func main() {
// 	fmt.Println(RelativelyPrimeProbability(1, 1000000000000, 1000000000))
// }
