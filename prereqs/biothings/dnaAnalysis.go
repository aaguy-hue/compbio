package main

import (
    "fmt"
    "bufio"
    "os"
    "strconv"
    "strings"
    "sort"
)

func FindClumps(text string, k, L, t int) []string {
    var retval []string
    for i := range text {
        if i > len(text)-L {
            break
        }
        table := FrequencyTable(text[i:i+L], k)
        for key, val := range table {
            if val >= t && !sliceContains(retval, key) {
                retval = append(retval, key)
            }
        }
    }
    return retval
}

func SkewArray(genome string) []int {
	// Find the G-C skew
    var retval []int
    retval = append(retval, 0)
    
    var skew int
    for _, ch := range genome {
        if string(ch) == "G" {
            skew += 1
        } else if string(ch) == "C" {
            skew -= 1
        }
        retval = append(retval, skew)
    }
    return retval
}

func MinimumSkew(genome string) []int {
    skewArr := SkewArray(genome)
    min := MinIntegerArray(skewArr)
    var retval []int
    for i, s := range skewArr {
        if s == min {
            retval = append(retval, i)
        }
    }
    return retval
}