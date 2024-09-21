package main

import ( 
    "fmt"
    "math"
    "strconv"
    "math/rand" // this should be helpful!
)

func RollDie() int {
    return rand.Intn(6)+1
}

func SumDice(numDice int) int {
    ans := 0
    for i := 0; i < numDice; i++ {
        ans += RollDie()
    }
    return ans
}

