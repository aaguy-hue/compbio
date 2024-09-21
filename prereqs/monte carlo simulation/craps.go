package main

import ( 
    "fmt"
    "math"
    "strconv"
    "math/rand" // this should be helpful!
)

func PlayCrapsOnce() bool {
	// simulate a game of craps
    firstRoll := SumDice(2)
    if firstRoll == 2 || firstRoll == 3 || firstRoll == 12 {
        return false
    } else if firstRoll == 7 || firstRoll == 11 {
        return true
    }
    
    for {
        newRoll := SumDice(2)
        if newRoll == firstRoll {
            return true
        } else if newRoll == 7 {
            return false
        }
    }
    return false
}

func ComputeCrapsHouseEdge(numTrials int) float64 {
	// calculate the "house edge" (i.e. how rigged the game is for the casino) in a game of craps
    count := 0
    for i := 0; i < numTrials; i++ {
        outcome := PlayCrapsOnce()
        if outcome == true {
            count++
        } else {
            count--
        }
    }
    return float64(count)/float64(numTrials)
}
