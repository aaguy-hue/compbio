package main

import (
	"os"      // for reading from files
	"strconv" // our old friend string conversion package
	"strings" // for working with strings
)

// return map of state to amount of electoral votes for that state
func ReadElectoralVotes(filename string) map[string]uint {
	electoralVotes := make(map[string]uint)

	file, err := os.ReadFile(filename)
	CheckErr(err)

	fileContent := string(file)
	lines := strings.Split(fileContent, "\n")

	for _, line := range lines {
		lineSlice := strings.Split(line, ",")
		numVotes, err := strconv.Atoi(lineSlice[1])
		CheckErr(err)

		electoralVotes[lineSlice[0]] = uint(numVotes)
	}

	return electoralVotes
}

// return map of state names to percentages of canidate 1
func ReadPollingData(filename string) map[string]float64 {
	// we only read in canidate 1's percentage since canidate 2 % = 1 - canidate 1 %
	percentages := make(map[string]float64)

	file, err := os.ReadFile(filename)
	CheckErr(err)

	lines := strings.Split(string(file), "\n")

	for _, line := range lines {
		lineSlice := strings.Split(line, ",")
		percent, err := strconv.ParseFloat(lineSlice[1], 64)
		CheckErr(err)

		percentages[lineSlice[0]] = percent / 100.0
	}

	return percentages
}

func CheckErr(err error) {
	if err != nil {
		panic(err)
	}
}
