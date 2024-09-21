package main
import (
    "os"
    "bufio"
    "strconv"
    "strings"
    "fmt"
)

// Beta diversity is the diversity between two habitats i.e. how different two habitats are
// Bray-Curtis distance and Jaccard distance are both good metrics
// closer to 0 = very similar, closer to 1 = very different

func BrayCurtisDistance(sample1, sample2 map[string]int) float64 {
    return 1-float64(SumOfMinima(sample1, sample2)) / (float64(SumOfValues(sample1) + SumOfValues(sample2))/2)
}

func BrayCurtisDistance(sample1, sample2 map[string]int) float64 {
    return 1-float64(SumOfMinima(sample1, sample2)) / (float64(SumOfValues(sample1) + SumOfValues(sample2))/2)
}