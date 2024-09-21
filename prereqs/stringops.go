package main
import (
    "os"
    "bufio"
    "strings"
    "fmt"
)

func StartingIndices(pattern, text string) []int {
    var retval []int
    for i := range text {
        if text[i] == pattern[0] && len(text)-i >= len(pattern) && text[i:i+len(pattern)] == pattern {
            retval = append(retval, i)
        }
    }
    return retval
}

func PatternCount(pattern, text string) int {
    var retval int = 0
    for i := range text {
        if text[i] == pattern[0] && len(text)-i >= len(pattern) && text[i:i+len(pattern)] == pattern {
            retval += 1
        }
    }
    return retval
}
