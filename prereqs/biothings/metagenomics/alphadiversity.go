package main
import (
    "os"
    "bufio"
    "strconv"
    "strings"
    "fmt"
)

// Richness and the Simpson's Index are two measures of Alpha Diversity, which is diversity within 
// a habitat
// Richness is to Simpson's Index as Range is to Standard Deviation

func Richness(sample map[string]int) int {
	// richness = how many distinct species, sampe = freq table
    var richness int
    for _, v := range sample {
        if v > 0 {
            richness += 1
        }
    }
    return richness
}

func SimpsonsIndex(sample map[string]int) float64 {
    // The probability of drawing something twice from a freq table, better than richness for
    // measuring diversity (accounts for skew in data). If the probability is lower, the data
    // is more even.
    total := SumOfValues(sample)
    totalsq := float64(total * total)
    var index float64
    for _, v := range sample {
        v2 := float64(v)
        index += (v2*v2)/totalsq
    }
    return index
}