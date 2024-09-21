package main

import (
    "fmt"
    "bufio"
    "os"
    "strings"
)

func Complement(dna string) string {
    retval := ""
    conversion_table := map[string]string{
        "A": "T",
        "C": "G",
        "G": "C",
        "T": "A",
    }
    for _, ch := range dna {
        retval += conversion_table[string(ch)]
    }
    return retval
}

func Reverse(s string) string {
    retval := make([]byte, len(s))
    for i := range s {
        retval[i] = s[len(s)-i-1]
    }
    return string(retval)
}

func ReverseComplement(dna string) string {
    return Reverse(Complement(dna))
}
