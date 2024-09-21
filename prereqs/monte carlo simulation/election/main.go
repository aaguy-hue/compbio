package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println("Election simulation")

	electoralVotes := ReadElectoralVotes("data/electoralVotes.csv")
	polls := ReadPollingData("data/debates.csv")

	numTrials := 1000000
	marginOfError := 0.1

	probability1, probability2, probabilityTie := SimulateMultipleElections(polls, electoralVotes, numTrials, marginOfError)

	fmt.Println("Canidate 1 winning probability: ", probability1*100, "%")
	fmt.Println("Canidate 2 winning probability: ", probability2*100, "%")
	fmt.Println("Tie probability: ", probabilityTie*100, "%")
}

func SimulateMultipleElections(polls map[string]float64, electoralVotes map[string]uint,
	numTrials int, marginOfError float64) (float64, float64, float64) {
	winCount1 := 0
	winCount2 := 0
	tieCount := 0

	for i := 0; i < numTrials; i++ {
		votes1, votes2 := SimulateOneElection(polls, electoralVotes, marginOfError)
		if votes1 > votes2 {
			winCount1++
		} else if votes2 > votes1 {
			winCount2++
		} else {
			tieCount++
		}
	}

	probability1 := float64(winCount1) / float64(numTrials)
	probability2 := float64(winCount2) / float64(numTrials)
	probabilityTie := float64(tieCount) / float64(numTrials)
	return probability1, probability2, probabilityTie
}

func SimulateOneElection(polls map[string]float64, electoralVotes map[string]uint,
	marginOfError float64) (uint, uint) {
	var collegeVotes1 uint = 0
	var collegeVotes2 uint = 0

	for state, pollingValue := range polls {
		numVotes := electoralVotes[state]
		adjustedPoll := AddNoise(pollingValue, marginOfError)

		if adjustedPoll >= 0.5 {
			collegeVotes1 += numVotes
		} else {
			collegeVotes2 += numVotes
		}
	}

	return collegeVotes1, collegeVotes2
}

func AddNoise(pollingValue, marginOfError float64) float64 {
	x := rand.NormFloat64() // 95% chance of being between -2 and 2
	x /= 2.0
	x *= marginOfError // 95% chance of being between -marginOfError and marginOfError

	return x + pollingValue
}
