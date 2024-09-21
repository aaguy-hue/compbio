package main

import (
	"math/rand"
)

func WeightedDie() int {
	num := rand.Float64()
	if num <= 0.5 {
		return 3
	} else if num <= 0.6 {
		return 1
	} else if num <= 0.7 {
		return 2
	} else if num <= 0.8 {
		return 4
	} else if num <= 0.9 {
		return 5
	} else {
		return 6
	}
}

func EstimatePi(numPoints int) float64 {
	var count float64 = 0
	for i := 0; i < numPoints; i++ {
		if ThrowDart() {
			count++
		}
	}
	return count / float64(numPoints) * 4
}

func ThrowDart() bool {
	a := rand.Float64()*2 - 1
	b := rand.Float64()*2 - 1
	if a*a+b*b <= 1 {
		return true
	} else {
		return false
	}
}

// func main() {
// 	fmt.Println(EstimatePi(10000000000))
// }
