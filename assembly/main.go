package main

import (
	"fmt"
)

func main() {
	// TestGreedyAssembler()
	SARSOverlapNetwork()
}

func SARSOverlapNetwork() {
	sarsGenome := ReadGenomeFromFASTA("Data/SARS-CoV_genome.fasta")
	fmt.Println("SARS-CoV genome read.")

	readLength := 150
	probability := 0.1
	reads := SimulateReadsClean(sarsGenome, readLength, probability)
	fmt.Println("Reads simulated! We have", len(reads), "total reads. Now building overlap graph.")

	match := 1.0
	mismatch := 5.0
	gap := 1.0
	threshold := 40.0
	adjList := MakeOverlapNetwork(reads, match, mismatch, gap, threshold)
	fmt.Println("Overlap network generated!")

	fmt.Println("The graph has", len(adjList), "reads")
	fmt.Println("Average outdegree:", AverageOutDegree(adjList))
}

func TestGreedyAssembler() {
	length := 10000
	k := 50

	genome := GenerateRandomGenome(length)
	fmt.Println("Genome generated.")

	kmers := KmerComposition(genome, k)
	shuffledKmers := ShuffleStrings(kmers)
	fmt.Println("Kmers generated and shuffled, running assembler.")

	reconstructedGenome := GreedyAssembler(shuffledKmers)
	AssertEquals(genome, reconstructedGenome, "genome is not equal to reconstructed genome.")

	reconstructedKmers := KmerComposition(reconstructedGenome, k)
	if StringSliceEquals(kmers, reconstructedKmers) {
		fmt.Println("Test for greedy assembler passed.")
	} else {
		fmt.Println("Test for greedy assembler failed.")
	}
}

func AssertEquals(a, b any, errorMsg string) {
	if a != b {
		panic("Assertion failed: " + errorMsg)
	}
}
