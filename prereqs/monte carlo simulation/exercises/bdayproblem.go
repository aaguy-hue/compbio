package main

import (
	"math/rand"
)

func HasRepeat(arr []int) bool {
	for i, a := range arr[0 : len(arr)-1] {
		for _, b := range arr[i+1 : len(arr)] {
			if a == b {
				return true
			}
		}
	}
	return false
}

func SimulateOneBirthdayTrial(numPeople int) bool {
	birthdays := make([]int, numPeople)
	for i := 0; i < numPeople; i++ {
		birthdays[i] = rand.Intn(365)
	}
	return HasRepeat(birthdays)
}

func SharedBirthdayProbability(numPeople, numTrials int) float64 {
	var count float64 = 0
	for i := 0; i < numTrials; i++ {
		if SimulateOneBirthdayTrial(numPeople) {
			count++
		}
	}
	return count / float64(numTrials)
}

// func main() {
// 	fmt.Println(SharedBirthdayProbability(40, 50000))
// }
