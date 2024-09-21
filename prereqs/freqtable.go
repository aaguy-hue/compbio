package main

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
    "sort"
    "math"
)

func FrequencyTable(text string, k int) map[string]int {
    retval := make(map[string]int)
    for i := range text {
        if i > len(text)-k {
            break
        }
        retval[text[i:i+k]] += 1 
    }
    return retval
}

func MaxMapValue(dict map[string]int) int {
    var maxval int = int(math.Inf(-1))
    for _, val := range dict {
        if val > maxval {
            maxval = val
        }
    }
    return maxval
}

func SumOfValues(dict map[string]int) int {
    // adds the values in the table
    var sum int = 0
    for _, val := range dict {
        sum += val
    }
    return sum
}

func SumOfMinima(sample1, sample2 map[string]int) int {
    // compares two tables, and at each key, will choose the number with the minimum value and add
    var ans int
    for k, v := range sample1 {
        ans += Min2(v, sample2[k])
    }
    return ans
}

func SumOfMaxima(sample1, sample2 map[string]int) int {
    // compares two tables, and at each key, will choose the number with the maximum value and add
    var ans int
    for k, v := range sample1 {
        // if a key isn't in the other table, just add the value of the key
        if s2val, exists := sample2[k]; exists {
            ans += Max2(v, s2val)
        } else {
            ans += v
        }
    }
    // add the keys in table 2 that aren't in table 1
    for k, v := range sample2 {
        if _, exists := sample1[k]; !exists {
            ans += v
        }
    }
    return ans
}

func FindFrequentWords(text string, k int) []string {
    table := FrequencyTable(text, k)
    num := MaxMapValue(table)
    var retval []string
    for key, val := range table {
        if val == num {
            retval = append(retval, key)
        }
    }
    return retval
}
